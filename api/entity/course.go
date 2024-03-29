package entity

// Courses 课程实体
type Courses struct {
	//ID          uint   `gorm:"type:int;autoIncrement;primaryKey" json:"id"`
	CourseCode  string `gorm:"type:varchar(255);primaryKey" json:"course_code"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
}

//// CourseIdDto 课程ID参数
//type CourseIdDto struct {
//	ID uint `json:"id"`
//}

// CreateCourseDto 新增课程参数
type CreateCourseDto struct {
	CourseCode  string `json:"course_code" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// UpdateCourseDto 更新课程参数
type UpdateCourseDto struct {
	CourseCode  string `json:"course_code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CourseInfoDto 课程信息 详情视图
type CourseInfoDto struct {
	//ID          uint   `json:"id"`
	CourseCode  string `json:"course_code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
