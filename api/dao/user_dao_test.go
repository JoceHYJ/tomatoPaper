package dao

import (
	"testing"
	"tomatoPaper/pkg/database"
	"tomatoPaper/web"
)

func TestCreateStudent(t *testing.T) {
	database.SetupDBLink()
	// 创建一个HTTP服务器
	server := web.NewHTTPServer()
	server.Post("/api/register/student", CreateStudent)
	server.Start(":8080")
}

func TestCreateTeacher(t *testing.T) {
	database.SetupDBLink()
	// 创建一个HTTP服务器
	server := web.NewHTTPServer()
	server.Post("/api/register/teacher", CreateTeacher)
	server.Start(":8080")
}

func TestCreateAdmin(t *testing.T) {
	database.SetupDBLink()
	// 创建一个HTTP服务器
	server := web.NewHTTPServer()
	server.Post("/api/register/admin", CreateAdmin)
	server.Start(":8080")
}
