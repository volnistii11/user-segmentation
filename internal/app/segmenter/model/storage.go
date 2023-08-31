package model

type Segment struct {
	Slug string `json:"slug"`
}

type UsersSegments struct {
	UserID    uint
	SegmentID uint
}
