package service

import (
	"context"
	"userSegmentation/internal/entity"
	"userSegmentation/internal/repo"
)

type User interface {
	CreateUser(ctx context.Context, user entity.User) (int, error)
	UserById(ctx context.Context, id int) ([]entity.Segment, error)
	AddDeleteSegment(ctx context.Context, id int, toAdd []string, toDelete []string) error
}
type Segment interface {
	Create(ctx context.Context, segment entity.Segment) (int, error)
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
