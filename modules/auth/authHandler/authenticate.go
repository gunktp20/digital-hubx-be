package handler

import (
	"net/http"

	"github.com/go-playground/validator"
	authDto "github.com/gunktp20/digital-hubx-be/modules/auth/authDto"
	"github.com/gunktp20/digital-hubx-be/pkg/response"
	"github.com/gunktp20/digital-hubx-be/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// @Summary Authenticate
// @Description This endpoint allows you to authenticate
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body authDto.AuthenticationReq true "Authentication detail"
// @Success 200 {object} authDto.AuthenticationRes
// @Router /auth/login [post]
func (h *authHttpHandler) Authenticate(c *fiber.Ctx) error {
	var body authDto.AuthenticationReq

	// ? Merge fiber http body with dto struct
	if err := c.BodyParser(&body); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "The input data is invalid", nil)
	}
	// ? Validate field in body with dynamic function
	if err := validator.New().Struct(&body); err != nil {
		validationErrors := utils.TranslateValidationError(err.(validator.ValidationErrors))
		return response.ErrResponse(c, http.StatusBadRequest, "The input data is invalid", &validationErrors)
	}

	res, err := h.authUsecase.Authenticate(&body)
	if err != nil {
		return response.ErrResponse(c, http.StatusUnauthorized, err.Error(), nil)
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
