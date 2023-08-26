package entity

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" validate:"required"`
}

type SegmentList struct {
	User     User      `json:"user"`
	Segments []Segment `json:"segments"`
}

type AddDelSegments struct {
	Id    int      `json:"id"`
	ToAdd []string `json:"to_add"`
	ToDel []string `json:"to_del"`
}
