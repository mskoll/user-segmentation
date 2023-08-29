package service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"userSegmentation/internal/entity"
	"userSegmentation/internal/lib/errTypes"
	"userSegmentation/internal/repo"
)

type UserService struct {
	repo repo.User
}

func NewUser(repo repo.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) UserById(ctx context.Context, id int) (entity.SegmentList, error) {

	var segments entity.SegmentList

	user, err := s.repo.UserById(ctx, id)
	if err != nil {
		return entity.SegmentList{}, err
	}

	segments.User = user

	segments.Segments, err = s.repo.UsersSegments(ctx, id)

	return segments, err
}

func (s *UserService) CreateUser(ctx context.Context, user entity.User) (int, error) {

	id, err := s.repo.CreateUser(ctx, user)

	return id, err
}

func (s *UserService) AddDeleteSegment(ctx context.Context, segments entity.AddDelSegments) error {

	if len(segments.ToDel) != 0 {

		idsToDel, err := s.repo.SegmentsIdsByName(ctx, segments.ToDel)
		if err != nil {
			return err
		}

		userSegmentToDel := make([]entity.UserSegment, len(segments.ToDel))
		for i := range segments.ToDel {
			userSegmentToDel[i] = entity.UserSegment{
				UserId:    segments.UserId,
				SegmentId: idsToDel[i],
			}
		}

		if err = s.repo.DeleteSegmentFromUser(ctx, userSegmentToDel); err != nil {
			return err
		}
	}

	if len(segments.ToAdd) == 0 {
		return nil
	}

	currentSegments, err := s.repo.UsersSegments(ctx, segments.UserId)
	if err != nil {
		return err
	}

	for _, segment := range segments.ToAdd {
		for _, segm := range currentSegments {
			if segment.Name == segm.Name {
				return errors.Wrap(errTypes.ErrAlreadyExists,
					fmt.Sprintf("Segment %s already exists for user %d", segment.Name, segments.UserId))
			}
		}
	}

	idsToAdd, err := s.repo.SegmentsIdsByName(ctx, segments.ToAdd)
	if err != nil {
		return err
	}

	userSegmentToAdd := make([]entity.UserSegment, len(segments.ToAdd))
	for i, segment := range segments.ToAdd {
		userSegmentToAdd[i] = entity.UserSegment{
			UserId:    segments.UserId,
			SegmentId: idsToAdd[i],
			DeletedAt: segment.DeletedAt,
		}
	}

	err = s.repo.AddSegment(ctx, userSegmentToAdd)

	return err
}

func (s *UserService) Operations(ctx context.Context, usersOperations entity.UserOperations) ([]entity.Operation, error) {

	operations, err := s.repo.Operations(ctx, usersOperations)
	if err != nil {
		return []entity.Operation{}, err
	}

	return operations, nil
}
