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

// DeleteUserById 删除用户
func DeleteUserById(c *web.Context) {
	var dto entity.UserIdDto
	_ = c.BindJson(&dto)
	service.UserService().DeleteUserById(c, dto)
}
