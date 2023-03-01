package pg

import (
	"antoine29/go/web-server/src/models"
	"fmt"
	"log"
 	"errors"

	config "antoine29/go/web-server/src"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getPGConn() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		config.EnvVarsDict["PG_HOST"],
		config.EnvVarsDict["PG_PORT"],
		config.EnvVarsDict["PG_USER"],
		config.EnvVarsDict["PG_PASSWORD"],
		config.EnvVarsDict["PG_DBNAME"],
		config.EnvVarsDict["PG_SCHEMA"],
	)

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

	defer closeDbConn(conn)

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

	defer closeDbConn(conn)

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

	defer closeDbConn(conn)

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

func PutToDo(todo models.ToDo) (*models.ToDo, error) {
	conn := getPGConn()
	if conn == nil {
		return nil, errors.New("error: Error getting pg connection")
	}

	defer closeDbConn(conn)

	result := conn.Create(&todo)
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

	defer closeDbConn(conn)

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

	defer closeDbConn(conn)

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

func IsDBHealthy() (bool, *string) {
	config.LoadEnvVarsDict(true)
	conn := getPGConn()
	if conn == nil {
		error := "Error connecting to DB"
		return false, &error
	}

	defer closeDbConn(conn)

	var todo models.ToDo
	result := conn.Take(&todo)
	if result.Error == nil || result.Error.Error() == "record not found" {
		return true, nil
	}

	errorStr := result.Error.Error()
	return false, &errorStr
}

func closeDbConn(db *gorm.DB) {
	dbInstance, _ := db.DB()
	_ = dbInstance.Close()
}
