package main

import (
	"testing"
	"tomatoPaper/api"
	"tomatoPaper/pkg/database"
	"tomatoPaper/web"
)

func TestCreateUser(t *testing.T) {
	server := web.NewHTTPServer()
	server.Post("/api/register", api.CreateUser)
	server.Start(":8080")
}

func init() {
	database.SetupDBLink()
}
