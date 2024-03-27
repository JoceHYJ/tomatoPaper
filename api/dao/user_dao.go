package dao

// DAO 层
// 用户 数据层

import (
	"tomatoPaper/api/entity"
	"tomatoPaper/pkg/database"
)

// UserDetail 用户详情
func UserDetail(dto entity.UserLoginDto) (user entity.Users) {
	username := dto.Username
	database.GormDB.Where("username = ?", username).First(&user)
	return user
}

// GetUserByUsername 根据用户名获取用户
func GetUserByUsername(username string) (user entity.Users) {
	database.GormDB.Select("id, username, role").Where("username = ?", username).First(&user)
	return user
}

// CreateUser 新增用户
func CreateUser(dto entity.CreateUserDto) bool {
	users := entity.Users{
		Username: dto.Username,
		Password: dto.Password,
		// TODO: 密码
		//Password: util.EncryptionMd5(dto.Password),
		Role: dto.Role,
	}
	_ = database.GormDB.AutoMigrate(&users)
	tx := database.GormDB.Create(&users)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}
