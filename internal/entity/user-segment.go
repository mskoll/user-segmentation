package entity

import "time"

type UserSegment struct {
	Id        int
	UserId    int `db:"user_id"`
	SegmentId int `db:"segment_id"`
	CreatedAt time.Time
	DeletedAt time.Time
}
