package repo

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"userSegmentation/internal/entity"
	"userSegmentation/internal/lib/errTypes"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(user entity.User) (int, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("UserRepo.CreateUser: %s", err.Error()))
	}

	userQuery := "INSERT INTO users (username) VALUES ($1) RETURNING id"
	row := tx.QueryRow(userQuery, user.Username)

	var userId int
	if err = row.Scan(&userId); err != nil {
		tx.Rollback()
		return 0, errors.Wrap(err, fmt.Sprintf("UserRepo.CreateUser: %s", err.Error()))
	}

	return userId, tx.Commit()
}

func (r *UserRepo) UserById(id int) (entity.User, error) {

	var user entity.User

	userQuery := "SELECT * FROM users WHERE id = $1"

	err := r.db.Get(&user, userQuery, id)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, errors.Wrap(errTypes.ErrNotFound, fmt.Sprintf("User %v not found", id))
		}

		return entity.User{}, errors.Wrap(err, fmt.Sprintf("UserRepo.UserById: %s", err.Error()))
	}

	return user, nil
}

func (r *UserRepo) UsersSegments(id int) ([]entity.Segment, error) {

	var segments []entity.Segment

	segmentQuery := "SELECT st.id, st.name, st.percent FROM segment st INNER JOIN user_segment ust ON st.id = ust.segment_id " +
		"WHERE ust.user_id = $1 AND (ust.deleted_at IS NULL OR ust.deleted_at > now())"

	err := r.db.Select(&segments, segmentQuery, id)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("UserRepo.UsersSegments: %s", err.Error()))
	}

	return segments, nil
}

func (r *UserRepo) AddSegment(segments []entity.UserSegment) error {

	tx, err := r.db.Begin()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("UserRepo.AddSegment: %s", err.Error()))
	}

	usersSegmentQuery := "INSERT INTO user_segment (user_id, segment_id, deleted_at) VALUES (:user_id, :segment_id, :deleted_at)"

	if _, err = r.db.NamedExec(usersSegmentQuery, segments); err != nil {
		tx.Rollback()
		return errors.Wrap(err, fmt.Sprintf("UserRepo.AddSegment: %s", err.Error()))
	}

	return tx.Commit()
}

func (r *UserRepo) DeleteSegmentFromUser(segments []entity.UserSegment) error {

	tx, err := r.db.Begin()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("UserRepo.DeleteSegmentFromUser: %s", err.Error()))
	}

	for _, segment := range segments {

		usersSegmentQuery := "UPDATE user_segment SET deleted_at = now() WHERE user_id = $1 AND segment_id = $2"
		if _, err = r.db.Exec(usersSegmentQuery, segment.UserId, segment.SegmentId); err != nil {
			tx.Rollback()
			return errors.Wrap(err, fmt.Sprintf("UserRepo.DeleteSegmentFromUser: %s", err.Error()))
		}
	}

	return tx.Commit()
}

func (r *UserRepo) Operations(usersOperations entity.UserOperations) ([]entity.Operation, error) {

	var operations []entity.Operation

	operationsQuery := "SELECT ust.user_id user_id, st.name segment_name, 'created' operation, ust.created_at datetime " +
		"FROM user_segment ust INNER JOIN segment st ON ust.segment_id = st.id " +
		"WHERE ust.user_id = $1 AND DATE_PART('month', ust.created_at) = $2 AND DATE_PART('year', ust.created_at) = $3 " +
		"UNION SELECT ust.user_id user_id, st.name segment_name, 'deleted' operation, ust.deleted_at datetime " +
		"FROM user_segment ust INNER JOIN segment st ON ust.segment_id = st.id " +
		"WHERE ust.user_id = $4 AND ust.deleted_at < NOW() AND ust.deleted_at IS NOT NULL AND " +
		"DATE_PART('month', ust.deleted_at) = $5 AND DATE_PART('year', ust.deleted_at) = $6 " +
		"ORDER BY datetime"

	if err := r.db.Select(&operations, operationsQuery, usersOperations.UserId, usersOperations.Month, usersOperations.Year,
		usersOperations.UserId, usersOperations.Month, usersOperations.Year); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("UserRepo.Operations: %s", err.Error()))
	}

	return operations, nil
}

func (r *UserRepo) SegmentsIdsByName(segments []entity.SegmentToUser) ([]int, error) {

	segmentQuery := "SELECT id FROM segment WHERE name LIKE $1"

	segmentIds := make([]int, len(segments))
	for i, segment := range segments {
		if err := r.db.Get(&segmentIds[i], segmentQuery, segment.Name); err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				return []int{}, errors.Wrap(errTypes.ErrNotFound, fmt.Sprintf("Segment %v not found", segment.Name))
			}

			return []int{}, errors.Wrap(err, fmt.Sprintf("UserRepo.SegmentsIdsByName: %s", err.Error()))
		}
	}

	return segmentIds, nil
}
