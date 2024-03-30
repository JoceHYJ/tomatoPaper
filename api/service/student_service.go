package service

import (
	"tomatoPaper/api/dao"
	"tomatoPaper/api/entity"
	"tomatoPaper/common/util"
	"tomatoPaper/pkg/jwt"
	"tomatoPaper/web"

	"github.com/go-playground/validator/v10"
)

// IStudentService 定义接口
type IStudentService interface {
	StudentLogin(c *web.Context, dto entity.StudentLoginDto)
	CreateStudent(c *web.Context, dto entity.CreateStudentDto)
	GetStudentByStudentName(c *web.Context, studentName string)
	GetStudentByStudentId(c *web.Context, studentId string)
	DeleteStudentByStudentId(c *web.Context, studentId string)
}

// StudentServiceImpl 实现接口
type StudentServiceImpl struct {
}

// StudentLogin 登录
func (u StudentServiceImpl) StudentLogin(c *web.Context, dto entity.StudentLoginDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		c.RespJSON(400, "参数校验失败")
		return
	}
	student := dao.StudentDetail(dto)
	//student := dao.GetStudentByStudentID(dto.StudentID)
	hashedPassword := util.EncryptionMd5(dto.Password)
	//if !bytes.Equal([]byte(hashedPassword), []byte(student.Password)) {
	//	c.RespJSON(400, "密码错误")
	//	return
	//}
	//fmt.Println("输入密码:", dto.Password)
	//fmt.Println("哈希化后的输入密码:", hashedPassword)
	//fmt.Println("数据库中的密码:", student.Password)
	if hashedPassword != student.Password {
		c.RespJSON(400, "密码错误")
		return
	}
	token, err := jwt.GenerateTokenStudent(student)
	if err != nil {
		c.RespJSON(500, "生成token失败")
		return
	}
	c.RespJSON(200, map[string]any{
		"msg":   "登录成功",
		"token": token,
	})
}

type StudentResponse struct {
	StudentID   string `json:"student_id"`
	StudentName string `json:"student_name"`
	//Role        string `json:"role"`
}

// GetStudentByStudentName 根据用户名获取用户信息
func (u StudentServiceImpl) GetStudentByStudentName(c *web.Context, studentName string) {
	defer func() {
		err := recover()
		if err != nil {
			c.RespJSON(400, err)
		}
	}()

	student := dao.GetStudentByStudentName(studentName)
	//c.RespJSON(200, student)
	resp := StudentResponse{
		StudentID:   student.StudentID,
		StudentName: studentName,
		//Role:        entity.RoleMap[student.Role],
	}
	c.RespJSON(200, resp)
	return
}

// GetStudentByStudentId 根据用户id获取用户信息
func (u StudentServiceImpl) GetStudentByStudentId(c *web.Context, studentId string) {
	defer func() {
		err := recover()
		if err != nil {
			c.RespJSON(400, err)
		}
	}()
	student := dao.GetStudentByStudentID(studentId)
	//c.RespJSON(200, student)
	resp := StudentResponse{
		StudentID:   student.StudentID,
		StudentName: student.StudentName,
		//Role:        entity.RoleMap[student.Role],
	}
	c.RespJSON(200, resp)
	return
}

// CreateStudent 创建用户
func (u StudentServiceImpl) CreateStudent(c *web.Context, dto entity.CreateStudentDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		c.RespJSON(400, err.Error())
		return
	}
	bool := dao.CreateStudent(dto)
	if !bool {
		c.RespJSON(400, "创建失败")
	}
	c.RespJSON(200, "创建成功")
	return
}

// DeleteStudentByStudentId 根据id删除用户信息
func (u StudentServiceImpl) DeleteStudentByStudentId(c *web.Context, studentId string) {
	defer func() {
		err := recover()
		if err != nil {
			c.RespJSON(400, err)
		}
	}()
	bool := dao.DeleteStudentByStudentId(studentId)
	if !bool {
		c.RespJSON(400, "删除失败")
	}
	c.RespJSON(200, "删除成功")
}

var studentService = StudentServiceImpl{}

func StudentService() IStudentService {
	return &studentService
}
