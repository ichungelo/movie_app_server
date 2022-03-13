package transport

type CreateReviewRequest struct {
	Review string `json:"review" validate:"required"`
}