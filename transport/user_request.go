package transport

type (
	RegisterRequest struct {
		Username        string `json:"username" validate:"required"`
		Email           string `json:"email" validate:"required,email"`
		FirstName       string `json:"first_name" validate:"required"`
		LastName        string `json:"last_name" validate:"required"`
		Password        string `json:"password" validate:"required,gte=8"`
		ConfirmPassword string `json:"confirm_password" validate:"required,gte=8"`
	}
	
	LoginRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
)
