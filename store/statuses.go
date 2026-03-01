package store

import "github.com/masputrawae/todo-cli/model"

var Statuses = map[string]model.Status{
	"planning":    {Name: "Planning", Emoji: "📋️"},
	"active":      {Name: "Active", Emoji: "🟢"},
	"in-progress": {Name: "In Progress", Emoji: "📊"},
	"done":        {Name: "Done", Emoji: "✅"},
	"cancelled":   {Name: "Cancelled", Emoji: "❌"},
	"archive":     {Name: "Archive", Emoji: "📦️"},
}
