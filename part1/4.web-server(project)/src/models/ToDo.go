package models

type ToDo struct {
	Id      string
	Content string
	IsDone  bool
}

type RawToDo struct {
	Content string
	IsDone  bool
}
