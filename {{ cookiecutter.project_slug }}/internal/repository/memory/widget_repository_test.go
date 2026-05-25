package memory

import (
	"context"
	"errors"
	"testing"

	"{{ cookiecutter.module_path }}/internal/domain"
)

func TestWidgetRepositorySaveAndGetWidget(t *testing.T) {
	t.Parallel()

	repo := NewWidgetRepository()
	widget := domain.Widget{ID: "widget-1", Name: "Platform"}

	if err := repo.SaveWidget(context.Background(), widget); err != nil {
		t.Fatalf("unexpected save error: %v", err)
	}

	got, err := repo.GetWidget(context.Background(), widget.ID)
	if err != nil {
		t.Fatalf("unexpected get error: %v", err)
	}

	if got != widget {
		t.Fatalf("expected %#v, got %#v", widget, got)
	}
}

func TestWidgetRepositoryGetMissingWidget(t *testing.T) {
	t.Parallel()

	repo := NewWidgetRepository()

	_, err := repo.GetWidget(context.Background(), "missing")
	if !errors.Is(err, ErrWidgetNotFound) {
		t.Fatalf("expected ErrWidgetNotFound, got %v", err)
	}
}
