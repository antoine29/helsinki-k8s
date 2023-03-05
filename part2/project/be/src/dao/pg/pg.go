package pg

import (
	"antoine29/go/web-server/src/models"
	"fmt"

	config "antoine29/go/web-server/src"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getPGConn() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		config.EnvVarsDict["PG_HOST"],
		config.EnvVarsDict["PG_PORT"],
		config.EnvVarsDict["PG_USER"],
		config.EnvVarsDict["PG_PASSWORD"],
		config.EnvVarsDict["PG_DBNAME"],
		config.EnvVarsDict["PG_SCHEMA"],
	)

	var (
		db  *gorm.DB
		err error
	)

	if db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		return nil, err
	}

	return db, nil
}

func GetToDos() ([]models.ToDo, error) {
	conn, err := getPGConn()
	if err != nil {
		return nil, err
	}

	defer closeDbConn(conn)

	todos := []models.ToDo{}
	if result := conn.Find(&todos); result.Error != nil {
		return nil, result.Error
	}

	return todos, nil
}

func GetToDo(id string) (*models.ToDo, error) {
	conn, err := getPGConn()
	if err != nil {
		return nil, err
	}

	defer closeDbConn(conn)

	todo := models.ToDo{}
	result := conn.First(&todo, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}

func InsertToDo(todo models.ToDo) (*models.ToDo, error) {
	conn, err := getPGConn()
	if err != nil {
		return nil, err
	}

	defer closeDbConn(conn)

	result := conn.Create(&todo)
	if result.Error != nil {
		return nil, err
	}

	return &todo, nil
}

func UpsertToDo(todo *models.ToDo) (*models.ToDo, error) {
	conn, err := getPGConn()
	if err != nil {
		return nil, err
	}

	defer closeDbConn(conn)

	if upsertResult := conn.Save(*todo); upsertResult.Error != nil {
		return nil, upsertResult.Error
	}

	return todo, nil
}

func DeleteToDo(id string) (*models.ToDo, error) {
	conn, err := getPGConn()
	if err != nil {
		return nil, err
	}

	defer closeDbConn(conn)

	todo2delete := models.ToDo{Id: id}
	// fmt.Printf("%+v\n", result)
	if result := conn.Delete(&todo2delete); result.Error != nil {
		return nil, result.Error
	}

	return &todo2delete, nil
}

func IsDBHealthy() error {
	config.LoadEnvVarsDict(true)
	conn, err := getPGConn()
	if err != nil {
		return err
	}

	defer closeDbConn(conn)

	var todo models.ToDo
	result := conn.Take(&todo)
	if result.Error == nil || result.Error.Error() == "record not found" {
		return nil
	}

	return result.Error
}

func closeDbConn(db *gorm.DB) {
	dbInstance, _ := db.DB()
	_ = dbInstance.Close()
}
