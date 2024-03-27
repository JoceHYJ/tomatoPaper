package main

import (
	"testing"
	"tomatoPaper/api/controller"
	"tomatoPaper/pkg/database"
	"tomatoPaper/web"
)

func TestCreateUser(t *testing.T) {
	server := web.NewHTTPServer()
	server.Post("/api/register", controller.CreateUser)
	server.Start(":8080")
}

func init() {
	database.SetupDBLink()
}
