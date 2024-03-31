package entity

// Courses 课程实体
type Courses struct {
	//ID          uint   `gorm:"type:int;autoIncrement;primaryKey" json:"id"`
	CourseCode  string `gorm:"type:varchar(255);primaryKey" json:"course_code"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	TeacherID   string `gorm:"type:varchar(255);not null" json:"teacher_id"`

	Teacher *Teachers `gorm:"foreignKey:TeacherID;references:TeacherID" json:"teacher"`
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
	TeacherID   string `json:"teacher_id" validate:"required"`
}

// UpdateCourseDto 更新课程参数
type UpdateCourseDto struct {
	CourseCode  string `json:"course_code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TeacherID   string `json:"teacher_id"`
}

//// CourseDetailDto 课程详情参数
//type CourseDetailDto struct {
//	Courses
//	Teacher Teachers
//}

// CourseInfoDto 课程信息 详情视图
type CourseInfoDto struct {
	//ID          uint   `json:"id"`
	CourseCode  string `json:"course_code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TeacherID   string `json:"teacher_id"`
}
