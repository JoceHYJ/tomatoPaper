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

func init() {
	database.SetupDBLink()
}
