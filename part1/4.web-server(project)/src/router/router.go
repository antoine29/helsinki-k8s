package router

import (
	"antoine29/go/web-server/docs"
	"antoine29/go/web-server/src/controllers"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupSwagger() gin.HandlerFunc {
	docs.SwaggerInfo.Title = "ToDo API"
	docs.SwaggerInfo.Description = "A GO API REST server to manage ToDo's"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Version = "1.0"

	return ginSwagger.WrapHandler(swaggerfiles.Handler)
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./todo-fe/dist", true)))

	r.GET("/image/:name", controllers.GetImage)
	r.GET("/swagger/*any", setupSwagger())

	api := r.Group("/api")
	{
		api.GET("/todos", controllers.GetTodos)
		api.GET("/todos/:id", controllers.GetTodo)
		api.POST("/todos", controllers.PostTodo)
		api.PATCH("/todos/:id", controllers.UpdateTodo)
		api.DELETE("/todos/:id", controllers.DeleteTodo)

		api.GET("/health", controllers.HealthCheck)
	}

	return r
}
