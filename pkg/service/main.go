package service

import (
	"errors"
	"strings"

	"github.com/masputrawae/go-todo-cli/pkg/model"
	"github.com/masputrawae/go-todo-cli/pkg/repo"
)

type Service struct {
	Repo repo.Manage
}

type Manage interface {
	Create(i *model.TodoCreateType) error
	Update(i *model.TodoUpdateType) error
	Delete(i *model.TodoDeleteType) error
	GetTodo(i *model.TodoFindType) ([]model.Todo, error)
}

func New(repo repo.Manage) Manage {
	return &Service{Repo: repo}
}

// utils
func getMetaID(m []model.Meta, id string) (string, error) {
	nID := func(s string) string {
		return strings.ReplaceAll(strings.ToLower(s), " ", "-")
	}

	for i := range m {
		if nID(m[i].ID) == nID(id) {
			return m[i].ID, nil
		}
	}

	return "", errors.New("meta tidak ditemukan")
}
