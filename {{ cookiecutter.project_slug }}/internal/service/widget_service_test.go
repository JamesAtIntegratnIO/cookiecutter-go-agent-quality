package service

import (
	"context"
	"errors"
	"testing"

	"{{ cookiecutter.module_path }}/internal/domain"
)

type fakeWidgetRepository struct {
	saved domain.Widget
	err   error
}

func (f *fakeWidgetRepository) SaveWidget(_ context.Context, widget domain.Widget) error {
	if f.err != nil {
		return f.err
	}

	f.saved = widget
	return nil
}

func (f *fakeWidgetRepository) GetWidget(_ context.Context, _ string) (domain.Widget, error) {
	if f.err != nil {
		return domain.Widget{}, f.err
	}

	return f.saved, nil
}

func TestWidgetServiceCreateWidget(t *testing.T) {
	t.Parallel()

	repo := &fakeWidgetRepository{}
	svc := NewWidgetService(repo)

	widget, err := svc.CreateWidget(context.Background(), "widget-1", "Platform")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if widget.ID != "widget-1" || repo.saved.ID != "widget-1" {
		t.Fatalf("expected widget to be saved, got %#v", repo.saved)
	}
}

func TestWidgetServiceCreateWidgetWrapsRepositoryError(t *testing.T) {
	t.Parallel()

	repositoryErr := errors.New("repository unavailable")
	svc := NewWidgetService(&fakeWidgetRepository{err: repositoryErr})

	_, err := svc.CreateWidget(context.Background(), "widget-1", "Platform")
	if !errors.Is(err, repositoryErr) {
		t.Fatalf("expected wrapped repository error, got %v", err)
	}
}
