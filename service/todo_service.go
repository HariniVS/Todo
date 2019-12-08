package service

import (
	"errors"
	"github.com/todo/model"
)

type TodoService interface {
	GetTodo(id int) (model.Todo, error)
	GetAllTodos() ([]model.Todo, error)
}

type todoService struct {

}

func (service todoService) GetAllTodos() ([]model.Todo, error) {
	todoOne := model.Todo{1, "Book tickets", false}
	todoTwo := model.Todo{2, "Read book", true}
	return []model.Todo{todoOne, todoTwo}, nil
}

func (service todoService) GetTodo(taskId int) (model.Todo, error) {
	todos, err := service.GetAllTodos()
	if err != nil {
		return model.Todo{}, errors.New("Unable to fetch all Todos")
	}
	for _, todo := range todos {
		if todo.Id == taskId {
			return todo, nil
		}
	}
	return model.Todo{}, errors.New("Todo not found")
}

func NewTodoService() TodoService {
	return todoService{}
}

