package service

import (
	"errors"
	"strings"
	"time"

	"github.com/masputrawae/todo-cli/model"
	"github.com/masputrawae/todo-cli/repo/todo"
)

var (
	ErrStatusNotMatch   = errors.New("status tidak cocok")
	ErrPriorityNotMatch = errors.New("prioritas tidak cocok")
	ErrNotChanged       = errors.New("tidak ada yang di ubah")
)

type Service struct {
	Todo       todo.Manage
	Statuses   map[string]model.Status
	Priorities map[string]model.Priority
}

type Manage interface {
	Create(model.TodoAdd) error
	Update(model.TodoEdit) error
	Delete(int) error
}

func New(t todo.Manage, s map[string]model.Status, p map[string]model.Priority) Manage {
	return &Service{Todo: t, Statuses: s, Priorities: p}
}

func (s *Service) Create(a model.TodoAdd) error {
	var todo model.Todo
	todo.ID = s.Todo.GenID()
	todo.Task = a.Task

	if a.Project != nil {
		todo.Project = a.Project
	}

	if a.Context != nil {
		todo.Context = a.Context
	}

	if a.Priority != nil {
		pr := strings.ToUpper(*a.Priority)
		_, ok := s.Priorities[pr]
		if !ok {
			return ErrPriorityNotMatch
		}
		todo.Priority = &pr
	}

	if a.Status != nil {
		st := strings.ReplaceAll(strings.ToLower(*a.Priority), " ", "-")
		_, ok := s.Statuses[st]
		if !ok {
			return ErrStatusNotMatch
		}
		todo.Status = &st
	}

	return s.Todo.Add(todo)
}

func (s *Service) Update(u model.TodoEdit) error {
	var todo model.Todo
	updated := false

	i, err := s.Todo.FindIndexByID(u.ID)
	if err != nil {
		return err
	}

	o, err := s.Todo.FindByIndex(i)
	if err != nil {
		return err
	}

	if u.Task != nil && u.Task != &o.Task {
		todo.Task = *u.Task
		updated = true
	}

	if u.Project != nil && u.Project != o.Project {
		todo.Project = u.Project
		updated = true
	}

	if u.Priority != nil && u.Priority != o.Priority {
		pr := strings.ToUpper(*u.Priority)
		_, ok := s.Priorities[pr]
		if !ok {
			return ErrPriorityNotMatch
		}

		todo.Priority = &pr
		updated = true
	}

	if u.Status != nil && u.Status != o.Status {
		st := strings.ReplaceAll(strings.ToLower(*u.Priority), " ", "-")
		_, ok := s.Statuses[st]
		if !ok {
			return ErrStatusNotMatch
		}

		todo.Project = u.Project
		updated = true
	}

	if !updated {
		return ErrNotChanged
	}

	dt := time.Now()
	todo.Lastmod = &dt
	return s.Todo.Edit(i, todo)
}

func (s *Service) Delete(id int) error {
	return s.Todo.Delete(id)
}
