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

func (s *SegmentService) Create(ctx context.Context, segment entity.Segment) (int, error) {
	id, err := s.repo.CreateSegment(ctx, segment)
	return id, err
}

func (s *SegmentService) Delete(ctx context.Context, name string) error {
	//TODO implement me
	panic("implement me")
}
