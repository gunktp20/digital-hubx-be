package repository

import (
	"gorm.io/gorm"
)

type questionGormRepository struct {
	db *gorm.DB
}

func NewQuestionGormRepository(db *gorm.DB) QuestionRepositoryService {
	return &questionGormRepository{db}
}
