package entity

type Segment struct {
	Id   int
	Name string
}

type History struct {
	UserId        int
	Segment_name  string
	Operation     string
	OperationTime string
}

type User struct {
	Id int
}
