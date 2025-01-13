package repository

import (
	"gorm.io/gorm"
)

type classGormRepository struct {
	db *gorm.DB
}

func NewClassGormRepository(db *gorm.DB) ClassRepositoryService {
	return &classGormRepository{db}
}
