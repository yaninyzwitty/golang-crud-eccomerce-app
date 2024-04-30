package service

import (
	"context"

	"github.com/yaninyzwitty/crud-eccomerce-app/model"
	"github.com/yaninyzwitty/crud-eccomerce-app/repository"
)

type CustomerService interface {
	GetCustomers(ctx context.Context) ([]model.Customer, error)
	GetCustomer(ctx context.Context, id string) (model.Customer, error)
	CreateCustomer(ctx context.Context, customer model.Customer) (model.Customer, error)
	UpdateCustomer(ctx context.Context, id string, customer model.Customer) (model.Customer, error)
	DeleteCustomer(ctx context.Context, id string) error
}

type customerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{repo}
}

func (s *customerService) GetCustomers(ctx context.Context) ([]model.Customer, error) {
	return s.repo.GetCustomers(ctx)
}

func (s *customerService) GetCustomer(ctx context.Context, id string) (model.Customer, error) {
	return s.repo.GetCustomer(ctx, id)
}

func (s *customerService) CreateCustomer(ctx context.Context, customer model.Customer) (model.Customer, error) {
	return s.repo.CreateCustomer(ctx, customer)
}

func (s *customerService) UpdateCustomer(ctx context.Context, id string, customer model.Customer) (model.Customer, error) {
	return s.repo.UpdateCustomer(ctx, id, customer)
}

func (s *customerService) DeleteCustomer(ctx context.Context, id string) error {
	return s.repo.DeleteCustomer(ctx, id)
}
