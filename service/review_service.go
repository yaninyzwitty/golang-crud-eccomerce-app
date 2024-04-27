package service

import (
	"context"

	"github.com/yaninyzwitty/crud-eccomerce-app/model"
	"github.com/yaninyzwitty/crud-eccomerce-app/repository"
)

type ReviewService interface {
	GetReviews(ctx context.Context) ([]model.Review, error)
	GetReview(ctx context.Context, id string) (model.Review, error)
	CreateReview(ctx context.Context, review model.Review) (model.Review, error)
	UpdateReview(ctx context.Context, id string, review model.Review) (model.Review, error)
	DeleteReview(ctx context.Context, id string) error
}

type reviewService struct {
	repo repository.ReviewRepoSitory
}

func NewReviewService(repo repository.ReviewRepoSitory) ReviewService {
	return &reviewService{repo}
}

func (s *reviewService) GetReviews(ctx context.Context) ([]model.Review, error) {
	return s.repo.GetReviews(ctx)

}
func (s *reviewService) GetReview(ctx context.Context, id string) (model.Review, error) {
	return s.repo.GetReview(ctx, id)

}

func (s *reviewService) CreateReview(ctx context.Context, review model.Review) (model.Review, error) {
	return s.repo.CreateReview(ctx, review)

}

func (s *reviewService) UpdateReview(ctx context.Context, id string, review model.Review) (model.Review, error) {
	return s.repo.UpdateReview(ctx, id, review)

}
func (s *reviewService) DeleteReview(ctx context.Context, id string) error {
	return s.repo.DeleteReview(ctx, id)

}
