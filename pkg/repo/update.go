package repo

import (
	"errors"
	"time"

	"github.com/masputrawae/go-todo-cli/pkg/model"
)

func (r *Repo) Update(i *model.TodoUpdateType) error {
	dt := time.Now()
	updated := false
	var index int

	for idx := range r.Todos {
		if r.Todos[idx].ID == *i.ID {
			index = idx
		}
	}

	if i.Task != nil {
		r.Todos[index].Task = *i.Task
		updated = true
	}
	if i.Status != nil {
		r.Todos[index].Status = *i.Status
		updated = true
	}
	if i.Priority != nil {
		r.Todos[index].Priority = *i.Priority
		updated = true
	}
	if i.Category != nil {
		r.Todos[index].Category = i.Category
		updated = true
	}

	if updated {
		r.Todos[index].Lastmod = &dt
		return save(r.Todos, r.TodoFile)
	}

	return errors.New("tidak ada yang di ubah")
}
