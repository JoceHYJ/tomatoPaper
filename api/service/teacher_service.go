package service

import (
	"tomatoPaper/api/dao"
	"tomatoPaper/api/entity"
	"tomatoPaper/common/util"
	"tomatoPaper/pkg/jwt"
	"tomatoPaper/web"

	"github.com/go-playground/validator/v10"
)

// ITeacherService 定义接口
type ITeacherService interface {
	TeacherLogin(c *web.Context, dto entity.TeacherLoginDto)
	CreateTeacher(c *web.Context, dto entity.CreateTeacherDto)
	GetTeacherByTeacherName(c *web.Context, teacherName string)
	GetTeacherByTeacherId(c *web.Context, teacherId string)
	DeleteTeacherByTeacherId(c *web.Context, teacherId string)
}

// TeacherServiceImpl 实现接口
type TeacherServiceImpl struct {
}

// TeacherLogin 登录
func (u TeacherServiceImpl) TeacherLogin(c *web.Context, dto entity.TeacherLoginDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		c.RespJSON(400, "参数校验失败")
		return
	}
	teacher := dao.TeacherDetail(dto)
	//teacher := dao.GetTeacherByTeacherID(dto.TeacherID)
	hashedPassword := util.EncryptionMd5(dto.Password)
	//if !bytes.Equal([]byte(hashedPassword), []byte(teacher.Password)) {
	//	c.RespJSON(400, "密码错误")
	//	return
	//}
	//fmt.Println("输入密码:", dto.Password)
	//fmt.Println("哈希化后的输入密码:", hashedPassword)
	//fmt.Println("数据库中的密码:", teacher.Password)
	if hashedPassword != teacher.Password {
		c.RespJSON(400, "密码错误")
		return
	}
	token, err := jwt.GenerateTokenTeacher(teacher)
	if err != nil {
		c.RespJSON(500, "生成token失败")
		return
	}
	c.RespJSON(200, map[string]any{
		"msg":   "登录成功",
		"token": token,
	})
}

type TeacherResponse struct {
	TeacherID   string `json:"teacher_id"`
	TeacherName string `json:"teacher_name"`
	//Role        string `json:"role"`
}

// GetTeacherByTeacherName 根据用户名获取用户信息
func (u TeacherServiceImpl) GetTeacherByTeacherName(c *web.Context, teacherName string) {
	defer func() {
		err := recover()
		if err != nil {
			c.RespJSON(400, err)
		}
	}()

	teacher := dao.GetTeacherByTeacherName(teacherName)
	//c.RespJSON(200, teacher)
	resp := TeacherResponse{
		TeacherID:   teacher.TeacherID,
		TeacherName: teacherName,
		//Role:        entity.RoleMap[teacher.Role],
	}
	c.RespJSON(200, resp)
	return
}

// GetTeacherByTeacherId 根据用户id获取用户信息
func (u TeacherServiceImpl) GetTeacherByTeacherId(c *web.Context, teacherId string) {
	defer func() {
		err := recover()
		if err != nil {
			c.RespJSON(400, err)
		}
	}()
	teacher := dao.GetTeacherByTeacherID(teacherId)
	//c.RespJSON(200, teacher)
	resp := TeacherResponse{
		TeacherID:   teacher.TeacherID,
		TeacherName: teacher.TeacherName,
		//Role:        entity.RoleMap[teacher.Role],
	}
	c.RespJSON(200, resp)
	return
}

// CreateTeacher 创建用户
func (u TeacherServiceImpl) CreateTeacher(c *web.Context, dto entity.CreateTeacherDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		c.RespJSON(400, err.Error())
		return
	}
	bool := dao.CreateTeacher(dto)
	if !bool {
		c.RespJSON(400, "创建失败")
	}
	c.RespJSON(200, "创建成功")
	return
}

// DeleteTeacherByTeacherId 根据id删除用户信息
func (u TeacherServiceImpl) DeleteTeacherByTeacherId(c *web.Context, teacherId string) {
	defer func() {
		err := recover()
		if err != nil {
			c.RespJSON(400, err)
		}
	}()
	bool := dao.DeleteTeacherByTeacherId(teacherId)
	if !bool {
		c.RespJSON(400, "删除失败")
	}
	c.RespJSON(200, "删除成功")
}

var teacherService = TeacherServiceImpl{}

func TeacherService() ITeacherService {
	return &teacherService
}
