package models

type ToDo struct {
	Id      string `gorm:"primaryKey"`
	Content string
	IsDone  bool
}

type RawToDo struct {
	Content string
	IsDone  bool
}

// implementing iface to set the pg table name
type Tabler interface {
	TableName() string
}

func (ToDo) TableName() string {
	return "todos"
}
