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
	AddSegment(ctx context.Context, id int, toAdd []string) error
	DeleteSegmentFromUser(ctx context.Context, id int, toDelete []string) error
}

type Segment interface {
	CreateSegment(ctx context.Context, segment entity.Segment) (int, error)
	AddToPercentUsers(ctx context.Context, segment entity.Segment) error
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
