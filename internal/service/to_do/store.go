package to_do

import (
	"errors"
	"sync"
	"time"

	"github.com/4lerman/go_to_do_list/types"
	"github.com/google/uuid"
)

type Store struct {
	db *sync.Map
}

func NewStore(db *sync.Map) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateToDo(toDo types.ToDo) (string, error) {

	exists := false
	s.db.Range(func(key, value interface{}) bool {
		if existingToDo, ok := value.(types.ToDo); ok {
			if existingToDo.Title == toDo.Title {
				exists = true
				return false
			}
		}
		return true
	})

	if exists {
		return "", errors.New("todo task with the same title already exists")
	}

	id := uuid.New().String()
	toDo.ID = id

	s.db.Store(id, toDo)

	return id, nil
}

func (s *Store) GetToDos(status string) ([]types.ToDo, error) {
	toDos := []types.ToDo{}
	now := time.Now()

	s.db.Range(func(key, value interface{}) bool {
		if toDo, ok := value.(types.ToDo); ok {
			switch status {
			case "active":
				if toDo.ActiveAt.After(now) {
					toDos = append(toDos, toDo)
				}
			case "done":
				if !toDo.ActiveAt.After(now) {
					toDos = append(toDos, toDo)
				}
			default:
				return false
			}
		}
		return true
	})

	if status != "active" && status != "done" {
		return nil, errors.New("invalid status")
	}

	return toDos, nil
}

func (s *Store) DeleteToDo(id string) error {
	if _, ok := s.db.Load(id); !ok {
		return errors.New("given task does not exist")
	}

	s.db.Delete(id)
	return nil
}

func (s *Store) UpdateToDo(id string, payload types.ToDo) error {
	task, ok := s.db.Load(id)
	if !ok {
		return errors.New("given task does not exist")
	}

	toDo, ok := task.(types.ToDo)
	if !ok {
		return errors.New("task is not of type ToDo")
	}

	if len(payload.Title) > 200 {
		return errors.New("title should be less than 200 characters")
	}

	toDo.ActiveAt = payload.ActiveAt
	toDo.Title = payload.Title

	s.db.Store(id, toDo)

	return nil
}

func (s *Store) ChangeToDone(id string) error {
	task, ok := s.db.Load(id)
	if !ok {
		return errors.New("given task does not exist")
	}

	toDo, ok := task.(types.ToDo)
	if !ok {
		return errors.New("task is not of type ToDo")
	}

	toDo.ActiveAt = time.Now()

	s.db.Store(id, toDo)

	return nil
}
