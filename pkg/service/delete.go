package service

import (
	"errors"

	"github.com/masputrawae/go-todo-cli/pkg/model"
)

func (s *Service) Delete(i *model.TodoDeleteType) error {
	if i == nil {
		return errors.New("parameter tidak boleh kosong")
	}
	if i.ID == nil {
		return errors.New("id tidak boleh kosong")
	}
	return s.Repo.Delete(i)
}
