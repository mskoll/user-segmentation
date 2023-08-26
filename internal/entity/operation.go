package entity

import "time"

type Operation struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	SegmentId int       `json:"segment_id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}
