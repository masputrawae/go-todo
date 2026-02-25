package store

import "github.com/masputrawae/go-todo-cli/pkg/model"

var Statuses = []model.Meta{
	{ID: "planning", Name: "Planning", Emoji: "📝"},
	{ID: "active", Name: "Active", Emoji: "📝"},
	{ID: "in-progress", Name: "In Progress"},
	{ID: "done", Name: "done", Emoji: "📝"},
	{ID: "cancelled", Name: "Cancelled", Emoji: "📝"},
	{ID: "archive", Name: "Archive", Emoji: "📝"},
}
