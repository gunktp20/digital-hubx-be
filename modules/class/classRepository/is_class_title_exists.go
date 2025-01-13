package repository

import "github.com/gunktp20/digital-hubx-be/pkg/models"

func (r *classGormRepository) IsClassTitleExists(classTitle string) bool {

	var class = new(models.Class)
	result := r.db.First(&class, "title = ?", classTitle)

	if result.RowsAffected > 0 {
		return true
	}

	return false
}
