/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/masputrawae/todo-cli/model"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var todo model.TodoEdit
		todo.ID = id

		if fTask != "" {
			todo.Task = &fTask
		}

		if fProject != "" {
			todo.Priority = &fProject
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

		if err := Svc.Update(todo); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().StringVarP(&fTask, "task", "t", "", "task")
	editCmd.Flags().StringVarP(&fProject, "project", "P", "", "project")
	editCmd.Flags().StringVarP(&fContext, "context", "c", "", "context")
	editCmd.Flags().StringVarP(&fPriority, "priority", "p", "", "priority")
	editCmd.Flags().StringVarP(&fStatus, "status", "s", "", "status")
}
