package utils

import (
	"encoding/json"
	"log"

	"github.com/masputrawae/todo-cli/model"
)

func LoadTodo(fp string) []model.Todo {
	var data []model.Todo
	file, err := ResolveFile(fp)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		if err.Error() == "EOF" {
			if err := json.NewEncoder(file).Encode(&data); err != nil {
				log.Fatal(err)
			}
			return data
		}
		log.Fatal(err)
	}

	return data
}
