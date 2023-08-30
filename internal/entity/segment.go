package entity

type Segment struct {
	Id      int    `json:"id"`
	Name    string `json:"name" validate:"required"`
	Percent int    `json:"percent"`
}
