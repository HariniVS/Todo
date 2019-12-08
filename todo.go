package main

import (
	"github.com/gin-gonic/gin"
	"github.com/todo/controller"
	"github.com/todo/service"
)

func main()  {
	router := gin.Default()

	todoService := service.NewTodoService()
	todoController := controller.NewTodoController(todoService)

	routerGroup := router.Group("/api/v1/todo")
	{
		routerGroup.GET("/", todoController.GetAllTodos)
		routerGroup.GET("/:id", todoController.GetTodo)
		routerGroup.POST("/", todoController.CreateATodo)
		routerGroup.PUT("/:id", todoController.UpdateATodo)
		routerGroup.DELETE("/:id", todoController.DeleteATodo)
	}
	router.Run()
}
