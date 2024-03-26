package main

import (
	"testing"
	"tomatoPaper/api/controller"
	"tomatoPaper/pkg/database"
	"tomatoPaper/web"
)

func TestCreateUser(t *testing.T) {
	// 设置测试数据库连接
	err := database.SetupDBLink()
	if err != nil {
		t.Fatalf("failed to setup database link: %v", err)
	}

	server := web.NewHTTPServer()
	server.Post("/api/register", controller.CreateUser)
	server.Start(":8081")
}
