package repository

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"todo/internal/model"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *model.Todo) error
	List(ctx context.Context) ([]*model.Todo, error)
	GetByID(ctx context.Context, id int) (*model.Todo, error)
	Update(ctx context.Context, todo *model.Todo) error
	Delete(ctx context.Context, id int) error
}

type todoRepository struct{ db *gorm.DB }

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) Create(ctx context.Context, todo *model.Todo) error {
	if err := r.db.WithContext(ctx).Create(todo).Error; err != nil {
		return fmt.Errorf("create todo: %w", err)
	}
	return nil
}

func (r *todoRepository) List(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo
	if err := r.db.WithContext(ctx).Order("id").Find(&todos).Error; err != nil {
		return nil, fmt.Errorf("list todos: %w", err)
	}
	return todos, nil
}

func (r *todoRepository) GetByID(ctx context.Context, id int) (*model.Todo, error) {
	var todo model.Todo
	if err := r.db.WithContext(ctx).First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.ErrNotFound
		}
		return nil, fmt.Errorf("get todo %d: %w", id, err)
	}
	return &todo, nil
}

func (r *todoRepository) Update(ctx context.Context, todo *model.Todo) error {
	if err := r.db.WithContext(ctx).Save(todo).Error; err != nil {
		return fmt.Errorf("update todo %d: %w", todo.ID, err)
	}
	return nil
}

func (r *todoRepository) Delete(ctx context.Context, id int) error {
	result := r.db.WithContext(ctx).Delete(&model.Todo{}, id)
	if result.Error != nil {
		return fmt.Errorf("delete todo %d: %w", id, result.Error)
	}
	if result.RowsAffected == 0 {
		return model.ErrNotFound
	}
	return nil
}
