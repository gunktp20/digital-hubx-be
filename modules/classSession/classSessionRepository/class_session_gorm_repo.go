package repository

import (
	"gorm.io/gorm"
)

type classSessionGormRepository struct {
	db *gorm.DB
}

func NewClassSessionGormRepository(db *gorm.DB) ClassSessionRepositoryService {
	return &classSessionGormRepository{db}
}
