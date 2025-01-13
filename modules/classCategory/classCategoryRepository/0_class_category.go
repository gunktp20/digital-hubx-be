package repository

import (
	"errors"

	classCategoryDto "github.com/gunktp20/digital-hubx-be/modules/classCategory/classCategoryDto"
	"github.com/gunktp20/digital-hubx-be/pkg/models"
)

type (
	ClassCategoryRepositoryService interface {
		CreateClassCategory(createClassCategoryReq *classCategoryDto.CreateClassCategoryReq) (*classCategoryDto.CreateClassCategoryRes, error)
		IsClassCategoryNameExists(classCategoryName string) bool
		GetClassCategoryById(categoryId string) (*models.ClassCategory, error)
		DeleteClassCategoryById(categoryId string) error
		GetAllClassCategories(keyword string, page int, limit int) (*[]models.ClassCategory, int64, error)
		IsClassCategoryIdExists(classCategoryId string) bool
	}
)

func (r *classCategoryGormRepository) CreateClassCategory(createClassCategoryReq *classCategoryDto.CreateClassCategoryReq) (*classCategoryDto.CreateClassCategoryRes, error) {

	classCategory := models.ClassCategory{
		CategoryName: createClassCategoryReq.ClassCategoryName,
	}

	if err := r.db.Create(&classCategory).Error; err != nil {
		return nil, err
	}

	return &classCategoryDto.CreateClassCategoryRes{
		ID:                classCategory.ID,
		ClassCategoryName: classCategory.CategoryName,
		CreatedAt:         classCategory.CreatedAt,
		UpdatedAt:         classCategory.UpdatedAt,
	}, nil
}

func (r *classCategoryGormRepository) IsClassCategoryNameExists(classCategoryName string) bool {
	var class = new(models.ClassCategory)
	result := r.db.First(&class, "category_name = ?", classCategoryName)

	if result.RowsAffected > 0 {
		return true
	}

	return false
}

func (r *classCategoryGormRepository) IsClassCategoryIdExists(classCategoryId string) bool {
	var class = new(models.ClassCategory)
	result := r.db.First(&class, "id = ?", classCategoryId)

	if result.RowsAffected > 0 {
		return true
	}

	return false
}

func (r *classCategoryGormRepository) GetClassCategoryById(categoryId string) (*models.ClassCategory, error) {
	var classCategory = new(models.ClassCategory)
	result := r.db.First(&classCategory, "id = ?", categoryId)

	if result.Error != nil {
		return &models.ClassCategory{}, result.Error
	}

	if result.RowsAffected == 0 {
		return &models.ClassCategory{}, errors.New("class category record not found")
	}

	return classCategory, nil
}

func (r *classCategoryGormRepository) DeleteClassCategoryById(categoryId string) error {
	var classCategory = new(models.ClassCategory)
	result := r.db.Delete(&classCategory, "id = ?", categoryId)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("class category record not found")
	}

	return nil
}

func (r *classCategoryGormRepository) GetAllClassCategories(keyword string, page int, limit int) (*[]models.ClassCategory, int64, error) {
	var classCategories []models.ClassCategory
	var total int64

	query := r.db.Model(&models.ClassCategory{})

	if keyword != "" {
		query = query.Where("category_name ILIKE ? ", "%"+keyword+"%")
	}

	query.Count(&total)

	offset := (page - 1) * limit
	result := query.
		Limit(limit).
		Offset(offset).
		Find(&classCategories)

	if result.Error != nil {
		return &[]models.ClassCategory{}, 0, result.Error
	}

	return &classCategories, total, nil
}
