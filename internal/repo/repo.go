package repo

import (
	"github.com/jmoiron/sqlx"
	"userSegmentation/internal/entity"
)

type User interface {
	CreateUser(user entity.User) (int, error)
	UserById(id int) (entity.User, error)
	UsersSegments(id int) ([]entity.Segment, error)
	AddSegment(segments []entity.UserSegment) error
	DeleteSegmentFromUser(segments []entity.UserSegment) error
	Operations(usersOperations entity.UserOperations) ([]entity.Operation, error)
	SegmentsIdsByName(segments []entity.SegmentToUser) ([]int, error)
}

type Segment interface {
	CreateSegment(segment entity.Segment) (int, error)
	AddUser(userSegment []entity.UserSegment) error
	UserIdsList(percent int) ([]int, error)
	DeleteSegment(name string) error
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
