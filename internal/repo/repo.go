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
	DeleteSegment(ctx context.Context, id int, toDelete []string) error
}

type Segment interface {
	CreateSegment(ctx context.Context, segment entity.Segment) (int, error)
	Delete(ctx context.Context, id int) error
}

type Operation interface {
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
