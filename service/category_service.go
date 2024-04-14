package service

import (
	"context"

	"github.com/yaninyzwitty/crud-eccomerce-app/model"
	"github.com/yaninyzwitty/crud-eccomerce-app/repository"
)

type CategoryService interface {
	GetCategories(ctx context.Context) ([]model.Category, error)
	GetCategory(ctx context.Context, id string) (model.Category, error)
	CreateCategory(ctx context.Context, category model.Category) (model.Category, error)
	DeleteCategory(ctx context.Context, id string) error
	UpdateCategory(ctx context.Context, id string, category model.Category) (model.Category, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo}
}

func (s *categoryService) GetCategories(ctx context.Context) ([]model.Category, error) {
	return s.repo.GetCategories(ctx)

}

func (s *categoryService) GetCategory(ctx context.Context, id string) (model.Category, error) {
	return s.repo.GetCategory(ctx, id)

}

func (s *categoryService) CreateCategory(ctx context.Context, category model.Category) (model.Category, error) {
	return s.repo.CreateCategory(ctx, category)

}

func (s *categoryService) UpdateCategory(ctx context.Context, id string, category model.Category) (model.Category, error) {
	return s.repo.UpdateCategory(ctx, id, category)

}

func (s *categoryService) DeleteCategory(ctx context.Context, id string) error {
	return s.repo.DeleteCategory(ctx, id)

}
