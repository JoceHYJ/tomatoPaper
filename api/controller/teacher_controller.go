package controller

import (
	"tomatoPaper/api/entity"
	"tomatoPaper/api/service"
	"tomatoPaper/web"
)

// Controller 层
// 用户 控制层

// CreateTeacher 新增用户
func CreateTeacher(c *web.Context) {
	var dto entity.CreateTeacherDto
	_ = c.BindJson(&dto)
	service.TeacherService().CreateTeacher(c, dto)
}

// GetTeacherByTeacherName 根据用户名获取用户信息
func GetTeacherByTeacherName(c *web.Context) {
	TeacherName, _ := c.QueryValue("teacher_name").String()
	service.TeacherService().GetTeacherByTeacherName(c, TeacherName)
}

// GetTeacherByTeacherId 根据用户id获取用户信息
func GetTeacherByTeacherId(c *web.Context) {
	TeacherId, _ := c.QueryValue("teacher_id").String()
	service.TeacherService().GetTeacherByTeacherId(c, TeacherId)
}

// DeleteTeacherByTeacherId 删除用户
func DeleteTeacherByTeacherId(c *web.Context) {
	TeacherId, _ := c.QueryValue("teacher_id").String()
	service.TeacherService().DeleteTeacherByTeacherId(c, TeacherId)
}

// TeacherLogin 用户登录
func TeacherLogin(c *web.Context) {
	var dto entity.TeacherLoginDto
	_ = c.BindJson(&dto)
	service.TeacherService().TeacherLogin(c, dto)
}
