package service

import (
	"context"

	"github.com/yaninyzwitty/crud-eccomerce-app/model"
	"github.com/yaninyzwitty/crud-eccomerce-app/repository"
)

type ProductService interface {
	GetProducts(ctx context.Context) ([]model.Product, error)
	GetProduct(ctx context.Context, id string) (model.Product, error)
	CreateProduct(ctx context.Context, product model.Product) (model.Product, error)
	UpdateProduct(ctx context.Context, id string, product model.Product) (model.Product, error)
	DeleteProduct(ctx context.Context, id string) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) GetProducts(ctx context.Context) ([]model.Product, error) {
	return s.repo.GetProducts(ctx)

}

func (s *productService) GetProduct(ctx context.Context, id string) (model.Product, error) {
	return s.repo.GetProduct(ctx, id)

}

func (s *productService) CreateProduct(ctx context.Context, product model.Product) (model.Product, error) {
	return s.repo.CreateProduct(ctx, product)

}

func (s *productService) UpdateProduct(ctx context.Context, id string, product model.Product) (model.Product, error) {
	return s.repo.UpdateProduct(ctx, id, product)

}

func (s *productService) DeleteProduct(ctx context.Context, id string) error {
	return s.repo.DeleteProduct(ctx, id)
}
