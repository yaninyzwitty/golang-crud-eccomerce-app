package repository

import (
	"context"
	"fmt"

	"github.com/gocql/gocql"
	"github.com/yaninyzwitty/crud-eccomerce-app/model"
)

type ReviewRepoSitory interface {
	GetReviews(ctx context.Context) ([]model.Review, error)
	GetReview(ctx context.Context, id string) (model.Review, error)
	CreateReview(ctx context.Context, review model.Review) (model.Review, error)
	UpdateReview(ctx context.Context, id string, review model.Review) (model.Review, error)
	DeleteReview(ctx context.Context, id string) error
}

type reviewRepository struct {
	ctx     context.Context
	session *gocql.Session
}

func NewReviewRepository(ctx context.Context, session *gocql.Session) ReviewRepoSitory {
	return &reviewRepository{
		ctx,
		session,
	}
}

func (r *reviewRepository) GetReviews(ctx context.Context) ([]model.Review, error) {
	var reviews []model.Review
	iter := r.session.Query(`SELECT review_id, comment, customer_id, product_id, rating FROM merchandise.reviews`).Iter()

	for {
		var review model.Review
		if !iter.Scan(&review.ID, &review.Comment, &review.CustomerID, &review.ProductID, &review.Rating) {
			break
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

func (r *reviewRepository) GetReview(ctx context.Context, id string) (model.Review, error) {
	var review model.Review

	if err := r.session.Query(`SELECT review_id, comment, customer_id, product_id, rating FROM merchandise.reviews WHERE review_id = ?`, id).Scan(&review.ID, &review.Comment, &review.CustomerID, &review.ProductID, &review.Rating); err != nil {
		return model.Review{}, fmt.Errorf("error getting review: %v", err)
	}
	return review, nil
}

func (r *reviewRepository) CreateReview(ctx context.Context, review model.Review) (model.Review, error) {
	review.ProductID = gocql.TimeUUID()
	review.ID = gocql.TimeUUID()
	review.CustomerID = gocql.TimeUUID()

	if err := r.session.Query(`INSERT INTO merchandise.reviews (review_id, comment, customer_id, product_id, rating) VALUES (?, ?, ?, ?, ?)`, review.ID, review.Comment, review.CustomerID, review.ProductID, review.Rating).Exec(); err != nil {
		return model.Review{}, fmt.Errorf("error creating review: %v", err)
	}

	return review, nil
}

func (r *reviewRepository) UpdateReview(ctx context.Context, id string, review model.Review) (model.Review, error) {
	if err := r.session.Query(`UPDATE merchandise.reviews SET comment = ?, customer_id = ?, product_id = ?, rating = ? WHERE review_id = ?`, review.Comment, review.CustomerID, review.ProductID, review.Rating, id).Exec(); err != nil {
		return model.Review{}, fmt.Errorf("error updating review: %v", err)
	}
	return review, nil
}

func (r *reviewRepository) DeleteReview(ctx context.Context, id string) error {
	if err := r.session.Query(`DELETE FROM merchandise.reviews WHERE review_id = ?`, id).Exec(); err != nil {
		return fmt.Errorf("error deleting review: %v", err)
	}
	return nil

}
