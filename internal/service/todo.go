package service

import (
	"context"
	"strings"
	"time"

	"todo/internal/model"
	"todo/internal/repository"
)

type TodoService interface {
	Add(ctx context.Context, title string) (*model.Todo, error)
	List(ctx context.Context) ([]*model.Todo, error)
	Toggle(ctx context.Context, id int) (*model.Todo, error)
	UpdateTitle(ctx context.Context, id int, title string) (*model.Todo, error)
	Delete(ctx context.Context, id int) error
}

type todoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) Add(ctx context.Context, title string) (*model.Todo, error) {
	title, err := cleanTitle(title)
	if err != nil {
		return nil, err
	}

	todo := &model.Todo{Title: title}
	if err := s.repo.Create(ctx, todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *todoService) List(ctx context.Context) ([]*model.Todo, error) {
	return s.repo.List(ctx)
}

func (s *todoService) Toggle(ctx context.Context, id int) (*model.Todo, error) {
	todo, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	todo.Done = !todo.Done
	if todo.Done {
		now := time.Now()
		todo.CompletedAt = &now
	} else {
		todo.CompletedAt = nil
	}

	if err := s.repo.Update(ctx, todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *todoService) UpdateTitle(ctx context.Context, id int, title string) (*model.Todo, error) {
	title, err := cleanTitle(title)
	if err != nil {
		return nil, err
	}

	todo, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	todo.Title = title
	if err := s.repo.Update(ctx, todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *todoService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

// cleanTitle trims spaces and rejects empty titles.
func cleanTitle(title string) (string, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return "", model.ErrEmptyTitle
	}
	return title, nil
}
