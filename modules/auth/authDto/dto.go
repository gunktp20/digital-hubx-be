package dto

type (
	AuthenticationReq struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	AuthenticationRes struct {
		AccessToken string `json:"access_token" validate:"required"`
		Email       string `json:"email" validate:"required"`
		UserID      string `json:"user_id" validate:"required"`
	}
)
