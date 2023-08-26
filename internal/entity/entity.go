package entity

type Segment struct {
	Id   int
	Name string
}

type History struct {
	UserId        int
	SegmentName   string
	Operation     string
	OperationTime string
}

type User struct {
	Id int `json:"id"`
}
