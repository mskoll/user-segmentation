package entity

type Segment struct {
	Id      int    `json:"id,omitempty"`
	Name    string `json:"name" validate:"required"`
	Percent int    `json:"percent,omitempty"`
}
