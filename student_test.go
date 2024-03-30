package main

import (
	"testing"
	"tomatoPaper/api/controller"
	"tomatoPaper/pkg/database"
	"tomatoPaper/web"
)

func TestCreateStu(t *testing.T) {
	server := web.NewHTTPServer()
	server.Post("/api/student/register", controller.CreateStudent)
	server.Start(":8080")
}

func TestQueryStu(t *testing.T) {
	server := web.NewHTTPServer()
	server.Get("/api/student/query/student_by_name", controller.GetStudentByStudentName)
	server.Get("/api/student/query/student_by_student_id", controller.GetStudentByStudentId)
	server.Get("/api/student/query/course_by_name", controller.GetCourseByCourseName)
	server.Get("/api/student/query/course_by_code", controller.GetCourseByCourseCode)
	server.Start(":8080")
}

func TestDeleteStu(t *testing.T) {
	server := web.NewHTTPServer()
	server.Delete("/api/student/delete", controller.DeleteStudentByStudentId)
	server.Delete("/api/student/delete/course", controller.DeleteCourseByCourseCode)
	server.Start(":8080")
}

func TestUpdateStu(t *testing.T) {
	server := web.NewHTTPServer()
	//server.Post("/api/update/password", controller.UpdatePassword)
	server.Post("/api/update/course", controller.UpdateCourse)
	server.Start(":8080")
}

// 测试用户登录
func TestLoginStu(t *testing.T) {
	server := web.NewHTTPServer()
	server.Post("/api/student/login", controller.StudentLogin)
	server.Start(":8080")
}

// 测试论文的上传下载
func TestPaperStu(t *testing.T) {

}

func init() {
	database.SetupDBLink()
}
