package models

type Message struct {
	ID     string `json:"id"`
	ToDo   string `json:"todo"`
	Status string `json:"status"`
}
