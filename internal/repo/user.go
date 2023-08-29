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

	segmentQuery := "SELECT st.id, st.name, st.percent FROM segment st INNER JOIN user_segment ust ON st.id = ust.segment_id " +
		"WHERE ust.user_id = $1 AND (ust.deleted_at IS NULL OR ust.deleted_at > now())"
	err := r.db.Select(&segments, segmentQuery, id)

	return segments, err
}

func (r *UserRepo) AddSegment(ctx context.Context, segments []entity.UserSegment) error {

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	usersSegmentQuery := "INSERT INTO user_segment (user_id, segment_id, deleted_at) VALUES (:user_id, :segment_id, :deleted_at)"

	if _, err = r.db.NamedExec(usersSegmentQuery, segments); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *UserRepo) DeleteSegmentFromUser(ctx context.Context, segments []entity.UserSegment) error {

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	for _, segment := range segments {

		usersSegmentQuery := "UPDATE user_segment SET deleted_at = now() WHERE user_id = $1 AND segment_id = $2"
		if _, err = r.db.Exec(usersSegmentQuery, segment.UserId, segment.SegmentId); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *UserRepo) Operations(ctx context.Context, usersOperations entity.UserOperations) ([]entity.Operation, error) {

	var operations []entity.Operation

	operationsQuery := "SELECT ust.user_id user_id, st.name segment_name, 'created' operation, ust.created_at datetime " +
		"FROM user_segment ust INNER JOIN segment st ON ust.segment_id = st.id " +
		"WHERE ust.user_id = $1 AND DATE_PART('month', ust.created_at) = $2 AND DATE_PART('year', ust.created_at) = $3 " +
		"UNION SELECT ust.user_id user_id, st.name segment_name, 'deleted' operation, ust.deleted_at datetime " +
		"FROM user_segment ust INNER JOIN segment st ON ust.segment_id = st.id " +
		"WHERE ust.user_id = $4 AND ust.deleted_at < NOW() AND ust.deleted_at IS NOT NULL AND " +
		"DATE_PART('month', ust.deleted_at) = $5 AND DATE_PART('year', ust.deleted_at) = $6 " +
		"ORDER BY datetime"

	if err := r.db.Select(&operations, operationsQuery, usersOperations.UserId, usersOperations.Month, usersOperations.Year, usersOperations.UserId, usersOperations.Month, usersOperations.Year); err != nil {
		return nil, err
	}

	return operations, nil
}

func (r *UserRepo) SegmentsIdsByName(ctx context.Context, segments []entity.SegmentToUser) ([]int, error) {

	segmentQuery := "SELECT id FROM segment WHERE name LIKE $1"

	segmentIds := make([]int, len(segments))
	for i, segment := range segments {
		if err := r.db.Get(&segmentIds[i], segmentQuery, segment.Name); err != nil {
			return []int{}, err
		}
	}

	return segmentIds, nil
}
