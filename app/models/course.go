package models

type TCourse struct {
	ID
	Name      string `json:"name" gorm:"not null;unique;comment:课程名称"`
	Cap       int    `json:"cap" gorm:"not null;default:0;comment:课程容量"`
	TeacherID string `json:"teacher_id" gorm:"comment:教师"`
}
