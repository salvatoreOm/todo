package service

import (
	"context"
	"errors"
	"testing"

	"todo/internal/model"
)

// fakeRepo is an in-memory repository used instead of a real database.
type fakeRepo struct {
	todos  map[int]*model.Todo
	nextID int
}

func newFakeRepo() *fakeRepo {
	return &fakeRepo{todos: map[int]*model.Todo{}, nextID: 1}
}

func (f *fakeRepo) Create(_ context.Context, t *model.Todo) error {
	t.ID = f.nextID
	f.nextID++
	f.todos[t.ID] = t
	return nil
}

func (f *fakeRepo) List(_ context.Context) ([]*model.Todo, error) {
	out := make([]*model.Todo, 0, len(f.todos))
	for id := 1; id < f.nextID; id++ {
		if t, ok := f.todos[id]; ok {
			out = append(out, t)
		}
	}
	return out, nil
}

func (f *fakeRepo) GetByID(_ context.Context, id int) (*model.Todo, error) {
	t, ok := f.todos[id]
	if !ok {
		return nil, model.ErrNotFound
	}
	return t, nil
}

func (f *fakeRepo) Update(_ context.Context, t *model.Todo) error {
	f.todos[t.ID] = t
	return nil
}

func (f *fakeRepo) Delete(_ context.Context, id int) error {
	if _, ok := f.todos[id]; !ok {
		return model.ErrNotFound
	}
	delete(f.todos, id)
	return nil
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name      string
		title     string
		wantTitle string
		wantErr   error
	}{
		{"valid", "buy milk", "buy milk", nil},
		{"trims spaces", "  buy milk  ", "buy milk", nil},
		{"empty", "", "", model.ErrEmptyTitle},
		{"only spaces", "   ", "", model.ErrEmptyTitle},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc := NewTodoService(newFakeRepo())

			got, err := svc.Add(context.Background(), tc.title)

			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("error = %v, want %v", err, tc.wantErr)
			}
			if err == nil && got.Title != tc.wantTitle {
				t.Errorf("title = %q, want %q", got.Title, tc.wantTitle)
			}
		})
	}
}

func TestToggle(t *testing.T) {
	ctx := context.Background()
	svc := NewTodoService(newFakeRepo())
	added, _ := svc.Add(ctx, "buy milk")

	done, err := svc.Toggle(ctx, added.ID)
	if err != nil {
		t.Fatal(err)
	}
	if !done.Done || done.CompletedAt == nil {
		t.Errorf("after first toggle: done=%v completedAt=%v, want done with a time", done.Done, done.CompletedAt)
	}

	undone, err := svc.Toggle(ctx, added.ID)
	if err != nil {
		t.Fatal(err)
	}
	if undone.Done || undone.CompletedAt != nil {
		t.Errorf("after second toggle: done=%v completedAt=%v, want not done and nil", undone.Done, undone.CompletedAt)
	}
}

func TestToggleNotFound(t *testing.T) {
	svc := NewTodoService(newFakeRepo())

	_, err := svc.Toggle(context.Background(), 99)

	if !errors.Is(err, model.ErrNotFound) {
		t.Errorf("error = %v, want ErrNotFound", err)
	}
}

func TestUpdateTitle(t *testing.T) {
	ctx := context.Background()
	svc := NewTodoService(newFakeRepo())
	added, _ := svc.Add(ctx, "buy milk")

	got, err := svc.UpdateTitle(ctx, added.ID, " buy oat milk ")
	if err != nil {
		t.Fatal(err)
	}
	if got.Title != "buy oat milk" {
		t.Errorf("title = %q, want %q", got.Title, "buy oat milk")
	}

	if _, err := svc.UpdateTitle(ctx, added.ID, "  "); !errors.Is(err, model.ErrEmptyTitle) {
		t.Errorf("empty title error = %v, want ErrEmptyTitle", err)
	}
	if _, err := svc.UpdateTitle(ctx, 99, "x"); !errors.Is(err, model.ErrNotFound) {
		t.Errorf("missing id error = %v, want ErrNotFound", err)
	}
}

func TestDelete(t *testing.T) {
	ctx := context.Background()
	svc := NewTodoService(newFakeRepo())
	added, _ := svc.Add(ctx, "buy milk")

	if err := svc.Delete(ctx, added.ID); err != nil {
		t.Fatal(err)
	}
	if err := svc.Delete(ctx, added.ID); !errors.Is(err, model.ErrNotFound) {
		t.Errorf("second delete error = %v, want ErrNotFound", err)
	}
}
