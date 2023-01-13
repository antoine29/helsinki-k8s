package pg

import (
	"antoine29/go/web-server/src/models"
  "log"
  "fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

// var dsn string = "host=localhost port=5432 user=postgres password=postgres dbname=postgres search_path=todo sslmode=disable"
var dsn string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=disable",
	os.Getenv("PG_HOST"),
	os.Getenv("PG_PORT"),
	os.Getenv("PG_USER"),
	os.Getenv("PG_PASSWORD"),
	os.Getenv("PG_DBNAME"),
	os.Getenv("PG_SCHEMA"),
)

func getPGConn() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error connecting to PG DB")
		log.Println(err)
		return nil
	}

	return db
}

func GetToDo_s() []models.ToDo {
	conn := getPGConn()
	if conn == nil {
		return nil
	}

	todos := []models.ToDo{}
	result := conn.Find(&todos)
	if result.Error != nil {
		log.Println("Error getting pg todos")
		log.Println(result.Error)
		return nil
	}

	return todos
}

func GetToDo(id string) *models.ToDo {
	conn := getPGConn()
	if conn == nil {
		return nil
	}

	todo := models.ToDo{}
	result := conn.First(&todo, "id = ?", id)
	if result.Error != nil {
		log.Println("Error getting todo")
		log.Println(result.Error)
		return nil
	}

	return &todo
}

func AddToDo(content string) *models.ToDo {
	conn := getPGConn()
	if conn == nil {
		return nil
	}

	newId := uuid.New().String()
	newToDo := models.ToDo{
		Id:      newId,
		Content: content,
		IsDone:  false,
	}

	result := conn.Create(&newToDo)
	if result.Error != nil {
		log.Println("Error creating todo")
		log.Println(result.Error)
		return nil
	}

	return &newToDo
}

func DeleteToDo(id string) *models.ToDo {
	conn := getPGConn()
	if conn == nil {
		return nil
	}

	todo2delete := models.ToDo{Id: id}
	result := conn.Delete(&todo2delete)
	// fmt.Printf("%+v\n", result)
	if result.RowsAffected == 0 {
		log.Println("Error deleting todo")
		return nil
	}

	return &todo2delete
}

func UpdateToDo(id string, todo models.ToDo) *models.ToDo {
	conn := getPGConn()
	if conn == nil {
		return nil
	}

	dbTodo := GetToDo(id)
	if dbTodo == nil {
		log.Printf("Todo with id: %s not found", id)
		return nil
	}

	dbTodo.Content = todo.Content
	dbTodo.IsDone = todo.IsDone

	result := conn.Save(dbTodo)
	// fmt.Printf("%+v\n", result)
	if result.RowsAffected == 0 {
		log.Println("Error updating todo")
		return nil
	}

	return dbTodo
}

