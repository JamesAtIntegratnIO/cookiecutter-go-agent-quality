package domain

import (
	"errors"
	"testing"
)

func TestNewWidget(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		id      string
		name    string
		wantErr error
	}{
		"valid":      {id: "widget-1", name: "Platform", wantErr: nil},
		"empty id":   {id: "", name: "Platform", wantErr: ErrInvalidWidgetName},
		"empty name": {id: "widget-1", name: "", wantErr: ErrInvalidWidgetName},
		"spaces":     {id: "   ", name: "   ", wantErr: ErrInvalidWidgetName},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewWidget(tc.id, tc.name)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("expected error %v, got %v", tc.wantErr, err)
			}

			if tc.wantErr == nil && got.ID == "" {
				t.Fatal("expected widget id to be set")
			}
		})
	}
}
