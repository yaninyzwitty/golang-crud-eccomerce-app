package repository

import (
	"context"
	"fmt"

	"github.com/gocql/gocql"
	"github.com/yaninyzwitty/crud-eccomerce-app/model"
)

type ProductRepository interface {
	GetProducts(ctx context.Context) ([]model.Product, error)
	GetProduct(ctx context.Context, productId string) (model.Product, error)
	CreateProduct(ctx context.Context, product model.Product) (model.Product, error)
	UpdateProduct(ctx context.Context, productId string, product model.Product) (model.Product, error)
	DeleteProduct(ctx context.Context, productId string) error
}

type productRepository struct {
	ctx     context.Context
	session *gocql.Session
}

func NewProductRepository(ctx context.Context, session *gocql.Session) ProductRepository {
	return &productRepository{ctx, session}
}

func (r *productRepository) GetProducts(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	iter := r.session.Query("SELECT product_id, name, description, price, quantity FROM merchandise.products").Iter()
	for {
		var product model.Product
		if !iter.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity) {
			break // no more rows

		}

		products = append(products, product)
	}

	if err := iter.Close(); err != nil {
		return []model.Product{}, err
	}

	return products, nil

}
func (r *productRepository) GetProduct(ctx context.Context, productId string) (model.Product, error) {
	var product model.Product

	if err := r.session.Query("SELECT product_id, name, description, price, quantity FROM merchandise.products WHERE product_id = ?", productId).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity); err != nil {
		return model.Product{}, fmt.Errorf("error getting product: %w", err)
	}

	return product, nil

}

func (r *productRepository) CreateProduct(ctx context.Context, product model.Product) (model.Product, error) {
	product.ID = gocql.TimeUUID()

	err := r.session.Query("INSERT INTO merchandise.products (product_id, name, description, price, quantity) VALUES (?, ?, ?, ?, ?)", product.ID, product.Name, product.Description, product.Price, product.Quantity).Exec()
	if err != nil {
		return model.Product{}, fmt.Errorf("error creating product: %w", err)
	}

	return product, nil

}

func (r *productRepository) UpdateProduct(ctx context.Context, productId string, product model.Product) (model.Product, error) {

	if err := r.session.Query("UPDATE merchandise.products SET name = ?, description = ?, price = ?, quantity = ? WHERE product_id = ?", product.Name, product.Description, product.Price, product.Quantity, product.ID).Exec(); err != nil {
		return model.Product{}, fmt.Errorf("error updating product: %w", err)
	}

	return product, nil

}

func (r *productRepository) DeleteProduct(ctx context.Context, productId string) error {
	if err := r.session.Query("DELETE FROM merchandise.products WHERE product_id = ?", productId).Exec(); err != nil {
		return fmt.Errorf("error deleting product: %w", err)

	}

	return nil

}
