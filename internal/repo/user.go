package repo

import (
	"context"
	"github.com/jmoiron/sqlx"
	"userSegmentation/internal/entity"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(ctx context.Context, user entity.User) (int, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	userQuery := "INSERT INTO users (username) VALUES ($1) RETURNING id"
	row := tx.QueryRow(userQuery, user.Username)

	var userId int
	if err = row.Scan(&userId); err != nil {
		tx.Rollback()
		return 0, err
	}

	return userId, tx.Commit()
}

func (r *UserRepo) UserById(ctx context.Context, id int) (entity.User, error) {

	var user entity.User

	userQuery := "SELECT * FROM users WHERE id = $1"
	err := r.db.Get(&user, userQuery, id)

	return user, err
}

func (r *UserRepo) UsersSegments(ctx context.Context, id int) ([]entity.Segment, error) {

	var segments []entity.Segment

	segmentQuery := "SELECT st.id, st.name FROM segment st INNER JOIN user_segment ust ON st.id = ust.segment_id WHERE ust.user_id = $1"
	err := r.db.Select(&segments, segmentQuery, id)

	return segments, err
}

func (r *UserRepo) AddSegment(ctx context.Context, id int, toAdd []string) error {
	// todo: add check if exists
	segments := make([]entity.UserSegment, len(toAdd))

	for i, segment := range toAdd {

		segments[i].UserId = id
		segmentQuery := "SELECT id FROM segment WHERE name LIKE $1"
		if err := r.db.Get(&segments[i].SegmentId, segmentQuery, segment); err != nil {
			return err
		}
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	usersSegmentQuery := "INSERT INTO user_segment (user_id, segment_id) VALUES (:user_id, :segment_id)"

	if _, err = r.db.NamedExec(usersSegmentQuery, segments); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *UserRepo) DeleteSegment(ctx context.Context, id int, toDelete []string) error {

	segments := make([]int, len(toDelete))

	for i, segment := range toDelete {

		segmentQuery := "SELECT id FROM segment WHERE name LIKE $1"
		if err := r.db.Get(&segments[i], segmentQuery, segment); err != nil {
			return err
		}
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	for _, segment := range segments {

		usersSegmentQuery := "DELETE FROM user_segment WHERE user_id = $1 AND segment_id = $2"
		if _, err = r.db.Exec(usersSegmentQuery, id, segment); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
