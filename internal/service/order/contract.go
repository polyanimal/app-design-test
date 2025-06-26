package order

import (
	"context"
	"time"

	"applicationDesignTest/internal/models"
)

type OrderRepository interface {
	AddOrder(ctx context.Context, order models.Order) error
	GetAvailableByDateAndRoomID(ctx context.Context, daysToBook []time.Time, roomID string) ([]models.RoomAvailability, error)
	UpdateAvailability(ctx context.Context, availabilityToUpdate []models.RoomAvailability) error
}
