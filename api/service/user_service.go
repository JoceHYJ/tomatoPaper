package service

import (
	"tomatoPaper/api/dao"
	"tomatoPaper/api/entity"
	"tomatoPaper/web"

	"github.com/go-playground/validator/v10"
)

// IUserService 定义接口
type IUserService interface {
	// Login(c *web.Context, dto entity.UserLoginDto)

	CreateUser(c *web.Context, dto entity.CreateUserDto)
	GetUserByUsername(c *web.Context, username string)
	DeleteUserByUserId(c *web.Context, userid string)
}

// UserServiceImpl 实现接口
type UserServiceImpl struct {
}

//func (u UserServiceImpl) Login(c *web.Context, dto entity.UserLoginDto) {
//	//TODO implement me
//	panic("implement me")
//}

type UserResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

// GetUserByUsername 根据用户名获取用户信息
func (u UserServiceImpl) GetUserByUsername(c *web.Context, username string) {
	defer func() {
		err := recover()
		if err != nil {
			c.RespJSON(400, err)
		}
	}()

	user := dao.GetUserByUsername(username)
	//c.RespJSON(200, user)
	resp := UserResponse{
		UserID:   user.UserID,
		Username: username,
		Role:     entity.RoleMap[user.Role],
	}
	c.RespJSON(200, resp)
	return
}

// CreateUser 创建用户
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

// DeleteUserByUserId 根据id删除用户信息
func (u UserServiceImpl) DeleteUserByUserId(c *web.Context, userid string) {
	defer func() {
		err := recover()
		if err != nil {
			c.RespJSON(400, err)
		}
	}()
	bool := dao.DeleteUserByUserId(userid)
	if !bool {
		c.RespJSON(400, "删除失败")
	}
	c.RespJSON(200, "删除成功")
}

var userService = UserServiceImpl{}

func UserService() IUserService {
	return &userService
}
