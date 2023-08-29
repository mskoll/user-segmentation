package entity

import "time"

type UserSegment struct {
	UserId    int `db:"user_id"`
	SegmentId int `db:"segment_id"`
	CreatedAt time.Time
	DeletedAt *time.Time `db:"deleted_at"`
}

type Operation struct {
	UserId      int       `json:"user_id" db:"user_id"`
	SegmentName string    `json:"segment_name" db:"segment_name"`
	Operation   string    `json:"operation"`
	Datetime    time.Time `json:"datetime"`
}
