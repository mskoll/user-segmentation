package service

import (
	"context"
	"userSegmentation/internal/entity"
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

		toDel, err := s.repo.SegmentsIdsByName(ctx, segments.ToDel)
		if err != nil {
			return err
		}
		if err = s.repo.DeleteSegmentFromUser(ctx, segments.UserId, toDel); err != nil {
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

	segmentsWithoutTtl := make([]entity.SegmentToUser, 0, len(segments.ToAdd))
	segmentsWithTtl := make([]entity.SegmentToUser, 0, len(segments.ToAdd))

	for _, segment := range segments.ToAdd {
		for _, segm := range currentSegments {
			if segment.Name == segm.Name {
				// todo : fix err
				return err
			}
			if segment.Ttl.IsZero() {
				segmentsWithoutTtl = append(segmentsWithoutTtl, segment)
			} else {
				segmentsWithTtl = append(segmentsWithTtl, segment)
			}
		}
	}
	toAdd, err := s.repo.SegmentsIdsByName(ctx, segmentsWithoutTtl)
	if err != nil {
		return err
	}
	toAddWithTtl, err := s.repo.SegmentsIdsByName(ctx, segmentsWithTtl)
	if err != nil {
		return err
	}

	userSegment := make([]entity.UserSegment, len(toAdd))
	for i := range userSegment {
		userSegment[i] = entity.UserSegment{
			UserId:    segments.UserId,
			SegmentId: toAdd[i].Id,
		}
	}

	userSegmentWithTtl := make([]entity.UserSegment, len(toAdd))
	for i := range userSegmentWithTtl {
		userSegmentWithTtl[i] = entity.UserSegment{
			UserId:    segments.UserId,
			SegmentId: toAdd[i].Id,
			DeletedAt: toAddWithTtl[i].Ttl,
		}
	}

	if err = s.repo.AddSegment(ctx, segments.UserId, userSegment); err != nil {
		return err
	}
	if err = s.repo.AddSegmentWithTtl(ctx, segments.UserId, userSegmentWithTtl); err != nil {
		return err
	}

	return nil
}

func (s *UserService) Operations(ctx context.Context, usersOperations entity.UsersOperations) ([]entity.Operation, error) {

	operations, err := s.repo.Operations(ctx, usersOperations)
	if err != nil {
		return []entity.Operation{}, err
	}

	return operations, nil
}
