package controller

import (
	"tomatoPaper/api/entity"
	"tomatoPaper/api/service"
	"tomatoPaper/web"
)

// Controller 层
// 用户 控制层

// CreateStudent 新增用户
func CreateStudent(c *web.Context) {
	var dto entity.CreateStudentDto
	_ = c.BindJson(&dto)
	service.StudentService().CreateStudent(c, dto)
}

// GetStudentByStudentName 根据用户名获取用户信息
func GetStudentByStudentName(c *web.Context) {
	StudentName, _ := c.QueryValue("student_name").String()
	service.StudentService().GetStudentByStudentName(c, StudentName)
}

// GetStudentByStudentId 根据用户id获取用户信息
func GetStudentByStudentId(c *web.Context) {
	StudentId, _ := c.QueryValue("student_id").String()
	service.StudentService().GetStudentByStudentId(c, StudentId)
}

// DeleteStudentByStudentId 删除用户
func DeleteStudentByStudentId(c *web.Context) {
	StudentId, _ := c.QueryValue("student_id").String()
	service.StudentService().DeleteStudentByStudentId(c, StudentId)
}

// StudentLogin 用户登录
func StudentLogin(c *web.Context) {
	var dto entity.StudentLoginDto
	_ = c.BindJson(&dto)
	service.StudentService().StudentLogin(c, dto)
}
