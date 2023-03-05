package controllers

import (
	"antoine29/go/web-server/src/models"
	todoService "antoine29/go/web-server/src/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get ToDo's list
// @Schemes http
// @Description Get ToDo's list
// @Produce json
// @Success 200 {array} models.ToDo
// @Router /todos [get]
func GetTodos(c *gin.Context) {
	if todos, err := todoService.GetToDos(); err != nil {
		fmt.Printf("%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, todos)
	}
}

// @Summary Get ToDo by id
// @Schemes http
// @Description Get ToDo by id
// @Produce json
// @Param	id	path	string	true	"ToDo Id"
// @Success 200 {object} models.ToDo
// @Router /todos/{id} [get]
func GetTodo(c *gin.Context) {
	id := c.Param("id")
	if todoPointer, err := todoService.GetToDo(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, *todoPointer)
	}
}

// @Summary Post ToDo
// @Schemes http
// @Description Post ToDo
// @Produce json
// @Param	ToDo	body	models.RawToDo	true	"ToDo json payload"
// @Success 200 {object} models.ToDo
// @Router /todos [post]
func PostTodo(c *gin.Context) {
	var newToDo models.ToDo
	if err := c.BindJSON(&newToDo); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if createdToDo, err := todoService.CreateToDo(newToDo.Content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.IndentedJSON(http.StatusCreated, createdToDo)
	}
}

// @Summary Delete ToDo
// @Schemes http
// @Description Delete ToDo
// @Produce json
// @Param	id	path	string	true	"Id of the ToDo to delete"
// @Success 200 {object} models.ToDo
// @Router /todos/{id} [delete]
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if deletedResult, err := todoService.DeleteToDo(id); err != nil && deletedResult == false {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.Status(http.StatusOK)
	}
}

// @Summary Update ToDo
// @Schemes http
// @Description Update ToDo
// @Produce json
// @Param	id	path	string	true	"Id of the ToDo to update"
// @Param	ToDo	body	models.RawToDo	true	"ToDo updated body"
// @Success 200 {object} models.ToDo
// @Router /todos/{id} [patch]
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var updatedBody models.ToDo
	if err := c.BindJSON(&updatedBody); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if updatedToDoPointer, err := todoService.UpdateToDo(id, updatedBody); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, *updatedToDoPointer)
	}
}

// @Summary Put ToDo
// @Schemes http
// @Description Put (update or create) a ToDo
// @Produce json
// @Param	ToDo	body	models.ToDo	true	"ToDo body"
// @Success 200 {object} models.ToDo
// @Router /todos [put]
func PutTodo(c *gin.Context) {
	var todo models.ToDo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if upsertedPointer, err := todoService.UpsertToDo(todo); err == nil {
		c.JSON(http.StatusCreated, upsertedPointer)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
