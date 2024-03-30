package service

import (
	"tomatoPaper/api/dao"
	"tomatoPaper/api/entity"
	"tomatoPaper/common/util"
	"tomatoPaper/pkg/jwt"
	"tomatoPaper/web"

	"github.com/go-playground/validator/v10"
)

// IAdminService 定义接口
type IAdminService interface {
	AdminLogin(c *web.Context, dto entity.AdminLoginDto)
	CreateAdmin(c *web.Context, dto entity.CreateAdminDto)
	GetAdminByAdminName(c *web.Context, adminName string)
	GetAdminByAdminId(c *web.Context, adminId string)
	DeleteAdminByAdminId(c *web.Context, adminId string)
}

// AdminServiceImpl 实现接口
type AdminServiceImpl struct {
}

// AdminLogin 登录
func (u AdminServiceImpl) AdminLogin(c *web.Context, dto entity.AdminLoginDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		c.RespJSON(400, "参数校验失败")
		return
	}
	admin := dao.AdminDetail(dto)
	//admin := dao.GetAdminByAdminID(dto.AdminID)
	hashedPassword := util.EncryptionMd5(dto.Password)
	//if !bytes.Equal([]byte(hashedPassword), []byte(admin.Password)) {
	//	c.RespJSON(400, "密码错误")
	//	return
	//}
	//fmt.Println("输入密码:", dto.Password)
	//fmt.Println("哈希化后的输入密码:", hashedPassword)
	//fmt.Println("数据库中的密码:", admin.Password)
	if hashedPassword != admin.Password {
		c.RespJSON(400, "密码错误")
		return
	}
	token, err := jwt.GenerateTokenAdmin(admin)
	if err != nil {
		c.RespJSON(500, "生成token失败")
		return
	}
	c.RespJSON(200, map[string]any{
		"msg":   "登录成功",
		"token": token,
	})
}

type AdminResponse struct {
	AdminID   string `json:"admin_id"`
	AdminName string `json:"admin_name"`
	//Role        string `json:"role"`
}

// GetAdminByAdminName 根据用户名获取用户信息
func (u AdminServiceImpl) GetAdminByAdminName(c *web.Context, adminName string) {
	defer func() {
		err := recover()
		if err != nil {
			c.RespJSON(400, err)
		}
	}()

	admin := dao.GetAdminByAdminName(adminName)
	//c.RespJSON(200, admin)
	resp := AdminResponse{
		AdminID:   admin.AdminID,
		AdminName: adminName,
		//Role:        entity.RoleMap[admin.Role],
	}
	c.RespJSON(200, resp)
	return
}

// GetAdminByAdminId 根据用户id获取用户信息
func (u AdminServiceImpl) GetAdminByAdminId(c *web.Context, adminId string) {
	defer func() {
		err := recover()
		if err != nil {
			c.RespJSON(400, err)
		}
	}()
	admin := dao.GetAdminByAdminID(adminId)
	//c.RespJSON(200, admin)
	resp := AdminResponse{
		AdminID:   admin.AdminID,
		AdminName: admin.AdminName,
		//Role:        entity.RoleMap[admin.Role],
	}
	c.RespJSON(200, resp)
	return
}

// CreateAdmin 创建用户
func (u AdminServiceImpl) CreateAdmin(c *web.Context, dto entity.CreateAdminDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		c.RespJSON(400, err.Error())
		return
	}
	bool := dao.CreateAdmin(dto)
	if !bool {
		c.RespJSON(400, "创建失败")
	}
	c.RespJSON(200, "创建成功")
	return
}

// DeleteAdminByAdminId 根据id删除用户信息
func (u AdminServiceImpl) DeleteAdminByAdminId(c *web.Context, adminId string) {
	defer func() {
		err := recover()
		if err != nil {
			c.RespJSON(400, err)
		}
	}()
	bool := dao.DeleteAdminByAdminId(adminId)
	if !bool {
		c.RespJSON(400, "删除失败")
	}
	c.RespJSON(200, "删除成功")
}

var adminService = AdminServiceImpl{}

func AdminService() IAdminService {
	return &adminService
}
