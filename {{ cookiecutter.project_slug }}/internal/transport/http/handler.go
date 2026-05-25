package http

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"{{ cookiecutter.module_path }}/internal/domain"
)

type WidgetCreator interface {
	CreateWidget(ctx context.Context, id string, name string) (domain.Widget, error)
	GetWidget(ctx context.Context, id string) (domain.Widget, error)
}

type Handler struct {
	widgets WidgetCreator
	logger  *slog.Logger
}

func NewHandler(widgets WidgetCreator, logger *slog.Logger) *Handler {
	return &Handler{widgets: widgets, logger: logger}
}

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", h.health)
	mux.HandleFunc("POST /widgets", h.createWidget)
	mux.HandleFunc("GET /widgets/{id}", h.getWidget)
	return mux
}

func (h *Handler) health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

type createWidgetRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) createWidget(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var request createWidgetRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json"})
		return
	}

	widget, err := h.widgets.CreateWidget(ctx, request.ID, request.Name)
	if err != nil {
		h.logger.Warn("create widget failed", "error", err)
		writeJSON(w, statusFromError(err), map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusCreated, widget)
}

func (h *Handler) getWidget(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimSpace(r.PathValue("id"))
	if id == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "missing id"})
		return
	}

	widget, err := h.widgets.GetWidget(r.Context(), id)
	if err != nil {
		h.logger.Warn("get widget failed", "error", err)
		writeJSON(w, statusFromError(err), map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, widget)
}

func statusFromError(err error) int {
	if errors.Is(err, domain.ErrInvalidWidgetName) {
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(value); err != nil {
		panic(err)
	}
}
