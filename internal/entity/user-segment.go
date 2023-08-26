package entity

type UserSegment struct {
	Id        int `json:"id"`
	UserId    int `json:"user_id" db:"user_id"`
	SegmentId int `json:"segment_id" db:"segment_id"`
	//TimeToLive time.Time `json:"time_to_live"`
}
