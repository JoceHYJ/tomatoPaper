package entity

// 用户相关结构体

// Users  用户实体
type Users struct {
	ID       int    `gorm:"type:int;autoIncrement;primaryKey" json:"id"`
	Username string `gorm:"type:varchar(255);not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"` // 1-> 学生 2-> 老师 3-> 管理员
}

// UserDTO 用户DTO
type UserDTO struct {
}
