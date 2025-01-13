package handler

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	classRegistrationDto "github.com/gunktp20/digital-hubx-be/modules/classRegistration/classRegistrationDto"
	classRegistrationUsecase "github.com/gunktp20/digital-hubx-be/modules/classRegistration/classRegistrationUsecase"
	"github.com/gunktp20/digital-hubx-be/pkg/response"
	"github.com/gunktp20/digital-hubx-be/pkg/utils"
)

type (
	ClassRegistrationHttpHandlerService interface {
		CreateClassRegistration(c *fiber.Ctx) error
		GetUserRegistrations(c *fiber.Ctx) error
	}

	classRegistrationHttpHandler struct {
		classRegistrationUsecase classRegistrationUsecase.ClassRegistrationUsecaseService
	}
)

func NewClassRegistrationHttpHandler(usecase classRegistrationUsecase.ClassRegistrationUsecaseService) ClassRegistrationHttpHandlerService {
	return &classRegistrationHttpHandler{classRegistrationUsecase: usecase}
}

func (h *classRegistrationHttpHandler) CreateClassRegistration(c *fiber.Ctx) error {
	var body classRegistrationDto.CreateClassRegistrationReq

	userIDStr, err := utils.GetUserIDFromContext(c)
	fmt.Print("=============================================================", userIDStr)
	if err != nil {
		return response.ErrResponse(c, http.StatusUnauthorized, err.Error(), nil)
	}

	// ? Merge fiber http body with dto struct
	if err := c.BodyParser(&body); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "The input data is invalid", nil)
	}

	// ? Validate field in body with dynamic function
	if err := validator.New().Struct(&body); err != nil {
		validationErrors := utils.TranslateValidationError(err.(validator.ValidationErrors))
		return response.ErrResponse(c, http.StatusBadRequest, "The input data is invalid", &validationErrors)
	}

	res, err := h.classRegistrationUsecase.CreateClassRegistration(&body, userIDStr)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *classRegistrationHttpHandler) GetUserRegistrations(c *fiber.Ctx) error {

	userIDStr, err := utils.GetUserIDFromContext(c)
	if err != nil {
		return response.ErrResponse(c, http.StatusUnauthorized, err.Error(), nil)
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	res, total, err := h.classRegistrationUsecase.GetUserRegistrations(userIDStr, page, limit)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return response.SuccessResponse(c, http.StatusOK, fiber.Map{
		"data":       res,
		"total":      total,
		"page":       page,
		"limit":      limit,
		"totalPages": int(math.Ceil(float64(total) / float64(limit))),
	})

}
