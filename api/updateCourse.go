package api

import (
	"tomatoPaper/api/entity"
	"tomatoPaper/pkg/database"
	"tomatoPaper/web"
)

func UpdateCourse(c *web.Context) {
	defer func() {
		err := recover()
		if err != nil {
			c.RespJSON(400, err)
		}
	}()
	var course entity.Courses
	err := c.BindJson(&course)
	if err != nil {
		c.RespJSON(400, err)
		return
	}
	var count int64
	database.GormDB.Model(&entity.Courses{}).Where("course_code=?", course.CourseCode).Count(&count)
	if count <= 0 {
		c.RespJSON(400, "课程不存在")
	}
	tx := database.GormDB.Model(&entity.Courses{}).Where("course_code=?", course.CourseCode).Updates(&course)
	if tx.RowsAffected > 0 {
		c.RespJSON(200, "更新成功")
		return
	}
	c.RespJSON(400, "更新失败")
	//fmt.Println(tx.Error)
}
