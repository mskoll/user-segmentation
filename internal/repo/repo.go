package repo

import (
	"context"
	"github.com/jmoiron/sqlx"
	"userSegmentation/internal/entity"
)

type User interface {
	CreateUser(ctx context.Context, user entity.User) (int, error)
	UserById(ctx context.Context, id int) (entity.User, error)
	UsersSegments(ctx context.Context, id int) ([]entity.Segment, error)
	AddSegment(ctx context.Context, segments []entity.UserSegment) error
	DeleteSegmentFromUser(ctx context.Context, segments []entity.UserSegment) error
	Operations(ctx context.Context, usersOperations entity.UserOperations) ([]entity.Operation, error)
	SegmentsIdsByName(ctx context.Context, segments []entity.SegmentToUser) ([]int, error)
}

type Segment interface {
	CreateSegment(ctx context.Context, segment entity.Segment) (int, error)
	AddUser(ctx context.Context, userSegment []entity.UserSegment) error
	UserIdsList(ctx context.Context, percent int) ([]int, error)
	DeleteSegment(ctx context.Context, name string) error
}

type Repository struct {
	User
	Segment
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		User:    NewUser(db),
		Segment: NewSegment(db),
	}
}
