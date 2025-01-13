package usecase

import (
	"errors"
	"time"

	classDto "github.com/gunktp20/digital-hubx-be/modules/class/classDto"
	classRepository "github.com/gunktp20/digital-hubx-be/modules/class/classRepository"
	classCategoryRepository "github.com/gunktp20/digital-hubx-be/modules/classCategory/classCategoryRepository"
	"github.com/gunktp20/digital-hubx-be/pkg/bucket"
	"github.com/gunktp20/digital-hubx-be/pkg/models"
	"github.com/gunktp20/digital-hubx-be/pkg/utils"
)

type (
	ClassUsecaseService interface {
		CreateClass(createClassReq *classDto.CreateClassReq, fileBytes []byte) (*classDto.CreateClassRes, error)
		GetAllClasses(class_level, keyword string, page int, limit int) (*[]models.Class, int64, error)
		GetClassById(classId string) (*models.Class, error)
	}

	classUsecase struct {
		classRepo         classRepository.ClassRepositoryService
		classCategoryRepo classCategoryRepository.ClassCategoryRepositoryService
		bucketClient      bucket.BucketClientService
	}
)

func NewClassUsecase(classRepo classRepository.ClassRepositoryService, classCategoryRepo classCategoryRepository.ClassCategoryRepositoryService, bucketClient bucket.BucketClientService) ClassUsecaseService {
	return &classUsecase{classRepo: classRepo, classCategoryRepo: classCategoryRepo, bucketClient: bucketClient}
}

func (u *classUsecase) CreateClass(createClassReq *classDto.CreateClassReq, fileBytes []byte) (*classDto.CreateClassRes, error) {
	// ? Check is new app group name is taken yet ?
	classTitleExists := u.classRepo.IsClassTitleExists(createClassReq.Title)
	if classTitleExists {
		return &classDto.CreateClassRes{}, errors.New("class title was taken")
	}

	// ? Is class category id that user provided is exists
	if createClassReq.ClassCategoryID != "" {
		classCategoryExists := u.classCategoryRepo.IsClassCategoryIdExists(createClassReq.ClassCategoryID)
		if !classCategoryExists {
			return &classDto.CreateClassRes{}, errors.New("class category that you provided doesn't exists")
		}
	}

	// ? Get file extension from fileBytes
	fileExtension, err := utils.GetImageFileExtension(fileBytes)
	if err != nil {
		return &classDto.CreateClassRes{}, err
	}

	// ? Generate a unique file name
	fileName := utils.GenerateFileName(16)

	// ? Upload file to S3
	err = u.bucketClient.UploadFile(fileName, fileBytes, fileExtension)
	if err != nil {
		return &classDto.CreateClassRes{}, err
	}

	return u.classRepo.CreateClass(createClassReq, fileName)
}

func (u *classUsecase) GetAllClasses(class_tier, keyword string, page int, limit int) (*[]models.Class, int64, error) {
	classes, total, err := u.classRepo.GetAllClasses(class_tier, keyword, page, limit)
	if err != nil {
		return nil, 0, err
	}

	// Loop ผ่านแต่ละ AppGroup และอัปเดต IconURL ด้วย signed URL
	for i, class := range *classes {
		// เรียก GetSignedURL เพื่อดึง URL ที่มีเวลาใช้ชั่วคราว
		signedUrl, err := u.bucketClient.GetSignedURL(class.CoverImage, 60*time.Second)
		if err != nil {
			// ถ้ามีข้อผิดพลาดในการดึง signed URL ให้ข้ามไปตัวถัดไป
			continue
		}

		// อัปเดต IconURL ของ AppGroup ด้วย signed URL
		(*classes)[i].CoverImage = signedUrl
	}

	return classes, total, nil
}

func (u *classUsecase) GetClassById(classId string) (*models.Class, error) {
	res, err := u.classRepo.GetClassById(classId)
	if err != nil {
		return &models.Class{}, nil
	}

	signedUrl, err := u.bucketClient.GetSignedURL(res.CoverImage, 20*time.Second)
	if err != nil {
		return &models.Class{}, err
	}

	res.CoverImage = signedUrl
	return res, nil
}
