package src

import (
	"antoine29/go/web-server/src/models"

	"github.com/google/uuid"
)

var todo_s map[string]*models.ToDo

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
