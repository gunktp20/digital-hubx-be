package usecase

import (
	"errors"

	classCategoryDto "github.com/gunktp20/digital-hubx-be/modules/classCategory/classCategoryDto"
	classCategoryRepository "github.com/gunktp20/digital-hubx-be/modules/classCategory/classCategoryRepository"
	"github.com/gunktp20/digital-hubx-be/pkg/models"
)

type (
	ClassCategoryUsecaseService interface {
		CreateClassCategory(createClassCategoryReq *classCategoryDto.CreateClassCategoryReq) (*classCategoryDto.CreateClassCategoryRes, error)
		GetAllClassCategories(keyword string, page int, limit int) (*[]models.ClassCategory, int64, error)
	}

	classCategoryUsecase struct {
		classCategoryRepo classCategoryRepository.ClassCategoryRepositoryService
	}
)

func NewClassCategoryUsecase(classCategoryRepo classCategoryRepository.ClassCategoryRepositoryService) ClassCategoryUsecaseService {
	return &classCategoryUsecase{classCategoryRepo: classCategoryRepo}
}

func (u *classCategoryUsecase) CreateClassCategory(createClassCategoryReq *classCategoryDto.CreateClassCategoryReq) (*classCategoryDto.CreateClassCategoryRes, error) {

	// ? Check is new app group name is taken yet ?
	classCategoryTitleExists := u.classCategoryRepo.IsClassCategoryNameExists(createClassCategoryReq.ClassCategoryName)

	if classCategoryTitleExists {
		return &classCategoryDto.CreateClassCategoryRes{}, errors.New("class category name was taken")
	}

	return u.classCategoryRepo.CreateClassCategory(createClassCategoryReq)
}

func (u *classCategoryUsecase) GetAllClassCategories(keyword string, page int, limit int) (*[]models.ClassCategory, int64, error) {
	ClassCategories, total, err := u.classCategoryRepo.GetAllClassCategories(keyword, page, limit)
	if err != nil {
		return nil, 0, err
	}

	return ClassCategories, total, nil
}
