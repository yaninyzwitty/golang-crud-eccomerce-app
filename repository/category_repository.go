package repository

import (
	"context"
	"fmt"

	"github.com/gocql/gocql"
	"github.com/yaninyzwitty/crud-eccomerce-app/model"
)

type CategoryRepository interface {
	GetCategories(ctx context.Context) ([]model.Category, error)
	GetCategory(ctx context.Context, id string) (model.Category, error)
	CreateCategory(ctx context.Context, category model.Category) (model.Category, error)
	UpdateCategory(ctx context.Context, id string, category model.Category) (model.Category, error)
	DeleteCategory(ctx context.Context, id string) error
}

type categoryRepository struct {
	ctx     context.Context
	session *gocql.Session
}

func NewCategoryRepository(ctx context.Context, session *gocql.Session) CategoryRepository {
	return &categoryRepository{
		ctx, session,
	}
}

func (r *categoryRepository) GetCategories(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category
	iter := r.session.Query("SELECT category_id, name, description FROM merchandise.categories").Iter()

	for {
		var category model.Category
		if !iter.Scan(&category.ID, &category.Name, &category.Description) {
			break // no more rows
		}

		categories = append(categories, category)
	}

	return categories, nil

}
func (r *categoryRepository) GetCategory(ctx context.Context, id string) (model.Category, error) {
	var category model.Category

	if err := r.session.Query("SELECT category_id, name, description FROM merchandise.categories WHERE category_id = ?", id).Scan(&category.ID, &category.Name, &category.Description); err != nil {
		return model.Category{}, fmt.Errorf("failed to get categories: %w", err)
	}

	return category, nil

}
func (r *categoryRepository) CreateCategory(ctx context.Context, category model.Category) (model.Category, error) {
	category.ID = gocql.TimeUUID()

	err := r.session.Query("INSERT INTO merchandise.categories (category_id, name, description) VALUES (?, ?, ?)", category.ID, category.Name, category.Description).Exec()
	if err != nil {
		return model.Category{}, fmt.Errorf("failed to create category: %w", err)
	}

	return category, nil

}

func (r *categoryRepository) UpdateCategory(ctx context.Context, id string, category model.Category) (model.Category, error) {
	if err := r.session.Query("UPDATE merchandise.categories SET name = ?, description = ? WHERE category_id = ?", category.Name, category.Description, id).Exec(); err != nil {
		return model.Category{}, fmt.Errorf("failed to update category: %w", err)
	}

	return category, nil

}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id string) error {
	if err := r.session.Query("DELETE FROM merchandise.categories WHERE category_id = ?", id).Exec(); err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}
	return nil

}
