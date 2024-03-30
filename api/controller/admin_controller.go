package controller

import (
	"tomatoPaper/api/entity"
	"tomatoPaper/api/service"
	"tomatoPaper/web"
)

// Controller 层
// 管理员 控制层

// CreateAdmin 新增管理员
func CreateAdmin(c *web.Context) {
	var dto entity.CreateAdminDto
	_ = c.BindJson(&dto)
	service.AdminService().CreateAdmin(c, dto)
}

// GetAdminByAdminName 根据管理员名获取管理员信息
func GetAdminByAdminName(c *web.Context) {
	AdminName, _ := c.QueryValue("admin_name").String()
	service.AdminService().GetAdminByAdminName(c, AdminName)
}

// GetAdminByAdminId 根据管理员id获取管理员信息
func GetAdminByAdminId(c *web.Context) {
	AdminId, _ := c.QueryValue("admin_id").String()
	service.AdminService().GetAdminByAdminId(c, AdminId)
}

// DeleteAdminByAdminId 删除管理员
func DeleteAdminByAdminId(c *web.Context) {
	AdminId, _ := c.QueryValue("admin_id").String()
	service.AdminService().DeleteAdminByAdminId(c, AdminId)
}

// AdminLogin 管理员登录
func AdminLogin(c *web.Context) {
	var dto entity.AdminLoginDto
	_ = c.BindJson(&dto)
	service.AdminService().AdminLogin(c, dto)
}
