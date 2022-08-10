package main

import (
	"antoine29/go/web-server/docs"
	controllers "antoine29/go/web-server/src/controllers"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary Check API health
// @Schemes http
// @Description Check API health
// @Produce json
// @Router /health [get]
func healthCheckController(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "Healty"})
}

func setupSwagger() gin.HandlerFunc {
	docs.SwaggerInfo.Title = "ToDo API"
	docs.SwaggerInfo.Description = "A GO API REST server to manage ToDo's"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Version = "1.0"

	return ginSwagger.WrapHandler(swaggerfiles.Handler)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/todos", controllers.GetTodos)
		api.GET("/todos/:id", controllers.GetTodo)
		api.POST("/todos", controllers.PostTodo)
		api.PATCH("/todos/:id", controllers.UpdateTodo)
		api.DELETE("/todos/:id", controllers.DeleteTodo)

		api.GET("/health", healthCheckController)
	}

	r.GET("/swagger/*any", setupSwagger())
	return r
}

func main() {
	port := os.Getenv("GO_PORT")
	if port == "" {
		fmt.Println("Warning: 'GO_PORT' environment variable was not set, using 8080 as default.")
		port = "8080"
	}

	r := setupRouter()

	fmt.Printf("Go to: 'http://localhost:%s/swagger/index.html' to check Swagger API docs.\n", port)
	// Listen and Server in 0.0.0.0:{port}
	r.Run(fmt.Sprintf(":%s", port))
}
