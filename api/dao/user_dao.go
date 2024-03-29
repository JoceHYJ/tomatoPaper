package dao

// DAO 层
// 用户 数据层

import (
	"tomatoPaper/api/entity"
	"tomatoPaper/common/util"
	"tomatoPaper/pkg/database"
)

// UserDetail 用户详情
func UserDetail(dto entity.UserLoginDto) (user entity.Users) {
	userid := dto.UserID
	database.GormDB.Where("user_id = ?", userid).First(&user)
	return user
}

// GetUserByUsername 根据用户名获取用户
func GetUserByUsername(username string) (user entity.Users) {
	database.GormDB.Select("user_id, username, role").Where("username = ?", username).First(&user)
	return user
}

// GetUserByUserID 根据用户ID获取用户
func GetUserByUserID(userID string) (user entity.Users) {
	database.GormDB.Select("user_id, username, role").Where("user_id = ?", userID).First(&user)
	return user
}

// CreateUser 新增用户
func CreateUser(dto entity.CreateUserDto) bool {
	users := entity.Users{
		UserID:   dto.UserID,
		Username: dto.Username,
		//Password: dto.Password,
		Password: util.EncryptionMd5(dto.Password),
		Role:     dto.Role,
	}
	_ = database.GormDB.AutoMigrate(&users)
	tx := database.GormDB.Create(&users)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}

// DeleteUserByUserId 根据用户id删除用户
func DeleteUserByUserId(userid string) bool {
	var count int64
	database.GormDB.Model(&entity.Users{}).Where("user_id = ?", userid).Count(&count)
	if count <= 0 {
		return false
	}
	tx := database.GormDB.Where("user_id = ?", userid).Delete(&entity.Users{})
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}
