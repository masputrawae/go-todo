package repo

import "github.com/masputrawae/go-todo-cli/pkg/model"

func (r *Repo) Find(i *model.TodoFindType) []model.Todo {
	var (
		results = make([]model.Todo, len(r.Todos))
		todo    model.Todo
	)

	isFilter := false
	for index := range r.Todos {
		if i.ID != nil {
			todo = r.Todos[index]
			isFilter = true
		}

		if i.Status != nil {
			todo = r.Todos[index]
			isFilter = true
		}

		if i.Priority != nil {
			todo = r.Todos[index]
			isFilter = true
		}

		if i.Category != nil {
			todo = r.Todos[index]
			isFilter = true
		}
	}

	if isFilter {
		results = append(results, todo)
		return results
	}

	results = r.Todos
	return results
}
