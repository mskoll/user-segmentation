package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"userSegmentation/internal/entity"
	"userSegmentation/internal/lib/errTypes"
)

type SegmentRepo struct {
	db *sqlx.DB
}

func NewSegment(db *sqlx.DB) *SegmentRepo {
	return &SegmentRepo{db: db}
}

func (r *SegmentRepo) CreateSegment(ctx context.Context, segment entity.Segment) (int, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("SegmentRepo.CreateSegment: %s", err.Error()))
	}

	var segmentId int
	segmentQuery := "INSERT INTO segment (name, percent) VALUES ($1, $2) RETURNING id"

	row := tx.QueryRow(segmentQuery, segment.Name, segment.Percent)
	if err = row.Scan(&segmentId); err != nil {
		tx.Rollback()
		return 0, errors.Wrap(err, fmt.Sprintf("SegmentRepo.CreateSegment: %s", err.Error()))
	}

	return segmentId, tx.Commit()
}

func (r *SegmentRepo) AddUser(ctx context.Context, userSegment []entity.UserSegment) error {

	tx, err := r.db.Begin()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("SegmentRepo.AddUser: %s", err.Error()))
	}

	usersSegmentQuery := "INSERT INTO user_segment (user_id, segment_id) VALUES (:user_id, :segment_id)"

	if _, err = r.db.NamedExec(usersSegmentQuery, userSegment); err != nil {
		tx.Rollback()
		return errors.Wrap(err, fmt.Sprintf("SegmentRepo.AddUser: %s", err.Error()))
	}

	return tx.Commit()
}

func (r *SegmentRepo) UserIdsList(ctx context.Context, percent int) ([]int, error) {

	var userIds []int

	userQuery := "SELECT id FROM users TABLESAMPLE BERNOULLI ($1)"
	if err := r.db.Select(&userIds, userQuery, percent); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("SegmentRepo.UserIdsList: %s", err.Error()))
	}

	return userIds, nil
}

func (r *SegmentRepo) DeleteSegment(ctx context.Context, name string) error {

	tx, err := r.db.Begin()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("SegmentRepo.DeleteSegment: %s", err.Error()))
	}

	segmentQuery := "UPDATE segment SET deleted_at = NOW() WHERE name LIKE $1 AND deleted_at IS NULL RETURNING id"
	row := tx.QueryRow(segmentQuery, name)

	var segmentId int
	if err = row.Scan(&segmentId); err != nil {
		tx.Rollback()
		if errors.Is(err, sql.ErrNoRows) {
			return errors.Wrap(errTypes.ErrNotFound, fmt.Sprintf("Segment %s not found", name))
		}
		return errors.Wrap(err, fmt.Sprintf("SegmentRepo.DeleteSegment: %s", err.Error()))
	}

	usersSegmentQuery := "UPDATE user_segment SET deleted_at = NOW() WHERE segment_id = $1 AND " +
		"(deleted_at IS NULL OR deleted_at > now())"
	row = tx.QueryRow(usersSegmentQuery, segmentId)

	return tx.Commit()
}
