package repo

import (
	"context"
	"fmt"
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
	segmentQuery := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", segmentTable)

	row := tx.QueryRow(segmentQuery, segment.Name)
	if err = row.Scan(&segmentId); err != nil {
		tx.Rollback()
		return 0, err
	}

	return segmentId, tx.Commit()
}

func (r *SegmentRepo) Delete(ctx context.Context, name string) error {
	//TODO implement me
	panic("implement me")
}
