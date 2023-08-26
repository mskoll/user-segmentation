package repo

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
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

	var userId int
	userQuery := fmt.Sprintf("INSERT INTO %s (username) VALUES ($1) RETURNING id", userTable)

	row := tx.QueryRow(userQuery, user.Username)
	if err = row.Scan(&userId); err != nil {
		tx.Rollback()
		return 0, err
	}

	return userId, tx.Commit()
}

func (r *UserRepo) GetById(ctx context.Context, id int) (entity.User, error) {
	var user entity.User

	userQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", userTable)

	err := r.db.Get(&user, userQuery, id)
	if err != nil {
		return entity.User{}, err
	}
	var segments []entity.Segment
	segmentQuery := fmt.Sprintf("SELECT st.name FROM %s st INNER JOIN %s slt ON st.id = slt.segment_id WHERE slt.user_id = $1", segmentTable, userSegmentTable)
	err = r.db.Select(&segments, segmentQuery, id)

	return user, nil
}
func (r *UserRepo) AddDeleteSegment(ctx context.Context, id int, toAdd []entity.Segment, toDelete []entity.Segment) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	segmentsToDelete := make([]entity.UserSegment, len(toDelete))
	for i, segment := range toDelete {
		segmentQuery := spew.Sprintf("SELECT id FROM %s WHERE NAME LIKE $1", segmentTable)
		err = r.db.Get(&segmentsToDelete[i].SegmentId, segmentQuery, segment.Name)
		segmentsToDelete[i].UserId = id
	}

	segmentsToAdd := make([]entity.UserSegment, len(toAdd))
	for i, segment := range toAdd {
		segmentQuery := spew.Sprintf("SELECT id FROM %s WHERE NAME LIKE $1", segmentTable)
		err = r.db.Get(&segmentsToAdd[i].SegmentId, segmentQuery, segment.Name)
		segmentsToDelete[i].UserId = id
	}

	_, err = r.db.NamedExec(fmt.Sprintf("INSERT INTO %s (user_id, segment_id) VALUES (:user_id, :segment_id)",
		userSegmentTable), segmentsToAdd)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
