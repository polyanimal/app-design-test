package order

import (
	"context"
	"encoding/json"
	"net/http"

	"applicationDesignTest/internal/models"
	"applicationDesignTest/pkg/logger"
)

type Handler struct {
	service OrderService
	logger  logger.Logger
}

func NewHandler(service OrderService, l logger.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  l,
	}
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder models.Order
	json.NewDecoder(r.Body).Decode(&newOrder)

	ctx := context.Background()

	if err := h.service.CreateOrder(ctx, newOrder); err != nil {
		http.Error(w, "Hotel room is not available for selected dates", http.StatusInternalServerError)
		h.logger.LogErrorf(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)

	h.logger.LogInfo("Order successfully created: %v", newOrder)
}
