package repo

import (
	"github.com/masputrawae/go-todo-cli/pkg/model"
	"sort"
)

func (r *Repo) Create(i *model.TodoCreateType) error {
	sort.Slice(r.Todos, func(i, j int) bool { return r.Todos[i].ID < r.Todos[j].ID })
	var todo model.Todo

	todo.ID = genID(r.Todos)
	todo.Task = *i.Task
	todo.Status = *i.Status
	todo.Priority = *i.Priority

	if i.Category != nil {
		todo.Category = i.Category
	}

	r.Todos = append(r.Todos, todo)
	return save(r.Todos, r.TodoFile)
}
