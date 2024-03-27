package entity

// Courses 课程实体
type Courses struct {
	ID          uint   `gorm:"type:int;autoIncrement;primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
}

// CourseIdDto 课程ID参数
type CourseIdDto struct {
	ID uint `json:"id"`
}

// CreateCourseDto 新增课程参数
type CreateCourseDto struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

// CourseInfoDto 课程信息 详情视图
type CourseInfoDto struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
