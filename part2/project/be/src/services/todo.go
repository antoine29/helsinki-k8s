package services

import (
	"antoine29/go/web-server/src/models"
	// "errors"
	"fmt"
	"log"

	config "antoine29/go/web-server/src"
  dao "antoine29/go/web-server/src/dao/pg"
	"github.com/google/uuid"
)

func CreateToDo(content string) *models.ToDo {
	newId := uuid.New().String()
	newToDo := models.ToDo{
		Id:      newId,
		Content: content,
		IsDone:  false,
	}

	result := dao.InsertToDo()
	if result.Error != nil {
		log.Println("Error creating todo")
		log.Println(result.Error)
		return nil
	}

	return &newToDo
}

func UpsertToDo(todo models.ToDo) (*models.ToDo, error) {
	// tries to update todo
	if _, err := GetToDo(todo.Id); err == nil {
		updatedTodoPointer, err := UpdateToDo(todo.Id, todo)
    return *updatedTodoPointer, nil
	}

	// tries to create todo
	createdToDo := conn.Create(&todo)
	if createdToDo.Error == nil {
		return &todo, nil
	} else {
		return nil, createdToDo.Error
	}
	}
}

func UpdateToDo(id string, todo models.ToDo) (*models.ToDo, error) {
	conn, err := getPGConn()
	if err != nil {
		return nil, err
	}

	defer closeDbConn(conn)

	// var existingTodoPointer *models.ToDo
	existingTodoPointer, err := GetToDo(id)
	if err != nil {
		return nil, err
	}

	existingTodo := *existingTodoPointer
	updatedTodo := models.ToDo{
		Id:      existingTodo.Id,
		Content: todo.Content,
		IsDone:  todo.IsDone,
	}

	// fmt.Printf("%+v\n", result)
	if updatingResult := conn.Save(*&updatedTodo); updatingResult.Error != nil {
		return nil, updatingResult.Error
	}

	return existingTodoPointer, nil
}

func UpsertToDo(todo models.ToDo) (*models.ToDo, error) {
	// tries to update todo
  existingTodo := dao.GetToDo()
	if _, err := dao.UpsertToDo(todo); err == nil {
		updatedTodoPointer, err := UpdateToDo(todo.Id, todo)
    return *updatedTodoPointer, nil
	}

