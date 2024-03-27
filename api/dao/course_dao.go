package dao

import (
	"tomatoPaper/api/entity"
	"tomatoPaper/pkg/database"
)

// CreateCourse 创建课程
func CreateCourse(dto entity.CreateCourseDto) bool {
	courses := entity.Courses{
		Name: dto.Name,
		//Deadline:    dto.Deadline,
		Description: dto.Description,
	}
	database.GormDB.AutoMigrate(&courses)
	tx := database.GormDB.Create(&courses)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}
