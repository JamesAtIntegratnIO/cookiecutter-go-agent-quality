package service

import (
	"context"
	"fmt"

	"{{ cookiecutter.module_path }}/internal/domain"
)

type WidgetRepository interface {
	SaveWidget(ctx context.Context, widget domain.Widget) error
	GetWidget(ctx context.Context, id string) (domain.Widget, error)
}

type WidgetService struct {
	repository WidgetRepository
}

func NewWidgetService(repository WidgetRepository) *WidgetService {
	return &WidgetService{repository: repository}
}

func (s *WidgetService) CreateWidget(ctx context.Context, id string, name string) (domain.Widget, error) {
	widget, err := domain.NewWidget(id, name)
	if err != nil {
		return domain.Widget{}, fmt.Errorf("create widget: %w", err)
	}

	if err := s.repository.SaveWidget(ctx, widget); err != nil {
		return domain.Widget{}, fmt.Errorf("save widget: %w", err)
	}

	return widget, nil
}

func (s *WidgetService) GetWidget(ctx context.Context, id string) (domain.Widget, error) {
	widget, err := s.repository.GetWidget(ctx, id)
	if err != nil {
		return domain.Widget{}, fmt.Errorf("get widget: %w", err)
	}

	return widget, nil
}
