package dao

import (
	"tomatoPaper/api/entity"
	"tomatoPaper/pkg/database"
)

// GetCourseByCourseName 通过课程名称获取课程信息
func GetCourseByCourseName(name string) (course entity.Courses) {
	database.GormDB.Where("name = ?", name).Preload("Teacher").First(&course)
	return course
}

// GetCourseByCourseCode 通过课程代码获取课程信息
func GetCourseByCourseCode(code string) (course entity.Courses) {
	database.GormDB.Where("course_code = ?", code).Preload("Teacher").First(&course)
	return course
}

func GetCoursesByTeacherID(teacherId string) (courses []entity.Courses) {
	database.GormDB.Where("teacher_id = ?", teacherId).Preload("Teacher").Find(&courses)
	return courses
}

// CreateCourse 创建课程
func CreateCourse(dto entity.CreateCourseDto) bool {
	courses := entity.Courses{
		Name:        dto.Name,
		CourseCode:  dto.CourseCode,
		Description: dto.Description,
		TeacherID:   dto.TeacherID,
	}
	database.GormDB.AutoMigrate(&courses)
	tx := database.GormDB.Create(&courses)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// DeleteCourseByCourseCode 删除课程
func DeleteCourseByCourseCode(code string) bool {
	var count int64
	database.GormDB.Model(&entity.Courses{}).Where("course_code = ?", code).Count(&count)
	if count <= 0 {
		return false
	}
	tx := database.GormDB.Where("course_code = ?", code).Delete(&entity.Courses{})
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// UpdateCourse 更新课程信息
func UpdateCourse(dto entity.UpdateCourseDto) bool {
	courses := entity.Courses{
		Name:        dto.Name,
		CourseCode:  dto.CourseCode,
		Description: dto.Description,
		TeacherID:   dto.TeacherID,
	}
	var count int64
	database.GormDB.Model(&entity.Courses{}).Where("course_code = ?", courses.CourseCode).Count(&count)
	if count < 0 {
		return false
	}
	tx := database.GormDB.Model(&entity.Courses{}).Where("course_code=?", courses.CourseCode).Updates(&courses)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}
