package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/todo/model"
	"testing"
)

type TodoServiceSuite struct {
	suite.Suite
	todoService TodoService
}

func (suite *TodoServiceSuite) SetupTest()  {
	suite.todoService = NewTodoService()
}

func (suite *TodoServiceSuite) TestShouldGetTodoById() {
	todo := model.Todo{
		Id:        1,
		Task:      "Book tickets",
		Completed: false,
	}

	actualTodo, err := suite.todoService.GetTodo(todo.Id)
	if err != nil {
		assert.Error(suite.T(), err)
	}
	assert.Equal(suite.T(), todo, actualTodo)
}

func (suite *TodoServiceSuite) TestShouldGetAllTodos() {
	todoOne := model.Todo{
		Id:        1,
		Task:      "Book tickets",
		Completed: false,
	}

	todoTwo := model.Todo{
		Id:        2,
		Task:      "Book tickets",
		Completed: false,
	}

	expectedTodos := []model.Todo{todoOne, todoTwo}
	actualTodos, _ := suite.todoService.GetAllTodos()

	assert.Equal(suite.T(), 2 , len(actualTodos))
	assert.Equal(suite.T(), expectedTodos[0].Id, 1)
	assert.Equal(suite.T(), expectedTodos[0].Task, actualTodos[0].Task)
	assert.Equal(suite.T(), expectedTodos[0].Completed, actualTodos[0].Completed)

	assert.Equal(suite.T(), expectedTodos[1].Id, 2)
	assert.Equal(suite.T(), expectedTodos[0].Task, actualTodos[0].Task)
	assert.Equal(suite.T(), expectedTodos[0].Completed, actualTodos[0].Completed)
}

func TestTodoServiceSuite(t *testing.T) {
	suite.Run(t, new(TodoServiceSuite))
}
