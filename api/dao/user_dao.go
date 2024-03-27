package dao

// DAO 层
// 用户 数据层

import (
	"tomatoPaper/api/entity"
	"tomatoPaper/pkg/database"
)

// CreateUser 新增用户
func CreateUser(dto entity.CreateUserDto) bool {
	users := entity.Users{
		Username: dto.Username,
		Password: dto.Password,
		// TODO: 密码
		//Password: util.EncryptionMd5(dto.Password),
		Role: dto.Role,
	}
	database.GormDB.AutoMigrate(&users)
	tx := database.GormDB.Create(&users)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}
