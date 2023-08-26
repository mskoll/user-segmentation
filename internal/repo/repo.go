package repo

import (
	"context"
	"github.com/jmoiron/sqlx"
	"userSegmentation/internal/entity"
)

type User interface {
	CreateUser(ctx context.Context, user entity.User) (int, error)
	GetById(ctx context.Context, id int) (entity.User, error)
	AddDeleteSegment(ctx context.Context, id int, toAdd []entity.Segment, toDelete []entity.Segment) error
}

type Segment interface {
	Create(ctx context.Context, segment entity.Segment) (int, error)
	Delete(ctx context.Context, id int) error
	// AddUser(ctx context.Context, segmId int, userId int) error
	// DeleteUser(ctx context.Context, segmId int, userId int) error
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
