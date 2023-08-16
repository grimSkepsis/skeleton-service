package graph

import (
	"context"
	dbmodel "skeleton-service/database/model"
	"skeleton-service/graph/model"
	"skeleton-service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestMutationResolver_CreateTodo(t *testing.T) {
	logger, _ := zap.NewProduction()

	defer logger.Sync()
	// Create a new mock TodoManager
	mockTodoManager := &mocks.TodoManager{}

	// Create a new resolver with the mock TodoManager
	resolver := NewResolver(logger, mockTodoManager)

	// Define the input for the CreateTodo mutation
	input := model.NewTodo{Text: "Test Todo", UserID: "testuser"}

	// Define the expected output from the CreateTodo mutation
	expectedOutput := &model.Todo{ID: "testid", Text: "Test Todo", Done: false, User: &model.User{ID: "testuser", Name: "Test User"}}

	// Set up the mock TodoManager's Create method to return the expected output
	mockTodoManager.On("Create", mock.AnythingOfType("*dbmodel.Todo")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*dbmodel.Todo)
		arg.ID = "testid"
	}).Once()

	// Call the CreateTodo mutation with the input
	output, err := resolver.Mutation().CreateTodo(context.Background(), input)

	// Assert that the output matches the expected output
	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, output)

	// Assert that the mock TodoManager's Create method was called with the correct argument
	mockTodoManager.AssertCalled(t, "Create", mock.AnythingOfType("*dbmodel.Todo"))
}
