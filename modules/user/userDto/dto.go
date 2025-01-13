package dto

type (
	CreateUserReq struct {
		Email    string `json:"email" form:"email" validate:"required,email,max=255"`
		Password string `json:"password" form:"password" validate:"required,max=32"`
	}

	CreateUserRes struct {
		Email string `json:"email" form:"email" validate:"required,email,max=255"`
	}
)
