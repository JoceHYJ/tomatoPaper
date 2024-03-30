package api

import (
	"bytes"
	"tomatoPaper/api/entity"
	"tomatoPaper/common/util"
	"tomatoPaper/pkg/database"
	"tomatoPaper/pkg/jwt"
	"tomatoPaper/web"
)

func Login(c *web.Context) {
	var user entity.Users
	c.BindJson(&user)
	userId := user.UserID
	password := user.Password
	database.GormDB.Where("user_id=?", userId).First(&user)
	if user.UserID == "" {
		c.RespJSON(400, "用户不存在")
		return
	}
	// 验证密码
	hashedPassword := util.EncryptionMd5(password)
	if !bytes.Equal([]byte(hashedPassword), []byte(user.Password)) {
		c.RespJSON(400, "密码错误")
		return
	}
	// 生成 token
	token, err := jwt.GenerateTokenUser(user)
	if err != nil {
		c.RespJSON(500, "生成token失败")
		return
	}
	c.RespJSON(200, map[string]any{
		"msg":   "登录成功",
		"token": token,
	})
}
