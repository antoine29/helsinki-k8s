package services

import (
	config "antoine29/go/web-server/src"
	"antoine29/go/web-server/src/models"
	"encoding/json"
	"errors"
	"log"

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

	if err := sendTodoQueueMessage(createdToDoPointer, "created"); err != nil {
		log.Println("Error sending created todo to publisher.", err.Error())
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

	if err := sendTodoQueueMessage(&updatedTodo, "updated"); err != nil {
		log.Println("Error sending updated todo to publisher.", err.Error())
	}

	return &updatedTodo, nil
}

func UpsertToDo(id string, todo models.ToDo) (*models.ToDo, error) {
	todo.Id = id
	if err := isValidToDo(todo); err != nil {
		return nil, err
	}

	upsertedTodoPointer, err := dao.UpsertToDo(&todo)
	if err != nil {
		return nil, err
	} else {

		if err := sendTodoQueueMessage(upsertedTodoPointer, "upserted"); err != nil {
			log.Println("Error sending updated todo to publisher.", err.Error())
		}

		return upsertedTodoPointer, nil
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

func sendTodoQueueMessage(todoPointer *models.ToDo, status string) error {
	publisherUrl, isPublisherUrlEnvVarSet := config.EnvVarsDict["QUEUE_PUBLISHER_URL"]
	if !isPublisherUrlEnvVarSet {
		return errors.New("QUEUE_PUBLISHER_URL env var not set")
	}

	todoMessage := models.TodoMessage{
		ToDo:   *todoPointer,
		Status: status,
	}

	createdTodoJsonMessage, err := json.Marshal(todoMessage)
	if err != nil {
		return err
	}

	if err := HttpPost(publisherUrl, createdTodoJsonMessage); err != nil {
		return err
	}

	return nil
}
