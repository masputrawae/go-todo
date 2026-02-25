package cmd

import (
	"github.com/olekukonko/tablewriter"
	"os"

	"github.com/masputrawae/go-todo-cli/pkg/model"
)

type ListCmd struct {
	ID       *int    `short:"i" help:"Cari berdasarkan ID"`
	Status   *string `short:"s" help:"Cari Berdasarkan status"`
	Priority *string `short:"p" help:"Cari Berdasarkan Prioritas"`
	Category *string `short:"c" help:"Cari Berdasarkan Kategori"`
}

func (a *ListCmd) Run(app *App) error {
	var filters model.TodoFindType

	if a.ID != nil {
		filters.ID = a.ID
	}
	if a.Status != nil {
		filters.Status = a.Status
	}
	if a.Priority != nil {
		filters.Priority = a.Priority
	}
	if a.Category != nil {
		filters.Priority = a.Category
	}
	todos, err := app.Srv.GetTodo(&filters)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Task", "Status", "Priority", "Category", "Lastmod"})

	for _, v := range todos {
		table.Append(v)
	}

	table.Render()
	return nil
}
