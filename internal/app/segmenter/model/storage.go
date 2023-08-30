package model

type Segment struct {
	ID   uint   `json:"id"`
	Slug string `json:"slug"`
}

type UsersSegments struct {
	UserID    uint
	SegmentID uint
}
