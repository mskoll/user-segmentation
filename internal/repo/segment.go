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

func (r *SegmentRepo) AddToPercentUsers(ctx context.Context, segment entity.Segment) error {

	var userIds []int
	usersQuery := "SELECT id FROM users ORDER BY random() LIMIT (SELECT count(*) FROM users)*$1/100"
	err := r.db.Select(&userIds, usersQuery, segment.Percent)

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	usersSegment := make([]entity.UserSegment, len(userIds))
	for i, us := range usersSegment {
		us.UserId = userIds[i]
		us.SegmentId = segment.Id
	}

	usersSegmentQuery := "INSERT INTO user_segment (user_id, segment_id) VALUES (:user_id, :segment_id) ON CONFLICT DO NOTHING"

	if _, err = r.db.NamedExec(usersSegmentQuery, usersSegment); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
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
