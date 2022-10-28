package inMemory

import (
	"antoine29/go/web-server/src/models"

	"github.com/google/uuid"

	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var todo0 = models.ToDo{
	Id:      "0",
	Content: "todo0",
	IsDone:  false,
}

var todo1 = models.ToDo{
	Id:      "1",
	Content: "todo1",
	IsDone:  false,
}

var todo2 = models.ToDo{
	Id:      "2",
	Content: "todo2",
	IsDone:  true,
}

var todo_s = map[string]*models.ToDo{
	"0": &todo0,
	"1": &todo1,
	"2": &todo2,
}

func initializeMap() {
	todo_s = make(map[string]*models.ToDo)
}

func GetToDo_s() []models.ToDo {
	todo_s_array := make([]models.ToDo, 0, len(todo_s))
	for _, value := range todo_s {
		todo_s_array = append(todo_s_array, *value)
	}

	return todo_s_array
}

func GetToDo(id string) *models.ToDo {
	value, ok := todo_s[id]

	if ok {
		return value
	} else {
		return nil
	}
}

func AddToDo(content string) models.ToDo {
	if len(todo_s) == 0 {
		initializeMap()
	}

	newId := uuid.New().String()
	newToDo := models.ToDo{
		Id:      newId,
		Content: content,
		IsDone:  false,
	}

	todo_s[newId] = &newToDo
	createdToDo := *todo_s[newId]
	return createdToDo
}

func DeleteToDo(id string) *models.ToDo {
	todoPointer := GetToDo(id)
	if todoPointer != nil {
		delete(todo_s, id)
		return todoPointer
	} else {
		return nil
	}
}

func UpdateToDo(id string, todo models.ToDo) *models.ToDo {
	todoPointer := GetToDo(id)
	if todoPointer != nil {
		todo_s[id].Content = todo.Content
		todo_s[id].IsDone = todo.IsDone
		return todoPointer
	} else {
		return nil
	}
}

func Test() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("error")
		return
	}

	todo := models.ToDo{}
	db.First(&todo)
	fmt.Println(todo)

}
