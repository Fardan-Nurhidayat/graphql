package graph

import (
	"context"
	"graphql/graph/model"
	"graphql/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodoCreateResolver(t *testing.T) {
	store := models.NewTodoStore()
	resolver := &Resolver{TodoStore: store}

	input := model.NewTodo{Text: "Learn C#"}

	todo, err := resolver.Mutation().CreateTodo(context.Background(), input)

	assert.NoError(t, err)
	assert.Equal(t, "Learn C#", todo.Text)
	assert.False(t, todo.Done)
	assert.NotEmpty(t, todo.ID)

}

func TestGetTodo(t *testing.T) {
	store := models.NewTodoStore()
	resolver := &Resolver{TodoStore: store}
	store.Create("Title 1")
	store.Create("Title 2")
	todos, err := resolver.Query().Todos(context.Background())

	assert.NoError(t, err)
	assert.Len(t, todos, 2)
}

func TestCreateTodo_Validation_Errors(t *testing.T) {
	store := models.NewTodoStore()
	resolver := &Resolver{TodoStore: store}

	input := model.NewTodo{Text: ""}

	todo, err := resolver.Mutation().CreateTodo(context.Background(), input)

	assert.Error(t, err)
	assert.Nil(t, todo)
	assert.Equal(t, "Validation failed for text: text cannot be empty", err.Error())

	input = model.NewTodo{Text: generateSlug(300)} // Assuming max length is 255

	todo, err = resolver.Mutation().CreateTodo(context.Background(), input)

	assert.Error(t, err)
	assert.Nil(t, todo)
	assert.Equal(t, "Validation failed for text: text cannot exceed 255 characters", err.Error())
}

func generateSlug(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = 'a' + byte(i%26)
	}
	return string(b)
}
