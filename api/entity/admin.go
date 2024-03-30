package entity

// Admins 管理员
type Admins struct {
	AdminID   string `gorm:"type:varchar(255);primaryKey" json:"admin_id" validate:"required"`
	AdminName string `gorm:"type:varchar(255);not null" json:"admin_name" validate:"required"`
	Password  string `gorm:"type:varchar(255);not null" json:"password" validate:"required"`
}

// CreateAdminDto 创建管理员参数
type CreateAdminDto struct {
	AdminID   string `json:"admin_id" validate:"required"`
	AdminName string `json:"admin_name" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

// AdminLoginDto 管理员登录参数
type AdminLoginDto struct {
	AdminID string `json:"admin_id" validate:"required"`
	//AdminName string `json:"admin_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// JwtAdminDto JWT中存储的管理员信息
type JwtAdminDto struct {
	AdminID  string `json:"admin_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}
