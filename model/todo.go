package model

import "time"

type Todo struct {
	ID       int        `json:"id"`
	Task     string     `json:"task"`
	Project  *string    `json:"project,omitempty"`
	Priority *string    `json:"priority,omitempty"`
	Context  *string    `json:"context,omitempty"`
	Status   *string    `json:"status,omitempty"`
	Lastmod  *time.Time `json:"lastmod,omitempty"`
}

type TodoAdd struct {
	ID       int
	Task     string
	Project  *string
	Priority *string
	Context  *string
	Status   *string
	Lastmod  *time.Time
}

type TodoEdit struct {
	ID       int
	Task     *string
	Project  *string
	Priority *string
	Context  *string
	Status   *string
	Lastmod  *time.Time
}

type TodoFind struct {
	ID       *int
	Project  *string
	Priority *string
	Context  *string
	Status   *string
}
