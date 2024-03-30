package entity

// Teachers 教师
type Teachers struct {
	TeacherID   string `gorm:"type:varchar(255);primaryKey" json:"teacher_id" validate:"required"`
	TeacherName string `gorm:"type:varchar(255);not null" json:"teacher_name" validate:"required"`
	Password    string `gorm:"type:varchar(255);not null" json:"password" validate:"required"`
}

// CreateTeacherDto 创建教师参数
type CreateTeacherDto struct {
	TeacherID   string `json:"teacher_id" validate:"required"`
	TeacherName string `json:"teacher_name" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

// TeacherLoginDto 教师登录参数
type TeacherLoginDto struct {
	TeacherID string `json:"teacher_id" validate:"required"`
	//TeacherName string `json:"teacher_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// JwtTeacherDto JWT中存储的教师信息
type JwtTeacherDto struct {
	TeacherID string `json:"teacher_id" validate:"required"`
	Password  string `json:"password" validate:"required"`
}
