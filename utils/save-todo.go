package utils

import (
	"encoding/json"

	"github.com/masputrawae/todo-cli/model"
)

func SaveTodo(data []model.Todo, fp string) error {
	file, err := ResolveFile(fp)
	if err != nil {
		return err
	}
	if err := json.NewEncoder(file).Encode(data); err != nil {
		return err
	}
	return file.Close()
}
