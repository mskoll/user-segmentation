package entity

import "time"

type UserSegment struct {
	Id        int
	UserId    int `db:"user_id"`
	SegmentId int `db:"segment_id"`
	CreatedAt time.Time
	DeletedAt time.Time
}

type Operation struct {
	UserId      int       `json:"user_id"`
	SegmentName string    `json:"segment_name"`
	Operation   string    `json:"operation"`
	Datetime    time.Time `json:"datetime"`
}
