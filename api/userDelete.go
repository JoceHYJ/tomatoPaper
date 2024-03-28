package api

import (
	"tomatoPaper/api/entity"
	"tomatoPaper/pkg/database"
	"tomatoPaper/web"
)

func DeleteUser(c *web.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.RespJSON(500, err)
		}
	}()
	id, _ := c.QueryValue("user_id").String()
	var count int64
	database.GormDB.Model(&entity.Users{}).Where("user_id = ?", id).Count(&count)
	if count <= 0 {
		c.RespJSON(404, "用户不存在")
		return
	}
	tx := database.GormDB.Where("user_id = ?", id).Delete(&entity.Users{})
	if tx.RowsAffected > 0 {
		c.RespJSON(200, "删除成功")
		return
	}
}
