package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yaninyzwitty/crud-eccomerce-app/model"
	"github.com/yaninyzwitty/crud-eccomerce-app/service"
)

type ReviewHandler struct {
	service service.ReviewService
}

func NewReviewHandler(service service.ReviewService) ReviewHandler {
	return ReviewHandler{service: service}
}

func (h *ReviewHandler) GetReviews(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reviews, err := h.service.GetReviews(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, reviews)

}
func (h *ReviewHandler) GetReview(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := r.Context()
	review, err := h.service.GetReview(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, review)

}

func (h *ReviewHandler) CreateReview(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var review model.Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	createdReview, err := h.service.CreateReview(ctx, review)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, createdReview)

}
func (h *ReviewHandler) UpdateReview(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := r.Context()
	var review model.Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	updatedReview, err := h.service.UpdateReview(ctx, id, review)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedReview)

}

func (h *ReviewHandler) DeleteReview(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	ctx := r.Context()
	err := h.service.DeleteReview(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")

}
