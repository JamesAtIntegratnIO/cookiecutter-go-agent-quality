package memory

import (
	"context"
	"errors"
	"sync"

	"{{ cookiecutter.module_path }}/internal/domain"
)

var ErrWidgetNotFound = errors.New("widget not found")

type WidgetRepository struct {
	mu      sync.RWMutex
	widgets map[string]domain.Widget
}

func NewWidgetRepository() *WidgetRepository {
	return &WidgetRepository{widgets: make(map[string]domain.Widget)}
}

func (r *WidgetRepository) SaveWidget(_ context.Context, widget domain.Widget) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.widgets[widget.ID] = widget
	return nil
}

func (r *WidgetRepository) GetWidget(_ context.Context, id string) (domain.Widget, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	widget, ok := r.widgets[id]
	if !ok {
		return domain.Widget{}, ErrWidgetNotFound
	}

	return widget, nil
}
