package cmd

import (
	"github.com/masputrawae/go-todo-cli/pkg/service"
)

type App struct {
	Srv service.Manage
}

var CLI struct {
	Add  AddCmd  `cmd:"" help:"Menambah tugas baru"`
	Rm   RmCmd   `cmd:"" help:"Hapus tugas"`
	Edit EditCmd `cmd:"" help:"Edit tugas"`
	List ListCmd `cmd:"" help:"Lihat tugas"`
}
