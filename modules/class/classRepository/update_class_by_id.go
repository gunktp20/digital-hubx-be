package repository

import (
	"errors"

	"github.com/gunktp20/digital-hubx-be/pkg/models"
)

func (r *classGormRepository) UpdateClassById(classId string) (*models.Class, error) {
	var class = new(models.Class)
	result := r.db.First(&class, "id = ?", classId)

	if result.Error != nil {
		return &models.Class{}, result.Error
	}

	if result.RowsAffected == 0 {
		return &models.Class{}, errors.New("class record not found")
	}

	return class, nil
}
