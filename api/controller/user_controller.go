package controller

import (
	"net/http"
	"tomatoPaper/api/entity"
	"tomatoPaper/api/service"
	"tomatoPaper/common/util"
	"tomatoPaper/web"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (c *UserController) CreateUser(ctx *web.Context) {
	var user entity.Users
	err := ctx.BindJson(&user)
	if err != nil {
		util.HandleResponse(ctx, http.StatusBadRequest, "参数错误", nil)
		return
	}
	// 调用服务创建用户
	err = c.UserService.CreateUser(&user)
	if err != nil {
		util.HandleResponse(ctx, http.StatusInternalServerError, "创建用户失败", nil)
		return
	}
	util.HandleResponse(ctx, http.StatusOK, "创建用户成功", nil)
}
