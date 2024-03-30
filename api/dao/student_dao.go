package dao

// DAO 层
// Student 数据层

import (
	"tomatoPaper/api/entity"
	"tomatoPaper/common/util"
	"tomatoPaper/pkg/database"
)

// StudentDetail 学生详情
func StudentDetail(dto entity.StudentLoginDto) (student entity.Students) {
	stuId := dto.StudentID
	database.GormDB.Where("student_id = ?", stuId).First(&student)
	return student
}

// GetStudentByStudentName 根据学生名获取学生
func GetStudentByStudentName(studentName string) (student entity.Students) {
	database.GormDB.Select("student_id, student_name").Where("student_name = ?", studentName).First(&student)
	return student
}

// GetStudentByStudentID 根据学生ID获取学生
func GetStudentByStudentID(studentID string) (student entity.Students) {
	database.GormDB.Select("student_id, student_name").Where("student_id = ?", studentID).First(&student)
	return student
}

// CreateStudent 新增学生
func CreateStudent(dto entity.CreateStudentDto) bool {
	students := entity.Students{
		StudentID:   dto.StudentID,
		StudentName: dto.StudentName,
		//Password: dto.Password,
		Password: util.EncryptionMd5(dto.Password),
	}
	_ = database.GormDB.AutoMigrate(&students)
	tx := database.GormDB.Create(&students)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// DeleteStudentByStudentId 根据学生id删除学生
func DeleteStudentByStudentId(studentId string) bool {
	var count int64
	database.GormDB.Model(&entity.Students{}).Where("student_id = ?", studentId).Count(&count)
	if count <= 0 {
		return false
	}
	tx := database.GormDB.Where("student_id = ?", studentId).Delete(&entity.Students{})
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}
