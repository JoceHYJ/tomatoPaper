package entity

// entity 层
// 用户 相关结构体

// Users  用户模型对象
type Users struct {
	//ID       uint   `gorm:"type:uint;autoIncrement;primaryKey" json:"id"`
	UserID   string `gorm:"type:varchar(255);primaryKey" json:"user_id"`
	Username string `gorm:"type:varchar(255);not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Role     uint   `gorm:"type:uint" json:"role"` // 1-> 学生 2-> 老师 3-> 管理员
}

var RoleMap = map[uint]string{
	1: "学生",
	2: "老师",
	3: "管理员",
}

//// UserIdDto 用户ID 传输对象
//type UserIdDto struct {
//	// ID uint `json:"id"`
//	UserID string `json:"user_id" validate:"required"`
//}

// CreateUserDto 新增用户参数
// 注册
type CreateUserDto struct {
	UserID   string `json:"user_id" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     uint   `json:"role" validate:"required"`
}

// UserLoginDto 登录对象
type UserLoginDto struct {
	UserID string `json:"user_id" validate:"required"`
	//Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// JwtUser 用户信息 用于jwt鉴权
type JwtUser struct {
	UserID string `json:"user_id"`
	//Username string `json:"username"`
	Password string `json:"password" validate:"required"`
	//Role     uint   `json:"role"`
}

// UserInfoDto 用户信息 详情视图
type UserInfoDto struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     uint   `json:"role"`
}
