package entity

type Segment struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Expire *string `json:"expire"`
}

type History struct {
	UserId        int    `json:"user_id"`
	SegmentName   string `json:"segment_name"`
	Operation     string `json:"operation"`
	OperationTime string `json:"operation_time"`
}

type User struct {
	Id int `json:"id"`
}
