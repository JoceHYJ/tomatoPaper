package service

import (
	"github.com/go-playground/validator/v10"
	"tomatoPaper/api/dao"
	"tomatoPaper/api/entity"
	"tomatoPaper/web"
)

// ICourseService 定义接口
type ICourseService interface {
	CreateCourse(c *web.Context, dao entity.CreateCourseDto)
}

// CourseServiceImpl 实现接口
type CourseServiceImpl struct {
}

func (cs CourseServiceImpl) CreateCourse(c *web.Context, dto entity.CreateCourseDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		c.RespJSON(400, err.Error())
		return
	}
	bool := dao.CreateCourse(dto)
	if !bool {
		c.RespJSON(400, "创建失败")
	}
	c.RespJSON(200, "创建成功")
	return
}

var courseService = CourseServiceImpl{}

func CourseService() ICourseService {
	return &courseService
}
