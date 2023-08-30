package service

import (
	"userSegmentation/internal/entity"
	"userSegmentation/internal/repo"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type User interface {
	CreateUser(user entity.User) (int, error)
	UserById(id int) (entity.SegmentList, error)
	AddDeleteSegment(segments entity.AddDelSegments) error
	Operations(userOperations entity.UserOperations) ([]entity.Operation, error)
}

type Segment interface {
	CreateSegment(segment entity.Segment) (int, error)
	DeleteSegment(name string) error
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
