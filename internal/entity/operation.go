package entity

import "time"

type Operation struct {
	ID            int       `json:"id"`
	UserId        int       `json:"user_id"`
	SegmentId     int       `json:"segment_id"`
	OperationType string    `json:"operation_type"`
	CreatedAt     time.Time `json:"created_at"`
}
