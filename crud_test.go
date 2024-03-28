package main

import (
	"testing"
	"tomatoPaper/api/controller"
	"tomatoPaper/pkg/database"
	"tomatoPaper/web"
)

func TestCreate(t *testing.T) {
	server := web.NewHTTPServer()
	server.Post("/api/register", controller.CreateUser)
	server.Post("/api/create_course", controller.CreateCourse)
	server.Start(":8080")
}

func TestQuery(t *testing.T) {
	server := web.NewHTTPServer()
	server.Get("/api/query/user", controller.GetUserByUsername)
	server.Get("/api/query/course_by_name", controller.GetCourseByCourseName)
	server.Get("/api/query/course_by_code", controller.GetCourseByCourseCode)
	server.Start(":8080")
}

func TestDelete(t *testing.T) {
	server := web.NewHTTPServer()
	server.Delete("/api/delete/user", controller.DeleteUserByUserId)
	server.Delete("/api/delete/course", controller.DeleteCourseByCourseCode)
	server.Start(":8080")
}

func init() {
	database.SetupDBLink()
}
