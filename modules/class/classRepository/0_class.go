package repository

import (
	classDto "github.com/gunktp20/digital-hubx-be/modules/class/classDto"
	"github.com/gunktp20/digital-hubx-be/pkg/models"
)

type (
	ClassRepositoryService interface {
		CreateClass(createClassReq *classDto.CreateClassReq, ThumbnailUrl string) (*classDto.CreateClassRes, error)
		IsClassTitleExists(classTitle string) bool
		GetAllClasses(class_tier, keyword string, page int, limit int) (*[]models.Class, int64, error)
		GetClassById(classId string) (*models.Class, error)
		DeleteClassById(classId string) (*models.Class, error)
	}
)
