package repo

import (
	"encoding/json"
	"os"
)

func (r *Repo) Load() error {
	file, err := os.ReadFile(r.TodoFile)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &r.Todos)
}
