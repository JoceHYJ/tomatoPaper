package dao

// DAO 层
// Teacher 数据层

import (
	"tomatoPaper/api/entity"
	"tomatoPaper/common/util"
	"tomatoPaper/pkg/database"
)

// TeacherDetail 教师详情
func TeacherDetail(dto entity.TeacherLoginDto) (teacher entity.Teachers) {
	stuId := dto.TeacherID
	database.GormDB.Where("teacher_id = ?", stuId).First(&teacher)
	return teacher
}

// GetTeacherByTeacherName 根据教师名获取教师
func GetTeacherByTeacherName(teacherName string) (teacher entity.Teachers) {
	database.GormDB.Select("teacher_id, teacher_name").Where("teacher_name = ?", teacherName).First(&teacher)
	return teacher
}

// GetTeacherByTeacherID 根据教师ID获取教师
func GetTeacherByTeacherID(teacherID string) (teacher entity.Teachers) {
	database.GormDB.Select("teacher_id, teacher_name").Where("teacher_id = ?", teacherID).First(&teacher)
	return teacher
}

// CreateTeacher 新增教师
func CreateTeacher(dto entity.CreateTeacherDto) bool {
	teachers := entity.Teachers{
		TeacherID:   dto.TeacherID,
		TeacherName: dto.TeacherName,
		//Password: dto.Password,
		Password: util.EncryptionMd5(dto.Password),
	}
	_ = database.GormDB.AutoMigrate(&teachers)
	tx := database.GormDB.Create(&teachers)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// DeleteTeacherByTeacherId 根据教师id删除教师
func DeleteTeacherByTeacherId(teacherId string) bool {
	var count int64
	database.GormDB.Model(&entity.Teachers{}).Where("teacher_id = ?", teacherId).Count(&count)
	if count <= 0 {
		return false
	}
	tx := database.GormDB.Where("teacher_id = ?", teacherId).Delete(&entity.Teachers{})
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}
