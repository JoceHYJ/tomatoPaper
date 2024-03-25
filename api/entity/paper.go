package entity

// Papers 论文实体
type Papers struct {
	ID       int    `gorm:"type:int;autoIncrement;primaryKey" json:"id"`
	Title    string `gorm:"type:varchar(255);not null" json:"title"`
	AuthorID int    `gorm:"type:int;not null" json:"author_id"`
	CourseID int    `gorm:"type:int;not null" json:"course_id"`
	FilePath string `gorm:"type:varchar(255);not null" json:"file_path"`
}
