package main

import (
	"log"

	"github.com/alecthomas/kong"
	"github.com/masputrawae/go-todo-cli/pkg/cmd"
	"github.com/masputrawae/go-todo-cli/pkg/model"
	"github.com/masputrawae/go-todo-cli/pkg/repo"
	"github.com/masputrawae/go-todo-cli/pkg/service"
)

func main() {
	todos := repo.New([]model.Todo{}, "data/todos.json")
	todos.Load()
	srv := service.New(todos)

	kctx := kong.Parse(&cmd.CLI)
	if err := kctx.Run(&cmd.App{Srv: srv}); err != nil {
		log.Fatal(err)
	}
}
