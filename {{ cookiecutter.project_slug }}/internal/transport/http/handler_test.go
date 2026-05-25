package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"{{ cookiecutter.module_path }}/internal/domain"
)

type fakeWidgetService struct {
	widget domain.Widget
	err    error
}

func (f *fakeWidgetService) CreateWidget(_ context.Context, _ string, _ string) (domain.Widget, error) {
	if f.err != nil {
		return domain.Widget{}, f.err
	}

	return f.widget, nil
}

func (f *fakeWidgetService) GetWidget(_ context.Context, _ string) (domain.Widget, error) {
	if f.err != nil {
		return domain.Widget{}, f.err
	}

	return f.widget, nil
}

func TestHealth(t *testing.T) {
	t.Parallel()

	handler := NewHandler(&fakeWidgetService{}, slog.New(slog.NewTextHandler(io.Discard, nil)))

	request := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/healthz", nil)
	response := httptest.NewRecorder()

	handler.Routes().ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}
}

func TestCreateWidget(t *testing.T) {
	t.Parallel()

	widget := domain.Widget{ID: "widget-1", Name: "Platform"}
	handler := NewHandler(&fakeWidgetService{widget: widget}, slog.New(slog.NewTextHandler(io.Discard, nil)))
	body := bytes.NewBufferString(`{"id":"widget-1","name":"Platform"}`)

	request := httptest.NewRequestWithContext(context.Background(), http.MethodPost, "/widgets", body)
	response := httptest.NewRecorder()

	handler.Routes().ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, response.Code)
	}

	var got domain.Widget
	if err := json.NewDecoder(response.Body).Decode(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if got != widget {
		t.Fatalf("expected %#v, got %#v", widget, got)
	}
}

func TestCreateWidgetInvalidJSON(t *testing.T) {
	t.Parallel()

	handler := NewHandler(&fakeWidgetService{}, slog.New(slog.NewTextHandler(io.Discard, nil)))
	request := httptest.NewRequestWithContext(context.Background(), http.MethodPost, "/widgets", bytes.NewBufferString(`{`))
	response := httptest.NewRecorder()

	handler.Routes().ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, response.Code)
	}
}

func TestCreateWidgetDomainError(t *testing.T) {
	t.Parallel()

	handler := NewHandler(&fakeWidgetService{err: domain.ErrInvalidWidgetName}, slog.New(slog.NewTextHandler(io.Discard, nil)))
	body := bytes.NewBufferString(`{"id":"","name":""}`)
	request := httptest.NewRequestWithContext(context.Background(), http.MethodPost, "/widgets", body)
	response := httptest.NewRecorder()

	handler.Routes().ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, response.Code)
	}
}

func TestGetWidget(t *testing.T) {
	t.Parallel()

	widget := domain.Widget{ID: "widget-1", Name: "Platform"}
	handler := NewHandler(&fakeWidgetService{widget: widget}, slog.New(slog.NewTextHandler(io.Discard, nil)))
	request := httptest.NewRequestWithContext(context.Background(), http.MethodGet, "/widgets/widget-1", nil)
	response := httptest.NewRecorder()

	handler.Routes().ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}
}

func TestStatusFromError(t *testing.T) {
	t.Parallel()

	if got := statusFromError(domain.ErrInvalidWidgetName); got != http.StatusBadRequest {
		t.Fatalf("expected bad request, got %d", got)
	}

	if got := statusFromError(errors.New("boom")); got != http.StatusInternalServerError {
		t.Fatalf("expected internal server error, got %d", got)
	}
}
