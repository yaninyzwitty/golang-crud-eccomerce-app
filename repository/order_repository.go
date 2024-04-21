package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"github.com/yaninyzwitty/crud-eccomerce-app/model"
)

type OrderRepository interface {
	GetOrders(ctx context.Context) ([]model.Order, error)
	GetOrder(ctx context.Context, id string) (model.Order, error)
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
	UpdateOrder(ctx context.Context, orderId string, order model.Order) (model.Order, error)
	DeleteOrder(ctx context.Context, id string) error
}

type orderRepository struct {
	ctx     context.Context
	session *gocql.Session
}

func NewOrderRepository(ctx context.Context, session *gocql.Session) OrderRepository {
	return &orderRepository{ctx, session}

}

func (r *orderRepository) GetOrders(ctx context.Context) ([]model.Order, error) {
	var orders []model.Order
	iter := r.session.Query("SELECT order_id, customer_id, order_date, total_amount FROM merchandise.orders").Iter()
	for {
		var order model.Order
		if !iter.Scan(&order.ID, &order.CustomerID, &order.OrderDate, &order.Amount) {
			break
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *orderRepository) GetOrder(ctx context.Context, id string) (model.Order, error) {
	var order model.Order
	if err := r.session.Query("SELECT order_id, customer_id, order_date, total_amount FROM merchandise.orders WHERE order_id = ?", id).Scan(&order.ID, &order.CustomerID, &order.OrderDate, &order.Amount); err != nil {
		return model.Order{}, fmt.Errorf("error getting order: %w", err)
	}
	return order, nil

}

func (r *orderRepository) CreateOrder(ctx context.Context, order model.Order) (model.Order, error) {
	order.ID = gocql.TimeUUID()
	order.OrderDate = time.Now()

	if err := r.session.Query("INSERT INTO merchandise.orders (order_id, customer_id, order_date, total_amount) VALUES (?, ?, ?, ?)", order.ID, order.CustomerID, order.OrderDate, order.Amount).Exec(); err != nil {
		return model.Order{}, fmt.Errorf("error creating order: %w", err)
	}

	return order, nil

}

func (r *orderRepository) UpdateOrder(ctx context.Context, orderId string, order model.Order) (model.Order, error) {
	if err := r.session.Query("UPDATE merchandise.orders SET customer_id = ? order_date = ?, total_amount = ? WHERE order_id = ?", order.CustomerID, order.OrderDate, order.Amount, orderId).Exec(); err != nil {
		return model.Order{}, fmt.Errorf("error updating order: %w", err)
	}
	return order, nil

}

func (r *orderRepository) DeleteOrder(ctx context.Context, orderId string) error {
	if err := r.session.Query("DELETE FROM merchandise.orders WHERE order_id = ?", orderId).Exec(); err != nil {
		return fmt.Errorf("error deleting order: %w", err)
	}
	return nil
}
