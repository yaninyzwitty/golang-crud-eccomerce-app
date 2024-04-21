package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gocql/gocql"
	"github.com/yaninyzwitty/crud-eccomerce-app/model"
	"github.com/yaninyzwitty/crud-eccomerce-app/service"
)

type OrderHandler struct {
	service service.OrderService
}

func NewOrderHandler(service service.OrderService) *OrderHandler {
	return &OrderHandler{
		service,
	}
}

func (o *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	orders, err := o.service.GetOrders(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// here you can have a functions that return a better json res
	respondCategoryWithJson(w, http.StatusOK, orders)

}
func (o *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	order, err := o.service.GetOrder(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	respondWithJSON(w, http.StatusOK, order)

}

func (o *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var order model.Order

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	// Validate order data
	if (order.CustomerID == gocql.UUID{}) || order.Amount <= 0 {
		http.Error(w, "Invalid or incomplete order data: customer ID and amount are required fields", http.StatusBadRequest)
		return
	}

	createdOrder, err := o.service.CreateOrder(ctx, order)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	respondWithJSON(w, http.StatusCreated, createdOrder)

}
func (o *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	var order model.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	updatedOrder, err := o.service.UpdateOrder(ctx, id, order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	respondWithJSON(w, http.StatusOK, updatedOrder)

}
func (o *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	if err := o.service.DeleteOrder(ctx, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")

}
