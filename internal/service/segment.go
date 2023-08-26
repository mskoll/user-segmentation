package service

import "userSegmentation/internal/repo"

type SegmentService struct {
	repo repo.Segment
}

func NewSegmentService(repo repo.Segment) *SegmentService {
	return &SegmentService{repo: repo}
}
