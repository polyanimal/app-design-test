package order

import (
	"applicationDesignTest/internal/models"
	"context"
)

type OrderService interface {
	CreateOrder(ctx context.Context, newOrder models.Order) error
}
