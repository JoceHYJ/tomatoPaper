package api

import (
	"fmt"
	"net/http"
	"tomatoPaper/api/entity"
	"tomatoPaper/common/util"
	"tomatoPaper/pkg/database"
	"tomatoPaper/web"
)

func CreateUser(c *web.Context) {
	defer func() {
		err := recover()
		if err != nil {
			util.HandleResponse(c, http.StatusBadRequest, "错误", err)
			return
		}
	}()
	var user entity.Users
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
