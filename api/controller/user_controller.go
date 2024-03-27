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
