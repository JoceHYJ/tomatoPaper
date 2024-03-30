package dao

// DAO 层
// Admin 数据层

import (
	"tomatoPaper/api/entity"
	"tomatoPaper/common/util"
	"tomatoPaper/pkg/database"
)

// AdminDetail 管理员详情
func AdminDetail(dto entity.AdminLoginDto) (admin entity.Admins) {
	stuId := dto.AdminID
	database.GormDB.Where("admin_id = ?", stuId).First(&admin)
	return admin
}

// GetAdminByAdminName 根据管理员名获取管理员
func GetAdminByAdminName(adminName string) (admin entity.Admins) {
	database.GormDB.Select("admin_id, admin_name").Where("admin_name = ?", adminName).First(&admin)
	return admin
}

// GetAdminByAdminID 根据管理员ID获取管理员
func GetAdminByAdminID(adminID string) (admin entity.Admins) {
	database.GormDB.Select("admin_id, admin_name").Where("admin_id = ?", adminID).First(&admin)
	return admin
}

// CreateAdmin 新增管理员
func CreateAdmin(dto entity.CreateAdminDto) bool {
	admins := entity.Admins{
		AdminID:   dto.AdminID,
		AdminName: dto.AdminName,
		//Password: dto.Password,
		Password: util.EncryptionMd5(dto.Password),
	}
	_ = database.GormDB.AutoMigrate(&admins)
	tx := database.GormDB.Create(&admins)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// DeleteAdminByAdminId 根据管理员id删除管理员
func DeleteAdminByAdminId(adminId string) bool {
	var count int64
	database.GormDB.Model(&entity.Admins{}).Where("admin_id = ?", adminId).Count(&count)
	if count <= 0 {
		return false
	}
	tx := database.GormDB.Where("admin_id = ?", adminId).Delete(&entity.Admins{})
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}
