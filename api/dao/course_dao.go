package dao

import (
	"tomatoPaper/api/entity"
	"tomatoPaper/pkg/database"
)

// GetCourseByCourseName 通过课程名称获取课程信息
func GetCourseByCourseName(name string) (course entity.Courses) {
	database.GormDB.Where("name = ?", name).First(&course)
	return course
}

// GetCourseByCourseCode 通过课程代码获取课程信息
func GetCourseByCourseCode(code string) (course entity.Courses) {
	database.GormDB.Where("course_code = ?", code).First(&course)
	return course
}

// CreateCourse 创建课程
func CreateCourse(dto entity.CreateCourseDto) bool {
	courses := entity.Courses{
		Name:        dto.Name,
		CourseCode:  dto.CourseCode,
		Description: dto.Description,
	}
	database.GormDB.AutoMigrate(&courses)
	tx := database.GormDB.Create(&courses)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}
