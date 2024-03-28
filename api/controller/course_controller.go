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

// GetCourseByCourseName 通过课程名称获取课程信息
func GetCourseByCourseName(c *web.Context) {
	name, _ := c.QueryValue("name").String()
	service.CourseService().GetCourseByCourseName(c, name)
}

// GetCourseByCourseCode 通过课程代码获取课程信息
func GetCourseByCourseCode(c *web.Context) {
	code, _ := c.QueryValue("course_code").String()
	service.CourseService().GetCourseByCourseCode(c, code)
}

// DeleteCourseByCourseCode 通过课程代码删除课程
func DeleteCourseByCourseCode(c *web.Context) {
	code, _ := c.QueryValue("course_code").String()
	service.CourseService().DeleteCourseByCourseCode(c, code)
}
