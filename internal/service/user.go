package service

import (
	"context"
	"log"
	"userSegmentation/internal/entity"
	"userSegmentation/internal/repo"
)

type UserService struct {
	repo repo.User
}

func (s *UserService) AddDeleteSegment(ctx context.Context, id int, toAdd []entity.Segment, toDelete []entity.Segment) error {
	//TODO implement me
	panic("implement me")
}

func NewUser(repo repo.User) *UserService {
	return &UserService{repo: repo}
}
func (u *UserService) GetById(ctx context.Context, id int) (entity.User, error) {
	user, err := u.repo.GetById(ctx, id)
	log.Println("in service")
	return user, err
}

func (u *UserService) CreateUser(ctx context.Context, user entity.User) (int, error) {
	userId, err := u.repo.CreateUser(ctx, user)
	return userId, err
}
