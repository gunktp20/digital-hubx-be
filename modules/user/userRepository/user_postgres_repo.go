package repository

import (
	userDto "github.com/gunktp20/digital-hubx-be/modules/user/userDto"
	"github.com/gunktp20/digital-hubx-be/pkg/models"
	"gorm.io/gorm"
)

type userGormRepository struct {
	db *gorm.DB
}

func NewUserGormRepository(db *gorm.DB) UserRepositoryService {
	return &userGormRepository{db}
}

func (r *userGormRepository) IsUniqueUser(email string) bool {
	return true
}

func (r *userGormRepository) CreateOneUser(createUserReq *userDto.CreateUserReq) (*userDto.CreateUserRes, error) {

	user := &models.User{
		Email:    createUserReq.Email,
		Password: createUserReq.Password,
	}

	if err := r.db.Create(user).Error; err != nil {
		return &userDto.CreateUserRes{}, err
	}

	return &userDto.CreateUserRes{
		Email: user.Email,
	}, nil
}

func (r *userGormRepository) GetOneUserByEmail(email string) (*models.User, error) {

	var user models.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &models.User{}, err
		}
		return &models.User{}, err
	}

	return &models.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
