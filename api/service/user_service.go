package service

import (
	"tomatoPaper/api/dao"
	"tomatoPaper/api/entity"
	"tomatoPaper/common/util"
	"tomatoPaper/pkg/jwt"
	"tomatoPaper/web"

	"github.com/go-playground/validator/v10"
)

// IUserService 定义接口
type IUserService interface {
	Login(c *web.Context, dto entity.UserLoginDto)
	CreateUser(c *web.Context, dto entity.CreateUserDto)
	GetUserByUsername(c *web.Context, username string)
	GetUserByUserId(c *web.Context, userid string)
	DeleteUserByUserId(c *web.Context, userid string)
}

// UserServiceImpl 实现接口
type UserServiceImpl struct {
}

// Login 登录
func (u UserServiceImpl) Login(c *web.Context, dto entity.UserLoginDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		c.RespJSON(400, "参数校验失败")
		return
	}
	user := dao.UserDetail(dto)
	//user := dao.GetUserByUserID(dto.UserID)
	hashedPassword := util.EncryptionMd5(dto.Password)
	//if !bytes.Equal([]byte(hashedPassword), []byte(user.Password)) {
	//	c.RespJSON(400, "密码错误")
	//	return
	//}
	//fmt.Println("输入密码:", dto.Password)
	//fmt.Println("哈希化后的输入密码:", hashedPassword)
	//fmt.Println("数据库中的密码:", user.Password)
	if hashedPassword != user.Password {
		c.RespJSON(400, "密码错误")
		return
	}
	token, err := jwt.GenerateToken(user)
	if err != nil {
		c.RespJSON(500, "生成token失败")
		return
	}
	c.RespJSON(200, map[string]any{
		"msg":   "登录成功",
		"token": token,
	})
}

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

// GetUserByUserId 根据用户id获取用户信息
func (u UserServiceImpl) GetUserByUserId(c *web.Context, userid string) {
	defer func() {
		err := recover()
		if err != nil {
			c.RespJSON(400, err)
		}
	}()
	user := dao.GetUserByUserID(userid)
	//c.RespJSON(200, user)
	resp := UserResponse{
		UserID:   user.UserID,
		Username: user.Username,
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
