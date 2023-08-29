package repo

import (
	"context"
	"github.com/jmoiron/sqlx"
	"userSegmentation/internal/entity"
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
		return 0, err
	}

	var segmentId int
	segmentQuery := "INSERT INTO segment (name, percent) VALUES ($1, $2) RETURNING id"

	row := tx.QueryRow(segmentQuery, segment.Name, segment.Percent)
	if err = row.Scan(&segmentId); err != nil {
		tx.Rollback()
		return 0, err
	}

	return segmentId, tx.Commit()
}

func (r *SegmentRepo) AddUser(ctx context.Context, userSegment []entity.UserSegment) error {

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	usersSegmentQuery := "INSERT INTO user_segment (user_id, segment_id) VALUES (:user_id, :segment_id)"

	if _, err = r.db.NamedExec(usersSegmentQuery, userSegment); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *SegmentRepo) UserIdsList(ctx context.Context, percent int) ([]int, error) {

	var userIds []int
	userQuery := "SELECT id FROM users TABLESAMPLE BERNOULLI ($1)"
	err := r.db.Select(&userIds, userQuery, percent)

	return userIds, err
}

func (r *SegmentRepo) DeleteSegment(ctx context.Context, name string) error {

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	segmentQuery := "UPDATE segment SET deleted_at = NOW() WHERE name LIKE $2 RETURNING id"
	row := tx.QueryRow(segmentQuery, name)

	var segmentId int
	if err = row.Scan(&segmentId); err != nil {
		tx.Rollback()
		return err
	}

	usersSegmentQuery := "UPDATE user_segment SET deleted_at = NOW() WHERE segment_id = $1"
	row = tx.QueryRow(usersSegmentQuery, segmentId)

	return tx.Commit()
}
