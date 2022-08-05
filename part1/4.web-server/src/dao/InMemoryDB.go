package src

import (
	"strconv"
	"antoine29/go/web-server/src/models"
)

var todos = []models.ToDo{}

func GetToDo_s() []models.ToDo {
	return todos 
}

func AddToDo(todo models.ToDo) models.ToDo {
	createdToDo := models.ToDo{
		strconv.Itoa(len(todos)),
		todo.Title,
		false,
	}
	
	todos = append(todos, createdToDo)

	return createdToDo
}
