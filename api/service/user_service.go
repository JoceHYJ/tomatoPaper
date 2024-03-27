package service

import (
	"github.com/go-playground/validator/v10"
	"tomatoPaper/api/dao"
	"tomatoPaper/api/entity"
	"tomatoPaper/web"
)

// IUserService 定义接口
type IUserService interface {
	CreateUser(c *web.Context, dto entity.CreateUserDto)
}

// UserServiceImpl 实现接口
type UserServiceImpl struct {
}

func (u UserServiceImpl) CreateUser(c *web.Context, dto entity.CreateUserDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		c.RespJSON(400, err.Error())
		return
	}
	bool := dao.CreateUser(dto)
	if !bool {
		c.RespJSON(400, "创建失败")
	}
	c.RespJSON(200, "创建成功")
	return
}

var userService = UserServiceImpl{}

func UserService() IUserService {
	return &userService
}
