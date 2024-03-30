package controller

import (
	"tomatoPaper/api/entity"
	"tomatoPaper/api/service"
	"tomatoPaper/web"
)

// Controller 层
// 用户 控制层

// CreateUser 新增用户
func CreateUser(c *web.Context) {
	var dto entity.CreateUserDto
	_ = c.BindJson(&dto)
	service.UserService().CreateUser(c, dto)
}

// GetUserByUsername 根据用户名获取用户信息
func GetUserByUsername(c *web.Context) {
	Username, _ := c.QueryValue("username").String()
	service.UserService().GetUserByUsername(c, Username)
}

// GetUserByUserId 根据用户id获取用户信息
func GetUserByUserId(c *web.Context) {
	UserId, _ := c.QueryValue("user_id").String()
	service.UserService().GetUserByUserId(c, UserId)
}

// DeleteUserByUserId 删除用户
func DeleteUserByUserId(c *web.Context) {
	UserId, _ := c.QueryValue("user_id").String()
	service.UserService().DeleteUserByUserId(c, UserId)
}

// Login 用户登录
func Login(c *web.Context) {
	var dto entity.UserLoginDto
	_ = c.BindJson(&dto)
	service.UserService().Login(c, dto)
}
