package domain

import (
	"errors"
	"strings"
)

var ErrInvalidWidgetName = errors.New("invalid widget name")

type Widget struct {
	ID   string
	Name string
}

func NewWidget(id string, name string) (Widget, error) {
	trimmedID := strings.TrimSpace(id)
	trimmedName := strings.TrimSpace(name)

	if trimmedID == "" || trimmedName == "" {
		return Widget{}, ErrInvalidWidgetName
	}

	return Widget{ID: trimmedID, Name: trimmedName}, nil
}
