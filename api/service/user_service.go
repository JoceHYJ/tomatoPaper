package service

import (
	"github.com/go-playground/validator/v10"
	"tomatoPaper/api/dao"
	"tomatoPaper/api/entity"
	"tomatoPaper/web"
)

// IUserService 定义接口
type IUserService interface {
	// Login(c *web.Context, dto entity.UserLoginDto)

	CreateUser(c *web.Context, dto entity.CreateUserDto)
	GetUserByUsername(c *web.Context, username string)
	DeleteUserById(c *web.Context, dto entity.UserIdDto)
}

// UserServiceImpl 实现接口
type UserServiceImpl struct {
}

//func (u UserServiceImpl) Login(c *web.Context, dto entity.UserLoginDto) {
//	//TODO implement me
//	panic("implement me")
//}

type UserResponse struct {
	ID       uint   `json:"id"`
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
		ID:       user.ID,
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

// DeleteUserById 根据id删除用户信息
func (u UserServiceImpl) DeleteUserById(c *web.Context, dto entity.UserIdDto) {
	dao.DeleteUserById(dto)
	c.RespJSON(200, "删除成功")
}

var userService = UserServiceImpl{}

func UserService() IUserService {
	return &userService
}
