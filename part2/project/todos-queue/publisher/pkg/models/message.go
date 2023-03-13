package models

type TodoMessage struct {
	ToDo   ToDo   `json:"todo"`
	Status string `json:"status"`
}

type ToDo struct {
	Id      string
	Content string
	IsDone  bool
}
