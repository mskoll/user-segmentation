package service

import (
	"context"
	"userSegmentation/internal/entity"
	"userSegmentation/internal/repo"
)

type User interface {
	CreateUser(ctx context.Context, user entity.User) (int, error)
	GetById(ctx context.Context, id int) (entity.User, error)
	AddDeleteSegment(ctx context.Context, id int, toAdd []entity.Segment, toDelete []entity.Segment) error
}
type Segment interface {
}
type Operation interface {
}
type Service struct {
	User
	Segment
}

func New(repo *repo.Repository) *Service {
	return &Service{
		User:    NewUser(repo),
		Segment: NewSegmentService(repo),
	}
}
