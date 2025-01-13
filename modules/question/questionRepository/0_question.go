package repository

import (
	"errors"

	questionDto "github.com/gunktp20/digital-hubx-be/modules/question/questionDto"
	"github.com/gunktp20/digital-hubx-be/pkg/models"
)

type (
	QuestionRepositoryService interface {
		CreateQuestion(createQuestionReq *questionDto.CreateQuestionReq) (*questionDto.CreateQuestionRes, error)
		GetQuestionsByClassID(classId string, page int, limit int) (*[]models.Question, int64, error)
		GetQuestionById(questionID string) (*models.Question, error)
	}
)

func (r *questionGormRepository) CreateQuestion(createQuestionReq *questionDto.CreateQuestionReq) (*questionDto.CreateQuestionRes, error) {

	question := models.Question{
		Description:  createQuestionReq.Description,
		ClassID:      createQuestionReq.ClassID,
		QuestionType: models.QuestionTypes(createQuestionReq.QuestionType),
	}

	if err := r.db.Create(&question).Error; err != nil {
		return &questionDto.CreateQuestionRes{}, err
	}

	if err := r.db.Preload("Class").Preload("Choices").First(&question).Error; err != nil {
		return &questionDto.CreateQuestionRes{}, err
	}

	return &questionDto.CreateQuestionRes{
		ID:          question.ID,
		Description: question.Description,
		ClassID:     question.ClassID,
		Class: questionDto.QuestionClass{
			Title:       question.Class.Title,
			Description: question.Class.Description,
			ClassTier:   question.Class.ClassTier,
			ClassLevel:  question.Class.ClassLevel,
		},
		QuestionType: question.QuestionType,
		CreatedAt:    question.CreatedAt,
		UpdatedAt:    question.UpdatedAt,

		Choices: question.Choices,
	}, nil
}

func (r *questionGormRepository) GetQuestionsByClassID(classId string, page int, limit int) (*[]models.Question, int64, error) {
	var questions []models.Question
	var total int64

	query := r.db.Model(&models.Question{})

	query.Count(&total)

	result := query.
		Preload("Choices").
		Find(&questions, "class_id = ?", classId)

	if result.Error != nil {
		return &[]models.Question{}, 0, result.Error
	}

	return &questions, total, nil

}

func (r *questionGormRepository) GetQuestionById(questionID string) (*models.Question, error) {
	var question = new(models.Question)
	result := r.db.First(&question, "id = ?", questionID)

	if result.Error != nil {
		return &models.Question{}, result.Error
	}

	if result.RowsAffected == 0 {
		return &models.Question{}, errors.New("question record not found")
	}

	return question, nil
}
