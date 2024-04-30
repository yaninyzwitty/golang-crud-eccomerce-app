package repository

import (
	"context"
	"fmt"

	"github.com/gocql/gocql"
	"github.com/yaninyzwitty/crud-eccomerce-app/model"
)

type CustomerRepository interface {
	GetCustomers(ctx context.Context) ([]model.Customer, error)
	GetCustomer(ctx context.Context, id string) (model.Customer, error)
	CreateCustomer(ctx context.Context, customer model.Customer) (model.Customer, error)
	UpdateCustomer(ctx context.Context, id string, customer model.Customer) (model.Customer, error)
	DeleteCustomer(ctx context.Context, id string) error
}

type customerRepository struct {
	ctx     context.Context
	session *gocql.Session
}

func NewCustomerRepository(ctx context.Context, session *gocql.Session) CustomerRepository {
	return &customerRepository{ctx, session}
}

func (r *customerRepository) GetCustomers(ctx context.Context) ([]model.Customer, error) {
	var customers []model.Customer
	iter := r.session.Query("SELECT customer_id, address, email, name FROM merchandise.customers").Iter()
	for {
		var customer model.Customer
		if !iter.Scan(&customer.ID, &customer.Address, &customer.Email, &customer.Name) {
			break
		}

		customers = append(customers, customer)
	}

	return customers, nil

}

func (r *customerRepository) GetCustomer(ctx context.Context, id string) (model.Customer, error) {
	var customer model.Customer
	if err := r.session.Query("SELECT customer_id, address, email, name FROM merchandise.customers WHERE customer_id = ?", id).Scan(&customer.ID, &customer.Address, &customer.Email, &customer.Name); err != nil {
		return model.Customer{}, fmt.Errorf("error getting customer: %w", err)
	}
	return customer, nil
}

func (r *customerRepository) CreateCustomer(ctx context.Context, customer model.Customer) (model.Customer, error) {
	customer.ID = gocql.TimeUUID()
	if err := r.session.Query("INSERT INTO merchandise.customers (customer_id, address, email, name) VALUES (?, ?, ?, ?)", customer.ID, customer.Address, customer.Email, customer.Name).Exec(); err != nil {
		return model.Customer{}, fmt.Errorf("error creating customer: %w", err)
	}

	return customer, nil
}
func (r *customerRepository) UpdateCustomer(ctx context.Context, id string, customer model.Customer) (model.Customer, error) {
	if err := r.session.Query("UPDATE merchandise.customers SET email = ?, name = ?, address = ? WHERE customer_id = ?", customer.Email, customer.Name, customer.Address, customer.ID).Exec(); err != nil {
		return model.Customer{}, fmt.Errorf("error updating customer: %w", err)
	}
	return customer, nil

}

func (r *customerRepository) DeleteCustomer(ctx context.Context, id string) error {
	if err := r.session.Query("DELETE FROM merchandise.customers WHERE customer_id = ?", id).Exec(); err != nil {
		return fmt.Errorf("error deleting customer: %w", err)
	}
	return nil
}
