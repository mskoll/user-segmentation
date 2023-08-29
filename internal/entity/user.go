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
	UserId int             `json:"user-id"`
	ToAdd  []SegmentToUser `json:"to_add"`
	ToDel  []SegmentToUser `json:"to_del"`
}

type SegmentToUser struct {
	Id   int       `json:"id"`
	Name string    `json:"name"`
	Ttl  time.Time `json:"ttl"`
}

type UsersOperations struct {
	Id    int `json:"id"`
	Month int `json:"month"`
	Year  int `json:"year"`
}
