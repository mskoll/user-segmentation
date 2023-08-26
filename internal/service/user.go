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

func (s *UserService) AddDeleteSegment(ctx context.Context, id int, toAdd []string, toDelete []string) error {

	if len(toAdd) != 0 {

		if err := s.repo.AddSegment(ctx, id, toAdd); err != nil {
			return err
		}
	}

	if len(toDelete) != 0 {

		if err := s.repo.DeleteSegment(ctx, id, toDelete); err != nil {
			return err
		}
	}

	return nil
}
