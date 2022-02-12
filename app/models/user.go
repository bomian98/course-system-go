package models

type User struct {
	ID
	Nickname string `json:"nickname" gorm:"not null;comment:用户昵称"`
	Username string `json:"username" gorm:"not null;comment:用户名称"`
	Password string `json:"password" gorm:"not null;default:'';comment:用户密码"`
	UserType UserType
	SoftDeletes
}

type UserType int

const (
	Admin   UserType = 1
	Student UserType = 2
	Teacher UserType = 3
)
