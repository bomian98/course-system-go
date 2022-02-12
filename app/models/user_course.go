package models

type UserCourse struct {
	ID
	UserID   uint `json:"user_id" gorm:"not null;comment:用户ID"`
	CourseID uint `json:"course_id" gorm:"not null;comment:课程ID"`
}
