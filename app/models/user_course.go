package models

type UserCourse struct {
	ID
	UserID   int64 `json:"user_id" gorm:"not null;comment:用户ID"`
	CourseID int64 `json:"course_id" gorm:"not null;comment:课程ID"`
}
