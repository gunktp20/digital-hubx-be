package usecase

import (
	choiceDto "github.com/gunktp20/digital-hubx-be/modules/choice/choiceDto"
	choiceRepository "github.com/gunktp20/digital-hubx-be/modules/choice/choiceRepository"
	questionRepository "github.com/gunktp20/digital-hubx-be/modules/question/questionRepository"
	"github.com/gunktp20/digital-hubx-be/pkg/models"
)

type (
	ChoiceUsecaseService interface {
		CreateChoice(createChoiceReq *choiceDto.CreateChoiceReq) (*choiceDto.CreateChoiceRes, error)
		GetChoicesByClassID(classId string, page int, limit int) (*[]models.Choice, int64, error)
	}

	choiceUsecase struct {
		choiceRepo   choiceRepository.ChoiceRepositoryService
		questionRepo questionRepository.QuestionRepositoryService
		// bucketClient bucket.BucketClientService
	}
)

func NewChoiceUsecase(choiceRepo choiceRepository.ChoiceRepositoryService, questionRepo questionRepository.QuestionRepositoryService) ChoiceUsecaseService {
	return &choiceUsecase{choiceRepo: choiceRepo, questionRepo: questionRepo}
}

func (u *choiceUsecase) CreateChoice(createChoiceReq *choiceDto.CreateChoiceReq) (*choiceDto.CreateChoiceRes, error) {

	_, err := u.questionRepo.GetQuestionById(createChoiceReq.QuestionID)
	if err != nil {
		return &choiceDto.CreateChoiceRes{}, err
	}

	return u.choiceRepo.CreateChoice(createChoiceReq)
}

func (u *choiceUsecase) GetChoicesByClassID(classId string, page int, limit int) (*[]models.Choice, int64, error) {

	res, total, err := u.choiceRepo.GetChoicesByClassID(classId, page, limit)
	if err != nil {
		return &[]models.Choice{}, total, nil
	}

	return res, total, nil
}
