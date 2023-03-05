package services

import (
	"antoine29/go/web-server/src/models"
	"errors"

	dao "antoine29/go/web-server/src/dao/pg"

	"github.com/google/uuid"
)

func GetToDos() ([]models.ToDo, error) {
	todos, err := dao.GetToDos()
	if err != nil {
		return nil, err
	} else {
		return todos, nil
	}
}

func GetToDo(id string) (*models.ToDo, error) {
	todoPointer, err := dao.GetToDo(id)
	if err != nil {
		return nil, err
	} else {
		return todoPointer, nil
	}
}

func CreateToDo(content string) (*models.ToDo, error) {
	newId := uuid.New().String()
	newToDo := models.ToDo{
		Id:      newId,
		Content: content,
		IsDone:  false,
	}

	if err := isValidToDo(newToDo); err != nil {
		return nil, err
	}

	createdToDoPointer, err := dao.InsertToDo(newToDo)
	if err != nil {
		return nil, err
	}

	return createdToDoPointer, nil
}

func UpdateToDo(id string, todo models.ToDo) (*models.ToDo, error) {
	existingTodoPointer, err := dao.GetToDo(id)
	if err != nil {
		return nil, err
	}

	existingTodo := *existingTodoPointer
	updatedTodo := models.ToDo{
		Id:      existingTodo.Id,
		Content: todo.Content,
		IsDone:  todo.IsDone,
	}

	if err := isValidToDo(updatedTodo); err != nil {
		return nil, err
	}

	// fmt.Printf("%+v\n", result)
	if _, err := dao.UpsertToDo(&updatedTodo); err != nil {
		return nil, err
	}

	return &updatedTodo, nil
}

func UpsertToDo(todo models.ToDo) (*models.ToDo, error) {
	_, err := dao.GetToDo(todo.Id)
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}

	if err := isValidToDo(todo); err != nil {
		return nil, err
	}

	if err != nil && err.Error() == "record not found" {
		upsertedPointer, err := dao.UpsertToDo(&todo)
		if err != nil {
			return nil, err
		} else {
			return upsertedPointer, nil
		}
	}

	updatedPointer, err := UpdateToDo(todo.Id, todo)
	if err != nil {
		return nil, err
	} else {
		return updatedPointer, nil
	}
}

func DeleteToDo(id string) (bool, error) {
	if _, err := dao.DeleteToDo(id); err != nil {
		return false, err
	}

	return true, nil
}

func isValidToDo(todo models.ToDo) error {
	if todo.Id == "" {
		return errors.New("'Id' field cannot be empty")
	}

	if todo.Content == "" {
		return errors.New("'Content' field cannot be empty")
	}

	if len(todo.Content) > 140 {
		return errors.New("'Content' field cannot exced 140 chars long.")
	}

	return nil
}
