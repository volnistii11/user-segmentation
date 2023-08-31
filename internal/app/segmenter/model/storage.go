package model

type Segment struct {
	Slug string `json:"slug"`
}

type User struct {
	ID uint `json:"user_id,string"`
}

type UserSegments struct {
	UserID           uint     `json:"user_id,string"`
	SegmentsToAdd    []string `json:"add_segments"`
	SegmentsToDelete []string `json:"delete_segments"`
}
