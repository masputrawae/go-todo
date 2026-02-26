package repo

import "github.com/masputrawae/go-todo-cli/pkg/model"

func (r *Repo) Find(i *model.TodoFindType) []model.Todo {
	var (
		results []model.Todo
		todo    model.Todo
	)

	isFilter := false
	for index := range r.Todos {
		if i.ID != nil {
			if r.Todos[index].ID == *i.ID {
				todo = r.Todos[index]
				isFilter = true
				results = append(results, todo)
				break
			}
		}

		if i.Status != nil {
			if r.Todos[index].Status == *i.Status {
				todo = r.Todos[index]
				results = append(results, todo)
				isFilter = true
			}
		}

		if i.Priority != nil {
			if r.Todos[index].Priority == *i.Priority {
				todo = r.Todos[index]
				results = append(results, todo)
				isFilter = true
			}
		}

		if i.Category != nil {
			if r.Todos[index].Category == i.Category {
				todo = r.Todos[index]
				results = append(results, todo)
				isFilter = true
			}
		}
	}

	if isFilter {
		return results
	}

	results = r.Todos
	return results
}
