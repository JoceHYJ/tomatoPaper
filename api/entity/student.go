package entity

// Students 学生
type Students struct {
	StudentID   string `gorm:"type:varchar(255);primaryKey" json:"student_id" validate:"required"`
	StudentName string `gorm:"type:varchar(255);not null" json:"student_name" validate:"required"`
	Password    string `gorm:"type:varchar(255);not null" json:"password" validate:"required"`
}

// CreateStudentDto 创建学生参数
type CreateStudentDto struct {
	StudentID   string `json:"student_id" validate:"required"`
	StudentName string `json:"student_name" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

// StudentLoginDto 学生登录参数
type StudentLoginDto struct {
	StudentID string `json:"student_id" validate:"required"`
	//StudentName string `json:"student_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// JwtStudentDto JWT中存储的学生信息
type JwtStudentDto struct {
	StudentID string `json:"student_id" validate:"required"`
	Password  string `json:"password" validate:"required"`
}
