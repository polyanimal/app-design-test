package order

import (
	"context"
	"time"

	"applicationDesignTest/internal/models"
	"applicationDesignTest/pkg/date"
)

type OrderRepository struct {
	storage *OrderStorage
}

func NewRepo(storage *OrderStorage) *OrderRepository {
	return &OrderRepository{
		storage: storage,
	}
}

func (r *OrderRepository) AddOrder(ctx context.Context, order models.Order) error {
	r.storage.Orders = append(r.storage.Orders, order)

	return nil
}

func (r *OrderRepository) GetAvailableByDateAndRoomID(ctx context.Context, daysToBook []time.Time, roomID string) ([]models.RoomAvailability, error) {
	availability := make([]models.RoomAvailability, 0)

	for _, dayToBook := range daysToBook {
		for _, a := range r.storage.Availability {
			if a.RoomID == roomID && a.Date.Equal(dayToBook) && a.Quota > 0 {
				availability = append(availability, a)
			}
		}
	}

	return availability, nil
}

func (r *OrderRepository) UpdateAvailability(ctx context.Context, availabilityToUpdate []models.RoomAvailability) error {
	for _, atu := range availabilityToUpdate {
		for i, a := range r.storage.Availability {
			if a.RoomID == atu.RoomID && a.Date.Equal(atu.Date) {
				r.storage.Availability[i].Quota = atu.Quota
			}
		}
	}

	return nil
}

type OrderStorage struct {
	Orders       []models.Order
	Availability []models.RoomAvailability
}

func NewStorage() *OrderStorage {
	Orders := []models.Order{}

	Availability := []models.RoomAvailability{
		{HotelID: "reddison", RoomID: "lux", Date: date.Date(2024, 1, 1), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: date.Date(2024, 1, 2), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: date.Date(2024, 1, 3), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: date.Date(2024, 1, 4), Quota: 1},
		{HotelID: "reddison", RoomID: "lux", Date: date.Date(2024, 1, 5), Quota: 0},
	}

	return &OrderStorage{
		Orders:       Orders,
		Availability: Availability,
	}
}
