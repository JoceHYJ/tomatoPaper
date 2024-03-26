package dao

import (
	"gorm.io/gorm"
	"tomatoPaper/api/entity"
)

type UserDao struct {
	DB *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{DB: db}
}

func (dao *UserDao) CreateUser(user *entity.Users) error {
	return dao.DB.Create(user).Error
}
