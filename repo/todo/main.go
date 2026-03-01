package todo

import (
	"errors"
	"sort"

	"github.com/masputrawae/todo-cli/model"
	"github.com/masputrawae/todo-cli/utils"
)

var (
	ErrIDNotFound   = errors.New("id tidak ditemukan")
	ErrTaskNotFound = errors.New("tugas tidak ditemukan")
)

type Todo struct {
	File string
	Data []model.Todo
}

type Manage interface {
	GenID() int
	FindIndexByID(int) (int, error)

	Add(model.Todo) error
	Edit(int, model.Todo) error
	Delete(int) error

	FindByID(int) (model.Todo, error)
	FindByIndex(int) (model.Todo, error)
	FindByProject(string) ([]model.Todo, error)
	FindByPriority(string) ([]model.Todo, error)
	FindByStatus(string) ([]model.Todo, error)
	FindByContext(string) ([]model.Todo, error)
}

func New(t []model.Todo, f string) Manage {
	return &Todo{File: f, Data: t}
}

// add todo
func (t *Todo) Add(d model.Todo) error {
	t.Data = append(t.Data, d)
	return utils.SaveTodo(t.Data, t.File)
}

// edit todo
func (t *Todo) Edit(i int, d model.Todo) error {
	t.Data[i] = d
	return utils.SaveTodo(t.Data, t.File)
}

// delete todo
func (t *Todo) Delete(i int) error {
	t.Data = append(t.Data[:i], t.Data[i+1:]...)
	return utils.SaveTodo(t.Data, t.File)
}

// generate id
func (t *Todo) GenID() int {
	// inisialisasi
	id := 1

	// urutkan datanya dulu
	sort.Slice(t.Data, func(i, j int) bool {
		return t.Data[i].ID < t.Data[j].ID
	})

	// cari id, dapatkan id yang belum digunakan
	for _, v := range t.Data {
		if v.ID == id {
			id++
		}
	}

	return id
}

// cari index berdasarkan id
func (t *Todo) FindIndexByID(id int) (int, error) {
	for i, v := range t.Data {
		if v.ID == id {
			return i, nil
		}
	}
	return -1, ErrIDNotFound
}

func (t *Todo) FindByIndex(i int) (model.Todo, error) {
	return t.Data[i], nil
}

func (t *Todo) FindByID(id int) (model.Todo, error) {
	for i, v := range t.Data {
		if v.ID == id {
			return t.Data[i], nil
		}
	}
	return model.Todo{}, ErrIDNotFound
}

func (t *Todo) FindByProject(p string) ([]model.Todo, error) {
	var results []model.Todo
	for i, v := range t.Data {
		if v.Project == &p {
			results = append(results, t.Data[i])
		}
	}
	if len(results) == 0 {
		return nil, ErrTaskNotFound
	}
	return results, nil
}

func (t *Todo) FindByPriority(p string) ([]model.Todo, error) {
	var results []model.Todo
	for i, v := range t.Data {
		if v.Priority == &p {
			results = append(results, t.Data[i])
		}
	}
	if len(results) == 0 {
		return nil, ErrTaskNotFound
	}
	return results, nil
}

func (t *Todo) FindByStatus(s string) ([]model.Todo, error) {
	var results []model.Todo
	for i, v := range t.Data {
		if v.Status == &s {
			results = append(results, t.Data[i])
		}
	}
	if len(results) == 0 {
		return nil, ErrTaskNotFound
	}
	return results, nil
}

func (t *Todo) FindByContext(c string) ([]model.Todo, error) {
	var results []model.Todo
	for i, v := range t.Data {
		if v.Status == &c {
			results = append(results, t.Data[i])
		}
	}
	if len(results) == 0 {
		return nil, ErrTaskNotFound
	}
	return results, nil
}
