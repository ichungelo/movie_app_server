package transport

type (
	CreateReviewRequest struct {
		Review string `json:"review" validate:"required"`
	}

	CreatePostReviewRequest struct {
		UserId  string `json:"user_id" validate:"required"`
		MovieId int    `json:"movie_id" validate:"required"`
		Review  string `json:"review" validate:"required"`
	}

	PutReviewRequest struct {
		ReviewId int    `json:"review_id" validate:"required"`
		UserId   string `json:"user_id" validate:"required"`
		MovieId  int    `json:"movie_id" validate:"required"`
		Review   string `json:"review" validate:"required"`
	}
)
