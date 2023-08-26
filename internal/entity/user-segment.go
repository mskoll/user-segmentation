package entity

type UserSegment struct {
	Id        int
	UserId    int `db:"user_id"`
	SegmentId int `db:"segment_id"`
}
