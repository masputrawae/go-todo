package model

import "time"

type Todo struct {
	ID       int        `json:"id"`
	Status   string     `json:"status"`
	Priority string     `json:"priority"`
	Task     string     `json:"task"`
	Category *string    `json:"category,omitempty"`
	Lastmod  *time.Time `json:"lastmod,omitempty"`
}

// type for create new todo
type TodoCreateType struct {
	Status   *string
	Priority *string
	Task     *string
	Category *string
}

// type for update todo
type TodoUpdateType struct {
	ID       *int
	Status   *string
	Priority *string
	Task     *string
	Category *string
}

// type for delete todo
type TodoDeleteType struct {
	ID *int
}

// type for find todo
type TodoFindType struct {
	ID       *int
	Status   *string
	Priority *string
	Category *string
}
