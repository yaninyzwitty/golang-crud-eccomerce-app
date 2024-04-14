package repository

import (
	"context"

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
	return []model.Category{}, nil

}
func (r *categoryRepository) GetCategory(ctx context.Context, id string) (model.Category, error) {
	return model.Category{}, nil

}
func (r *categoryRepository) CreateCategory(ctx context.Context, category model.Category) (model.Category, error) {
	return model.Category{}, nil

}

func (r *categoryRepository) UpdateCategory(ctx context.Context, id string, category model.Category) (model.Category, error) {
	return model.Category{}, nil

}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id string) error {
	return nil

}
