package service

import (
	"context"
	"userSegmentation/internal/entity"
	"userSegmentation/internal/repo"
)

type UserService struct {
	repo repo.User
}

func NewUser(repo repo.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) UserById(ctx context.Context, id int) (entity.SegmentList, error) {

	var segments entity.SegmentList

	user, err := s.repo.UserById(ctx, id)
	if err != nil {
		return entity.SegmentList{}, err
	}

	segments.User = user

	segments.Segments, err = s.repo.UsersSegments(ctx, id)

	return segments, err
}

func (s *UserService) CreateUser(ctx context.Context, user entity.User) (int, error) {

	id, err := s.repo.CreateUser(ctx, user)

	return id, err
}

func (s *UserService) AddDeleteSegment(ctx context.Context, segments entity.AddDelSegments) error {

	if len(segments.ToAdd) != 0 {

		if err := s.repo.AddSegment(ctx, segments.Id, segments.ToAdd); err != nil {
			return err
		}
	}

	if len(segments.ToDel) != 0 {

		if err := s.repo.DeleteSegmentFromUser(ctx, segments.Id, segments.ToDel); err != nil {
			return err
		}
	}

	return nil
}
