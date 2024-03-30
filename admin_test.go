package main

import (
	"testing"
	"tomatoPaper/api/controller"
	"tomatoPaper/pkg/database"
	"tomatoPaper/web"
)

func TestCreateAdmin(t *testing.T) {
	server := web.NewHTTPServer()
	server.Post("/api/admin/register", controller.CreateAdmin)
	server.Start(":8080")
}

func TestQueryAdmin(t *testing.T) {
	server := web.NewHTTPServer()
	server.Get("/api/admin/query/admin_by_name", controller.GetAdminByAdminName)
	server.Get("/api/admin/query/admin_by_admin_id", controller.GetAdminByAdminId)
	server.Get("/api/admin/query/course_by_name", controller.GetCourseByCourseName)
	server.Get("/api/admin/query/course_by_code", controller.GetCourseByCourseCode)
	server.Start(":8080")
}

func TestDeleteAdmin(t *testing.T) {
	server := web.NewHTTPServer()
	server.Delete("/api/admin/delete", controller.DeleteAdminByAdminId)
	server.Delete("/api/admin/delete/course", controller.DeleteCourseByCourseCode)
	server.Start(":8080")
}

func TestUpdateAdmin(t *testing.T) {
	server := web.NewHTTPServer()
	//server.Post("/api/update/password", controller.UpdatePassword)
	server.Post("/api/update/course", controller.UpdateCourse)
	server.Start(":8080")
}

// 测试用户登录
func TestLoginAdmin(t *testing.T) {
	server := web.NewHTTPServer()
	server.Post("/api/admin/login", controller.AdminLogin)
	server.Start(":8080")
}

// 测试论文的上传下载
func TestPaperAdmin(t *testing.T) {

}

func init() {
	database.SetupDBLink()
}
