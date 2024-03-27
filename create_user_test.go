package main

import (
	"fmt"
	"net/http"
	"testing"
	"tomatoPaper/common/util"
	"tomatoPaper/pkg/database"
	"tomatoPaper/web"
)

func TestCreateUser(t *testing.T) {
	server := web.NewHTTPServer()
	server.Post("/api/register", createUser)
	server.Start(":8080")
}

type Users struct {
	ID       int    `gorm:"type:int;autoIncrement;primaryKey" json:"id"`
	Username string `gorm:"type:varchar(255);not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"` // 1-> 学生 2-> 老师 3-> 管理员
}

func createUser(c *web.Context) {
	defer func() {
		err := recover()
		if err != nil {
			util.HandleResponse(c, http.StatusBadRequest, "错误", err)
			return
		}
	}()
	var user Users
	err := c.BindJson(&user)
	if err != nil {
		util.HandleResponse(c, http.StatusBadRequest, "参数错误", err)
		return
	}
	database.GormDB.AutoMigrate(&user)
	tx := database.GormDB.Create(&user)
	if tx.RowsAffected > 0 {
		util.HandleResponse(c, http.StatusOK, "注册成功", "OK")
		return
	}
	fmt.Println("注册失败, err:", tx.Error)
	util.HandleResponse(c, http.StatusBadRequest, "注册失败", tx)
	fmt.Println(tx)
}

func init() {
	database.SetupDBLink()
}
