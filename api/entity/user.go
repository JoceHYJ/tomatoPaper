package entity

// entity 层
// 用户 相关结构体

// Users  用户模型对象
type Users struct {
	ID       uint   `gorm:"type:uint;autoIncrement;primaryKey" json:"id"`
	Username string `gorm:"type:varchar(255);not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Role     uint   `gorm:"type:uint" json:"role"` // 1-> 学生 2-> 老师 3-> 管理员
}

var RoleMap = map[uint]string{
	1: "学生",
	2: "老师",
	3: "管理员",
}

// UserIdDto 用户ID 传输对象
type UserIdDto struct {
	Id uint `json:"id"`
}

// CreateUserDto 新增用户参数
// 注册
type CreateUserDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     uint   `json:"role" validate:"required"`
}

// UserLoginDto 登录对象
type UserLoginDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	//IdKey    string
}

// UserInfoDto 用户信息 详情视图
type UserInfoDto struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Role     uint   `json:"role"`
}
