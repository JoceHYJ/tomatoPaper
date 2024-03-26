package service

import (
	"tomatoPaper/api/dao"
	"tomatoPaper/api/entity"
)

type UserService struct {
	UserDao *dao.UserDao
}

func NewUserService(userDao *dao.UserDao) *UserService {
	return &UserService{UserDao: userDao}
}

func (s *UserService) CreateUser(user *entity.Users) error {
	return s.UserDao.CreateUser(user)
}
