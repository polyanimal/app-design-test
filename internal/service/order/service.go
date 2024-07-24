package order

import (
	"context"
	"fmt"
	"time"

	"applicationDesignTest/internal/models"
	"applicationDesignTest/pkg/date"
)

type Service struct {
	repo OrderRepository
}

func NewService(repo OrderRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateOrder(ctx context.Context, newOrder models.Order) error {
	daysToBook := date.DaysBetween(newOrder.From, newOrder.To)

	unavailableDays := make(map[time.Time]struct{})
	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}

	availability, err := s.repo.GetAvailableByDateAndRoomID(ctx, daysToBook, newOrder.RoomID)
	if err != nil {
		return err
	}

	availabilityToUpdate := make([]models.RoomAvailability, 0)
	for _, dayToBook := range daysToBook {
		for _, av := range availability {
			if !av.Date.Equal(dayToBook) || av.Quota < 1 {
				continue
			}
			av.Quota -= 1
			availabilityToUpdate = append(availabilityToUpdate, av)
			delete(unavailableDays, dayToBook)
		}
	}

	if len(unavailableDays) != 0 {
		return fmt.Errorf("hotel room is not available for selected dates:\n%v\n%v", newOrder, unavailableDays)
	}

	if err := s.repo.UpdateAvailability(ctx, availabilityToUpdate); err != nil {
		return err
	}

	return s.repo.AddOrder(ctx, newOrder)
}
