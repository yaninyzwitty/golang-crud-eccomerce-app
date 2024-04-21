package service

import (
	"context"

	"github.com/yaninyzwitty/crud-eccomerce-app/model"
	"github.com/yaninyzwitty/crud-eccomerce-app/repository"
)

type OrderService interface {
	GetOrders(ctx context.Context) ([]model.Order, error)
	GetOrder(ctx context.Context, id string) (model.Order, error)
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
	UpdateOrder(ctx context.Context, id string, order model.Order) (model.Order, error)
	DeleteOrder(ctx context.Context, id string) error
}

type orderService struct {
	handler repository.OrderRepository
}

func NewOrderService(handler repository.OrderRepository) OrderService {
	return &orderService{
		handler: handler,
	}
}

func (s *orderService) GetOrders(ctx context.Context) ([]model.Order, error) {
	return s.handler.GetOrders(ctx)
}

func (s *orderService) GetOrder(ctx context.Context, id string) (model.Order, error) {
	return s.handler.GetOrder(ctx, id)
}
func (s *orderService) CreateOrder(ctx context.Context, order model.Order) (model.Order, error) {
	return s.handler.CreateOrder(ctx, order)
}

func (s *orderService) UpdateOrder(ctx context.Context, id string, order model.Order) (model.Order, error) {
	return s.handler.UpdateOrder(ctx, id, order)

}

func (s *orderService) DeleteOrder(ctx context.Context, id string) error {
	return s.handler.DeleteOrder(ctx, id)
}
