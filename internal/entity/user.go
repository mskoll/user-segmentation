package entity

import "time"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" validate:"required"`
}

type SegmentList struct {
	User     User      `json:"user"`
	Segments []Segment `json:"segments"`
}

type AddDelSegments struct {
	UserId int             `json:"user_id"`
	ToAdd  []SegmentToUser `json:"to_add"`
	ToDel  []SegmentToUser `json:"to_del"`
}

type SegmentToUser struct {
	Name      string     `json:"name"`
	DeletedAt *time.Time `json:"ttl" db:"deleted_at"`
}

type UserOperations struct {
	UserId int `json:"user_id"`
	Month  int `json:"month"`
	Year   int `json:"year"`
}
