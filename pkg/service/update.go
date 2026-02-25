package service

import (
	"errors"

	"github.com/masputrawae/go-todo-cli/pkg/model"
)

func (s *Service) Update(i *model.TodoUpdateType) error {
	if i == nil {
		return errors.New("parameter tidak boleh kosong")
	}
	if i.ID == nil {
		return errors.New("id tidak boleh kosong")
	}

	var todo model.TodoUpdateType
	todo.ID = i.ID

	if i.Task != nil {
		todo.Task = i.Task
	}
	if i.Status != nil {
		todo.Status = i.Status
	}
	if i.Priority != nil {
		todo.Priority = i.Priority
	}
	if i.Category != nil {
		todo.Category = i.Category
	}

	return s.Repo.Update(&todo)
}
