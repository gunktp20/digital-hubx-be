package repository

import (
	classDto "github.com/gunktp20/digital-hubx-be/modules/class/classDto"
	"github.com/gunktp20/digital-hubx-be/pkg/models"
)

func (r *classGormRepository) CreateClass(createClassReq *classDto.CreateClassReq, CoverImageUrl string) (*classDto.CreateClassRes, error) {

	class := models.Class{
		Title:           createClassReq.Title,
		Description:     createClassReq.Description,
		CoverImage:      CoverImageUrl,
		ClassCategoryID: createClassReq.ClassCategoryID,
		ClassLevel:      createClassReq.ClassLevel,
		ClassTier:       createClassReq.ClassTier,
	}

	if err := r.db.Create(&class).Error; err != nil {
		return &classDto.CreateClassRes{}, err
	}

	return &classDto.CreateClassRes{
		ID:              class.ID,
		Title:           class.Title,
		Description:     class.Description,
		CoverImage:      CoverImageUrl,
		ClassCategoryID: class.ClassCategoryID,
		ClassLevel:      class.ClassLevel,
		ClassTier:       class.ClassTier,
		IsActive:        class.IsActive,
		IsRemove:        class.IsRemove,
		CreatedAt:       class.CreatedAt,
		UpdatedAt:       class.UpdatedAt,
	}, nil
}
