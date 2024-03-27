package controller

import (
	"tomatoPaper/api/entity"
	"tomatoPaper/api/service"
	"tomatoPaper/web"
)

// CreateCourse 创建课程
func CreateCourse(c *web.Context) {
	var dto entity.CreateCourseDto
	_ = c.BindJson(&dto)
	service.CourseService().CreateCourse(c, dto)
}
