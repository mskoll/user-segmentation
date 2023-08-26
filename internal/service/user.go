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
func (s *UserService) UserById(ctx context.Context, id int) ([]entity.Segment, error) {
	user, err := s.repo.UserById(ctx, id)
	return user, err
}

func (s *UserService) CreateUser(ctx context.Context, user entity.User) (int, error) {
	userId, err := s.repo.CreateUser(ctx, user)
	return userId, err
}

func (s *UserService) AddDeleteSegment(ctx context.Context, id int, toAdd []string, toDelete []string) error {
	err := s.repo.AddSegment(ctx, id, toAdd)
	if err != nil {
		return err
	}
	//err = s.repo.DeleteSegment(ctx, id, toDelete)
	return err
}
