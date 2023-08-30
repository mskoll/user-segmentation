package entity

import "time"

type User struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username" validate:"required"`
}

type SegmentList struct {
	User     User      `json:"user"`
	Segments []Segment `json:"segments"`
}

type AddDelSegments struct {
	UserId int             `json:"user_id" validate:"required"`
	ToAdd  []SegmentToUser `json:"to_add,omitempty"`
	ToDel  []SegmentToUser `json:"to_del,omitempty"`
}

type SegmentToUser struct {
	Name      string     `json:"name" validate:"required"`
	DeletedAt *time.Time `json:"ttl,omitempty" db:"deleted_at"`
}

type UserOperations struct {
	UserId int `json:"user_id" validate:"required"`
	Month  int `json:"month" validate:"required"`
	Year   int `json:"year" validate:"required"`
}
