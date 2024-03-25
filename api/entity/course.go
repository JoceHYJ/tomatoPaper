package entity

import "time"

// Courses 课程实体
type Courses struct {
	ID          int       `gorm:"type:int;autoIncrement;primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	SubTime     time.Time `gorm:"type:datetime;not null" json:"sub_time"`
	Description string    `gorm:"type:text" json:"description"`
}
