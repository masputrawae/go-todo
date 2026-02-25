package service

import (
	"errors"
	"github.com/masputrawae/go-todo-cli/pkg/model"
	"github.com/masputrawae/go-todo-cli/pkg/store"
)

func (s *Service) Create(i *model.TodoCreateType) error {
	if i == nil {
		return errors.New("parameter tidak boleh kosong")
	}
	if i.Task == nil {
		return errors.New("tugas tidak boleh kosong")
	}

	var todo model.TodoCreateType
	todo.Task = i.Task
	todo.Status = &store.Default.Status
	todo.Priority = &store.Default.Priority

	if i.Status != nil {
		id, err := getMetaID(store.Statuses, *i.Status)
		if err != nil {
			return err
		}
		todo.Status = &id
	}

	if i.Priority != nil {
		id, err := getMetaID(store.Priorities, *i.Priority)
		if err != nil {
			return err
		}
		todo.Priority = &id
	}

	if i.Category != nil {
		todo.Category = i.Category
	}

	return s.Repo.Create(&todo)
}
