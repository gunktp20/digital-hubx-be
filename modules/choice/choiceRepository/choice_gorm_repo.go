package repository

import (
	"gorm.io/gorm"
)

type choiceGormRepository struct {
	db *gorm.DB
}

func NewChoiceGormRepository(db *gorm.DB) ChoiceRepositoryService {
	return &choiceGormRepository{db}
}
