/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/masputrawae/todo-cli/repo/todo"
	"github.com/masputrawae/todo-cli/service"
	"github.com/masputrawae/todo-cli/store"
	"github.com/masputrawae/todo-cli/utils"
	"github.com/spf13/cobra"
)

var (
	File = "todos.json"
	Data = utils.LoadTodo(File)
	Todo = todo.New(Data, File)
	Svc  = service.New(Todo, store.Statuses, store.Priorities)
)

// flags variables
var (
	fTask     string
	fProject  string
	fContext  string
	fPriority string
	fStatus   string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "A brief description of your application",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
