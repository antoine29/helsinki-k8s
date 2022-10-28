package controllers

import (
	// dao "antoine29/go/web-server/src/dao/inMemory"
	dao "antoine29/go/web-server/src/dao/pg"
	"antoine29/go/web-server/src/models"
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
	todos := dao.GetToDo_s()
	c.IndentedJSON(http.StatusOK, todos)
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
	todo := dao.GetToDo(id)
	if todo == nil {
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, *todo)
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
		c.IndentedJSON(http.StatusUnprocessableEntity, err)
		return
	}

	if newToDo.Content == "" {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"error": "'content' field expected in the payload body"})
		return
	}

	createdToDo := dao.AddToDo(newToDo.Content)
	c.IndentedJSON(http.StatusCreated, createdToDo)
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
	deletedToDoPointer := dao.DeleteToDo(id)

	if deletedToDoPointer == nil {
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, *deletedToDoPointer)
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
		c.IndentedJSON(http.StatusUnprocessableEntity, err)
		return
	}

	updatedToDoPointer := dao.UpdateToDo(id, updatedBody)

	if updatedToDoPointer == nil {
		c.Status(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, *updatedToDoPointer)
	}
}

// @Summary Check API health
// @Schemes http
// @Description Check API health
// @Produce json
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "Healty"})
}
