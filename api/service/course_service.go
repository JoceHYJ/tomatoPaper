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
	GetCourseByCourseName(c *web.Context, name string)
	GetCourseByCourseCode(c *web.Context, code string)
	DeleteCourseByCourseCode(c *web.Context, code string)
}

// CourseServiceImpl 实现接口
type CourseServiceImpl struct {
}

// GetCourseByCourseName  通过课程名称获取课程信息
func (cs CourseServiceImpl) GetCourseByCourseName(c *web.Context, name string) {
	course := dao.GetCourseByCourseName(name)
	c.RespJSON(200, course)
}

// GetCourseByCourseCode 通过课程代码获取课程信息
func (cs CourseServiceImpl) GetCourseByCourseCode(c *web.Context, code string) {
	course := dao.GetCourseByCourseCode(code)
	c.RespJSON(200, course)
}

// CreateCourse 创建课程
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

// DeleteCourseByCourseCode 通过课程代码删除课程
func (cs CourseServiceImpl) DeleteCourseByCourseCode(c *web.Context, code string) {
	defer func() {
		err := recover()
		if err != nil {
			c.RespJSON(400, err)
		}
	}()
	bool := dao.DeleteCourseByCourseCode(code)
	if !bool {
		c.RespJSON(400, "删除失败")
	}
	c.RespJSON(200, "删除成功")
}

var courseService = CourseServiceImpl{}

func CourseService() ICourseService {
	return &courseService
}
