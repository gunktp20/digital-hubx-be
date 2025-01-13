package repository

import (
	"gorm.io/gorm"
)

type classRegistrationGormRepository struct {
	db *gorm.DB
}

func NewClassRegistrationGormRepository(db *gorm.DB) ClassRegistrationRepositoryService {
	return &classRegistrationGormRepository{db}
}
