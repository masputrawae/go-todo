package store

import "github.com/masputrawae/go-todo-cli/pkg/model"

var Priorities = []model.Meta{
	{ID: "highest", Name: "Highest"},
	{ID: "high", Name: "High"},
	{ID: "medium", Name: "Medium"},
	{ID: "low", Name: "Low"},
	{ID: "lowest", Name: "Lowest"},
}
