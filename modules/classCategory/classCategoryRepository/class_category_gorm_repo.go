package repository

import (
	"gorm.io/gorm"
)

type classCategoryGormRepository struct {
	db *gorm.DB
}

func NewClassCategoryGormRepository(db *gorm.DB) ClassCategoryRepositoryService {
	return &classCategoryGormRepository{db}
}
