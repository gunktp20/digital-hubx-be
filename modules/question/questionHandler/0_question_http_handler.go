package handler

import (
	"math"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	questionDto "github.com/gunktp20/digital-hubx-be/modules/question/questionDto"
	questionUsecase "github.com/gunktp20/digital-hubx-be/modules/question/questionUsecase"
	"github.com/gunktp20/digital-hubx-be/pkg/response"
	"github.com/gunktp20/digital-hubx-be/pkg/utils"
)

type (
	QuestionHttpHandlerService interface {
		CreateQuestion(c *fiber.Ctx) error
		GetQuestionsByClassID(c *fiber.Ctx) error
	}

	questionHttpHandler struct {
		questionUsecase questionUsecase.QuestionUsecaseService
	}
)

func NewQuestionHttpHandler(usecase questionUsecase.QuestionUsecaseService) QuestionHttpHandlerService {
	return &questionHttpHandler{questionUsecase: usecase}
}

func (h *questionHttpHandler) CreateQuestion(c *fiber.Ctx) error {

	var body questionDto.CreateQuestionReq

	// ? Merge fiber http body with dto struct
	if err := c.BodyParser(&body); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "The input data is invalid", nil)
	}

	// ? Validate field in body with dynamic function
	if err := validator.New().Struct(&body); err != nil {
		validationErrors := utils.TranslateValidationError(err.(validator.ValidationErrors))
		return response.ErrResponse(c, http.StatusBadRequest, "The input data is invalid", &validationErrors)
	}

	res, err := h.questionUsecase.CreateQuestion(&body)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *questionHttpHandler) GetQuestionsByClassID(c *fiber.Ctx) error {

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	res, total, err := h.questionUsecase.GetQuestionsByClassID(c.Params("class_id"), page, limit)

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
