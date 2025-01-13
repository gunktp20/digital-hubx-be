package usecase

import (
	classRepository "github.com/gunktp20/digital-hubx-be/modules/class/classRepository"
	questionDto "github.com/gunktp20/digital-hubx-be/modules/question/questionDto"
	questionRepository "github.com/gunktp20/digital-hubx-be/modules/question/questionRepository"
	"github.com/gunktp20/digital-hubx-be/pkg/models"
)

type (
	QuestionUsecaseService interface {
		CreateQuestion(createQuestionReq *questionDto.CreateQuestionReq) (*questionDto.CreateQuestionRes, error)
		GetQuestionsByClassID(classId string, page int, limit int) (*[]models.Question, int64, error)
	}

	questionUsecase struct {
		questionRepo questionRepository.QuestionRepositoryService
		classRepo    classRepository.ClassRepositoryService
		// bucketClient bucket.BucketClientService
	}
)

func NewQuestionUsecase(questionRepo questionRepository.QuestionRepositoryService, classRepo classRepository.ClassRepositoryService) QuestionUsecaseService {
	return &questionUsecase{questionRepo: questionRepo, classRepo: classRepo}
}

func (u *questionUsecase) CreateQuestion(createQuestionReq *questionDto.CreateQuestionReq) (*questionDto.CreateQuestionRes, error) {

	_, err := u.classRepo.GetClassById(createQuestionReq.ClassID)
	if err != nil {
		return &questionDto.CreateQuestionRes{}, err
	}

	return u.questionRepo.CreateQuestion(createQuestionReq)
}

func (u *questionUsecase) GetQuestionsByClassID(classId string, page int, limit int) (*[]models.Question, int64, error) {

	res, total, err := u.questionRepo.GetQuestionsByClassID(classId, page, limit)
	if err != nil {
		return &[]models.Question{}, total, nil
	}

	return res, total, nil
}
