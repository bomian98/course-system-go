package models

import "course-system/app/common"

type User struct {
	ID
	Username string `json:"username" gorm:"not null;comment:用户名称"`
	Nickname string `json:"nickname" gorm:"not null;comment:用户昵称"`
	Password string `json:"password" gorm:"not null;default:'';comment:用户密码"`
	UserType common.UserType
	SoftDeletes
}
