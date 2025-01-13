package repository

import (
	"github.com/gunktp20/digital-hubx-be/pkg/models"
)

func (r *classGormRepository) GetAllClasses(class_tier, keyword string, page int, limit int) (*[]models.Class, int64, error) {
	var classes []models.Class
	var total int64

	query := r.db.Model(&models.Class{})

	// Filter by keyword
	if keyword != "" {
		query = query.Where("title ILIKE ? OR description ILIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	// Filter by class_tier
	if class_tier != "" {
		query = query.Where("class_tier = ?", class_tier)
	}

	query.Count(&total)

	// คำนวณ OFFSET และเพิ่ม Limit และ Offset
	offset := (page - 1) * limit
	result := query.
		Preload("ClassCategory").
		Preload("ClassSessions").
		Limit(limit).   // กำหนด Limit
		Offset(offset). // กำหนด Offset
		Find(&classes)

	if result.Error != nil {
		return &[]models.Class{}, 0, result.Error
	}

	return &classes, total, nil
}
