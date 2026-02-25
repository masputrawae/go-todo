package cmd

import "github.com/masputrawae/go-todo-cli/pkg/model"

type AddCmd struct {
	Task     string  `arg:"" required:"" help:"Deskripsi tugas"`
	Status   *string `short:"s" help:"Status tugas"`
	Priority *string `short:"p" help:"Prioritas tugas"`
	Category *string `short:"c" help:"Kategori tugas"`
}

func (a *AddCmd) Run(app *App) error {
	var todo model.TodoCreateType
	todo.Task = &a.Task

	if a.Status != nil {
		todo.Status = a.Status
	}
	if a.Priority != nil {
		todo.Priority = a.Priority
	}
	if a.Category != nil {
		todo.Category = a.Category
	}

	return app.Srv.Create(&todo)
}
