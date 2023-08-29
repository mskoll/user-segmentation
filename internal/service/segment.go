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
		return 0, err
	}

	if segment.Percent == 0 {
		return id, nil
	}

	userIds, err := s.repo.UserIdsList(ctx, segment.Percent)
	if len(userIds) == 0 {
		return id, nil
	}
	if err != nil {
		return 0, err
	}

	userSegment := make([]entity.UserSegment, len(userIds))
	for i := range userIds {
		userSegment[i].UserId = userIds[i]
		userSegment[i].SegmentId = id
	}

	err = s.repo.AddUser(ctx, userSegment)

	return id, err
}

func (s *SegmentService) DeleteSegment(ctx context.Context, name string) error {
	err := s.repo.DeleteSegment(ctx, name)
	return err
}
