package repo

import (
	"errors"

	"github.com/masputrawae/go-todo-cli/pkg/model"
)

func (r *Repo) Delete(i *model.TodoDeleteType) error {
	for index := range r.Todos {
		if r.Todos[index].ID == *i.ID {
			r.Todos = append(r.Todos[:index], r.Todos[index+1:]...)
			return save(r.Todos, r.TodoFile)
		}
	}

	return errors.New("id tidak ditemukan")
}
