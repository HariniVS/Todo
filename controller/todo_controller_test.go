package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/todo/mocks"
	"github.com/todo/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TodoControllerSuite struct {
	suite.Suite
	recorder *httptest.ResponseRecorder
	context *gin.Context
	mockController *gomock.Controller
	mockTodoService *mocks.MockTodoService
	todoController TodoController
}

func (suite *TodoControllerSuite) SetupTest()  {
	suite.recorder = httptest.NewRecorder()
	suite.mockController = gomock.NewController(suite.T())
	suite.mockTodoService = mocks.NewMockTodoService(suite.mockController)
	suite.todoController = NewTodoController(suite.mockTodoService)
	suite.context, _ = gin.CreateTestContext(suite.recorder)
}

func (suite *TodoControllerSuite) TearDownTest()  {
	suite.mockController.Finish()
}

func (suite *TodoControllerSuite) TestShouldGetAllTodos() {
	expectedTodos := []model.Todo{{1, "Book tickets", false},{2, "Read book", false}}
	suite.context.Request = httptest.NewRequest("GET", "/api/v1/todo/", nil)
	suite.mockTodoService.EXPECT().GetAllTodos().Return(expectedTodos, nil).Times(1)

	suite.todoController.GetAllTodos(suite.context)


	type actualTodoResponse struct {
		Id        int	`json:"id"`
		Task      string	`json:"task"`
		Completed bool	`json:"completed"`
	}

	var actualTodos []actualTodoResponse
	err := json.Unmarshal(suite.recorder.Body.Bytes(), &actualTodos)

	if err != nil {
		panic(fmt.Sprint("Unable to get all Todos"))
	}

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)
	assert.Equal(suite.T(), len(expectedTodos), len(actualTodos))

	assert.Equal(suite.T(), expectedTodos[0].Id, actualTodos[0].Id)
	assert.Equal(suite.T(), expectedTodos[0].Task, actualTodos[0].Task)
	assert.Equal(suite.T(), expectedTodos[0].Completed, actualTodos[0].Completed)

	assert.Equal(suite.T(), expectedTodos[1].Id, actualTodos[1].Id)
	assert.Equal(suite.T(), expectedTodos[1].Task, actualTodos[1].Task)
	assert.Equal(suite.T(), expectedTodos[1].Completed, actualTodos[1].Completed)
}

func (suite *TodoControllerSuite) TestShouldGetATodo() {
	suite.context.Params = gin.Params{
		gin.Param{
			Key:   "id",
			Value: "2",
		},
	}
	suite.context.Request = httptest.NewRequest("GET", "/api/v1/todo/2", nil)

	expectedTodo := model.Todo{2, "Read book", false}
	suite.mockTodoService.EXPECT().GetTodo(2).Return(expectedTodo, nil).Times(1)

	suite.todoController.GetTodo(suite.context)

	type actualTodoResponse struct {
		Id        int	`json:"id"`
		Task      string	`json:"task"`
		Completed bool	`json:"completed"`
	}

	var actualTodo actualTodoResponse
	err := json.Unmarshal(suite.recorder.Body.Bytes(), &actualTodo)

	if err != nil {
		panic(fmt.Sprint("Unable to get todo for requested id"))
	}

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)

	assert.Equal(suite.T(), expectedTodo.Id, actualTodo.Id)
	assert.Equal(suite.T(), expectedTodo.Task, actualTodo.Task)
	assert.Equal(suite.T(), expectedTodo.Completed, actualTodo.Completed)
}

func TestTodoControllerSuite(t *testing.T) {
	suite.Run(t, new(TodoControllerSuite))
}
