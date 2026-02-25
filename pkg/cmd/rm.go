package cmd

import "github.com/masputrawae/go-todo-cli/pkg/model"

type RmCmd struct {
	ID int `arg:"" required:"" help:"ID tugas"`
}

func (a *RmCmd) Run(app *App) error {
	return app.Srv.Delete(&model.TodoDeleteType{ID: &a.ID})
}
