package repository

import (
	choiceDto "github.com/gunktp20/digital-hubx-be/modules/choice/choiceDto"
	"github.com/gunktp20/digital-hubx-be/pkg/models"
)

type (
	ChoiceRepositoryService interface {
		CreateChoice(createChoiceReq *choiceDto.CreateChoiceReq) (*choiceDto.CreateChoiceRes, error)
		GetChoicesByClassID(classId string, page int, limit int) (*[]models.Choice, int64, error)
	}
)

func (r *choiceGormRepository) CreateChoice(createChoiceReq *choiceDto.CreateChoiceReq) (*choiceDto.CreateChoiceRes, error) {

	choice := models.Choice{
		Description: createChoiceReq.Description,
		QuestionID:  createChoiceReq.QuestionID,
	}

	if err := r.db.Create(&choice).Error; err != nil {
		return &choiceDto.CreateChoiceRes{}, err
	}

	return &choiceDto.CreateChoiceRes{
		ID:          choice.ID,
		Description: choice.Description,
		QuestionID:  choice.QuestionID,
	}, nil
}

func (r *choiceGormRepository) GetChoicesByClassID(classId string, page int, limit int) (*[]models.Choice, int64, error) {
	var choices []models.Choice
	var total int64

	query := r.db.Model(&models.Choice{})

	query.Count(&total)

	result := query.
		Preload("Choices").
		Find(&choices, "class_id = ?", classId)

	if result.Error != nil {
		return &[]models.Choice{}, 0, result.Error
	}

	return &choices, total, nil

}
