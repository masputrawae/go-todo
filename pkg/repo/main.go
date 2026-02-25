package repo

import (
	"encoding/json"
	"os"

	"github.com/masputrawae/go-todo-cli/pkg/model"
)

type Repo struct {
	TodoFile string
	Todos    []model.Todo
}

type Manage interface {
	Load() error
	Create(i *model.TodoCreateType) error
	Update(i *model.TodoUpdateType) error
	Delete(i *model.TodoDeleteType) error
	Find(i *model.TodoFindType) []model.Todo
}

func New(todos []model.Todo, todoFile string) Manage {
	return &Repo{Todos: todos, TodoFile: todoFile}
}

// utils
func save(todos []model.Todo, todoFile string) error {
	jData, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(todoFile, jData, 0644)
}

func genID(t []model.Todo) int {
	initID := 1
	for i := range t {
		if t[i].ID == initID {
			initID++
		}
	}
	return initID
}
