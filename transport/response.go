package transport

type (
	GeneralResponse struct {
		Success bool   `json:"success"`
		Message interface{} `json:"message"`
	}

	LoginResponse struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Token   string `json:"token"`
	}

	CreateReviewResponse struct {
		UserId string `json:"user_id" validate:"required"`
		MovieId int `json:"movie_id" validate:"required"`
		Review string `json:"review" validate:"required"`
	}
)
