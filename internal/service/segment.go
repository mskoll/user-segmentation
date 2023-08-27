package service

import (
	"context"
	"userSegmentation/internal/entity"
	"userSegmentation/internal/repo"
)

type SegmentService struct {
	repo repo.Segment
}

func NewSegmentService(repo repo.Segment) *SegmentService {
	return &SegmentService{repo: repo}
}

func (s *SegmentService) CreateSegment(ctx context.Context, segment entity.Segment) (int, error) {
	id, err := s.repo.CreateSegment(ctx, segment)
	if err != nil {
		return id, err
	}
	if segment.Percent != 0 {
		segment.Id = id
		err = s.repo.AddToPercentUsers(ctx, segment)
	}
	return id, err
}

func (s *SegmentService) DeleteSegment(ctx context.Context, name string) error {
	err := s.repo.DeleteSegment(ctx, name)
	return err
}
