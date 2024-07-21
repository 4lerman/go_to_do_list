package types

import "time"

type ToDoStore interface {
	CreateToDo(ToDo) (string, error)
	GetToDos(string) ([]ToDo, error)
	DeleteToDo(string) error
	ChangeToDone(string) error
	UpdateToDo(string, ToDo) error
}

type ToDo struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	ActiveAt time.Time `json:"activeAt"`
}

type CreateToDoPayload struct {
	Title    string    `json:"title" validate:"required" example:"title"`
	ActiveAt time.Time `json:"activeAt" validate:"required" example:"2024-07-22T00:55:37.481555+05:00"`
}

type UpdateToDoPayload struct {
	Title    string    `json:"title" validate:"required,max=200" example:"updated title"`
	ActiveAt time.Time `json:"activeAt" validate:"required" example:"2024-07-22T00:55:37.481555+05:00"`
}
