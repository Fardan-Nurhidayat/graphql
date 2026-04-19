package graph

import (
	"graphql/graph/model"
	"graphql/models"
)

func convertTodo(todo *models.Todo) *model.Todo {
	if todo == nil {
		return nil
	}
	return &model.Todo{
		ID:        todo.ID,
		Text:      todo.Text,
		Done:      todo.Done,
		CreatedAt: todo.CreatedAt,
	}
}

func convertTodos(todos []*models.Todo) []*model.Todo {
	result := make([]*model.Todo, 0, len(todos))
	for _, todo := range todos {
		result = append(result, convertTodo(todo))
	}
	return result
}
