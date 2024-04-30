package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yaninyzwitty/crud-eccomerce-app/model"
	"github.com/yaninyzwitty/crud-eccomerce-app/service"
)

type CustomerHandler struct {
	service service.CustomerService
}

func NewCustomerHandler(service service.CustomerService) *CustomerHandler {
	return &CustomerHandler{service}
}

func (h *CustomerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	customers, err := h.service.GetCustomers(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, customers)

}
func (h *CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := r.Context()

	customer, err := h.service.GetCustomer(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, customer)

}
func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var customer model.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	createdCustomer, err := h.service.CreateCustomer(ctx, customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, createdCustomer)

}
func (h *CustomerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := r.Context()
	var customer model.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	updatedCustomer, err := h.service.UpdateCustomer(ctx, id, customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedCustomer)

}
func (h *CustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := r.Context()

	if err := h.service.DeleteCustomer(ctx, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")

}
