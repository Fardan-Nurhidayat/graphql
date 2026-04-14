package models

import (
	"fmt"
	"sync"
	"time"
)

type Todo struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	Done      bool   `json:"done"`
	CreatedAt string `json:"createdAt"`
}

// TodoStore Manage the todos in memory
type TodoStore struct {
	mu     sync.Mutex
	todos  map[string]*Todo // deklarasi map dengan pointer ke Todo bertujuan untuk menghindari penyalinan data yang besar ketika mengakses atau memodifikasi todo, sehingga meningkatkan performa aplikasi
	nextID int
}

func NewTodoStore() *TodoStore {
	return &TodoStore{
		todos:  make(map[string]*Todo),
		nextID: 1,
	}
}

// GetTodos returns all todos
func (s *TodoStore) GetAllTodos() []*Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	todos := make([]*Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		todos = append(todos, todo)
	}
	return todos
}

func (s *TodoStore) GetById(id string) *Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.todos[id]
}

func (s *TodoStore) GetByStatus(done bool) []*Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	var filtered []*Todo
	for _, todo := range s.todos {
		if todo.Done == done {
			filtered = append(filtered, todo)
		}
	}
	return filtered
}

// Create a new todo
func (s *TodoStore) Create(text string) *Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	todo := &Todo{
		ID:        fmt.Sprintf("%d", s.nextID),
		Text:      text,
		Done:      false,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	s.nextID++
	s.todos[todo.ID] = todo
	return todo
}

// Update an existing todo
func (s *TodoStore) Update(id string, text *string, done *bool) *Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	todo := s.todos[id]
	if todo == nil {
		return nil
	}

	if text != nil {
		todo.Text = *text
	}
	if done != nil {
		todo.Done = *done
	}
	return todo
}

func (s *TodoStore) Delete(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.todos[id]; !exists {
		return false
	}
	delete(s.todos, id)
	return true
}

// Toggle switches the done status of a todo

func (s *TodoStore) Toggle(id string) *Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	todo := s.todos[id]
	if todo == nil {
		return nil
	}
	todo.Done = !todo.Done
	return todo
}
