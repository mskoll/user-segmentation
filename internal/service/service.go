package service

import (
	"context"
	"userSegmentation/internal/entity"
	"userSegmentation/internal/repo"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type User interface {
	CreateUser(ctx context.Context, user entity.User) (int, error)
	UserById(ctx context.Context, id int) (entity.SegmentList, error)
	AddDeleteSegment(ctx context.Context, segments entity.AddDelSegments) error
	Operations(ctx context.Context, userOperations entity.UserOperations) ([]entity.Operation, error)
}

type Segment interface {
	CreateSegment(ctx context.Context, segment entity.Segment) (int, error)
	DeleteSegment(ctx context.Context, name string) error
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
