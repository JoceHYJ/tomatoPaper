package main

import (
	"testing"
	"tomatoPaper/api/controller"
	"tomatoPaper/pkg/database"
	"tomatoPaper/web"
)

func TestCreateTeacher(t *testing.T) {
	server := web.NewHTTPServer()
	server.Post("/api/teacher/register", controller.CreateTeacher)
	server.Start(":8080")
}

func TestQueryTeacher(t *testing.T) {
	server := web.NewHTTPServer()
	server.Get("/api/teacher/query/teacher_by_name", controller.GetTeacherByTeacherName)
	server.Get("/api/teacher/query/teacher_by_teacher_id", controller.GetTeacherByTeacherId)
	server.Get("/api/teacher/query/course_by_name", controller.GetCourseByCourseName)
	server.Get("/api/teacher/query/course_by_code", controller.GetCourseByCourseCode)
	server.Start(":8080")
}

func TestDeleteTeacher(t *testing.T) {
	server := web.NewHTTPServer()
	server.Delete("/api/teacher/delete", controller.DeleteTeacherByTeacherId)
	server.Delete("/api/teacher/delete/course", controller.DeleteCourseByCourseCode)
	server.Start(":8080")
}

func TestUpdateTeacher(t *testing.T) {
	server := web.NewHTTPServer()
	//server.Post("/api/update/password", controller.UpdatePassword)
	server.Post("/api/update/course", controller.UpdateCourse)
	server.Start(":8080")
}

// 测试用户登录
func TestLoginTeacher(t *testing.T) {
	server := web.NewHTTPServer()
	server.Post("/api/teacher/login", controller.TeacherLogin)
	server.Start(":8080")
}

// 测试论文的上传下载
func TestPaperTeacher(t *testing.T) {

}

func init() {
	database.SetupDBLink()
}
