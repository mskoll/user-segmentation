package entity

type Segment struct {
	ID      int    `json:"id"`
	Name    string `json:"name" db:"name"`
	Percent int    `json:"percent"`
}
