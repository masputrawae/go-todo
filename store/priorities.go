package store

import "github.com/masputrawae/todo-cli/model"

var Priorities = map[string]model.Priority{
	"A": {Name: "Highest", Color: "\033[31;1m"},
	"B": {Name: "High", Color: "\033[33;1m"},
	"C": {Name: "Medium", Color: "\033[33;1m"},
	"D": {Name: "Low", Color: "\033[34;1m"},
	"E": {Name: "Lowest", Color: "\033[35;1m"},
}
