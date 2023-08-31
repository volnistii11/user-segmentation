package model

type Segment struct {
	Slug string `json:"slug"`
}

type User struct {
	ID uint `json:"user_id,string"`
}

type UsersSegments struct {
	UserID    uint
	SegmentID uint
}
