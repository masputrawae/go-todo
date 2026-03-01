/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/masputrawae/todo-cli/model"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var todo model.TodoAdd
		todo.Task = args[0]

		if fProject != "" {
			todo.Project = &fProject
		}
		if fContext != "" {
			todo.Context = &fContext
		}
		if fPriority != "" {
			todo.Priority = &fPriority
		}
		if fStatus != "" {
			todo.Status = &fStatus
		}

		if err := Svc.Create(todo); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&fProject, "project", "P", "", "project")
	addCmd.Flags().StringVarP(&fPriority, "priority", "p", "", "priority")
	addCmd.Flags().StringVarP(&fStatus, "status", "s", "", "status")
	addCmd.Flags().StringVarP(&fContext, "context", "c", "", "context")
}
