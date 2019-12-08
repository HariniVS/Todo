package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/todo/service"
	"strconv"
)

type TodoController interface {
	GetTodo(ctx *gin.Context)
	GetAllTodos(ctx *gin.Context)
	CreateATodo(ctx *gin.Context)
	UpdateATodo(ctx *gin.Context)
	DeleteATodo(ctx *gin.Context)
}

type todoController struct {
	todoService service.TodoService
}

func (controller todoController) GetTodo(ctx *gin.Context) {
	var param string
	param = ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(500, err)
	}
	todos, err := controller.todoService.GetTodo(id)
	ctx.JSON(200, todos)
}

func (controller todoController) GetAllTodos(ctx *gin.Context) {
	todos, err := controller.todoService.GetAllTodos()
	if err != nil {
		ctx.JSON(500, err)
	}
	ctx.JSON(200, todos)
}

func (controller todoController) CreateATodo(ctx *gin.Context) {
	panic("implement me")
}

func (controller todoController) UpdateATodo(ctx *gin.Context) {
	panic("implement me")
}

func (controller todoController) DeleteATodo(ctx *gin.Context) {
	panic("implement me")
}

func NewTodoController(todoService service.TodoService) TodoController {
	return todoController{todoService:todoService}
}