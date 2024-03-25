package dao

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"tomatoPaper/api/entity"
	"tomatoPaper/common/util"
	"tomatoPaper/web"
)

var db *gorm.DB

// 注册

// CreateStudent 创建学生
func CreateStudent(c *web.Context) {
	defer func() {
		err := recover()
		if err != nil {
			util.HandleResponse(c, http.StatusBadRequest, "error", err)
		}
	}()
	var stu entity.Users
	err := c.BindJson(&stu)
	if err != nil {
		util.HandleResponse(c, http.StatusBadRequest, "参数错误", err)
		return
	}
	tx := db.Create(&stu)
	if tx.RowsAffected > 0 {
		util.HandleResponse(c, http.StatusOK, "写入成功", "OK")
		return
	}
	fmt.Printf("写入失败, err:%v\n", err)
	util.HandleResponse(c, http.StatusBadRequest, "写入失败", tx)
	fmt.Println(tx)
}

// CreateTeacher 创建教师
func CreateTeacher(c *web.Context) {
	defer func() {
		err := recover()
		if err != nil {
			util.HandleResponse(c, http.StatusBadRequest, "error", err)
		}
	}()
	var teacher entity.Users
	err := c.BindJson(&teacher)
	if err != nil {
		util.HandleResponse(c, http.StatusBadRequest, "参数错误", err)
		return
	}
	tx := db.Create(&teacher)
	if tx.RowsAffected > 0 {
		util.HandleResponse(c, http.StatusOK, "写入成功", "OK")
		return
	}
	fmt.Printf("写入失败, err:%v\n", err)
	util.HandleResponse(c, http.StatusBadRequest, "写入失败", tx)
	fmt.Println(tx)
}

// CreateAdmin 创建管理员
func CreateAdmin(c *web.Context) {
	defer func() {
		err := recover()
		if err != nil {
			util.HandleResponse(c, http.StatusBadRequest, "error", err)
		}
	}()
	var admin entity.Users
	err := c.BindJson(&admin)
	if err != nil {
		util.HandleResponse(c, http.StatusBadRequest, "参数错误", err)
		return
	}
	tx := db.Create(&admin)
	if tx.RowsAffected > 0 {
		util.HandleResponse(c, http.StatusOK, "写入成功", "OK")
		return
	}
	fmt.Printf("写入失败, err:%v\n", err)
	util.HandleResponse(c, http.StatusBadRequest, "写入失败", tx)
	fmt.Println(tx)
}
