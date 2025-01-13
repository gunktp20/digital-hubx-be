package handler

import (
	"math"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	classDto "github.com/gunktp20/digital-hubx-be/modules/class/classDto"
	classUsecase "github.com/gunktp20/digital-hubx-be/modules/class/classUsecase"
	"github.com/gunktp20/digital-hubx-be/pkg/models"
	"github.com/gunktp20/digital-hubx-be/pkg/response"
	"github.com/gunktp20/digital-hubx-be/pkg/utils"
)

type (
	ClassHttpHandlerService interface {
		CreateClass(c *fiber.Ctx) error
		GetAllClasses(c *fiber.Ctx) error
		GetClassById(c *fiber.Ctx) error
	}

	classHttpHandler struct {
		classUsecase classUsecase.ClassUsecaseService
	}
)

func NewClassHttpHandler(usecase classUsecase.ClassUsecaseService) ClassHttpHandlerService {
	return &classHttpHandler{classUsecase: usecase}
}

func (h *classHttpHandler) CreateClass(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("cover_image")
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "Class cover image is required", nil)
	}

	// ? Convert Multipart file to bytes
	fileBytes, err := utils.ConvertMultipartFileToBytes(fileHeader)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, "Failed to convert multipart file to bytes", nil)
	}
	// ? Define allowed types & max file size
	allowedTypes := []string{"image/png", "image/jpg"}
	maxFileSize := int64(5 * 1024 * 1024)

	classLevelStr := c.FormValue("class_level")
	var classLevel int
	// ? If class level is sent strconv atoi for check class level is an integer
	if classLevelStr != "" {
		classLevel, err = strconv.Atoi(classLevelStr)
		if err != nil {
			return response.ErrResponse(c, http.StatusBadRequest, "Invalid class_level, must be an integer", nil)
		}
	}

	body := classDto.CreateClassReq{
		Title:           c.FormValue("title"),
		Description:     c.FormValue("description"),
		CoverImage:      "",
		ClassCategoryID: c.FormValue("class_category_id"),
		ClassLevel:      classLevel, // classLevel จะเป็น 0 ถ้าไม่ได้รับค่า
		ClassTier:       models.ClassTier(c.FormValue("class_tier")),
	}

	// ? Validate file with allowed types & max file size
	if err := utils.ValidateFile(fileBytes, allowedTypes, maxFileSize); err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error(), nil)
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

	res, err := h.classUsecase.CreateClass(&body, fileBytes)
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *classHttpHandler) GetAllClasses(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	res, total, err := h.classUsecase.GetAllClasses(c.Query("class_tier"), c.Query("keyword"), page, limit)
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

func (h *classHttpHandler) GetClassById(c *fiber.Ctx) error {
	res, err := h.classUsecase.GetClassById(c.Params("class_id"))
	if err != nil {
		return response.ErrResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
