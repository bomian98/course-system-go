package services

import (
	"course-system/app/common"
	"course-system/app/dao"
	"course-system/app/models"
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

// 创建学生-课程服务对象，以防服务层函数过多，控制层调用函数时，函数重名的情况
type userSevice struct {
}

var UserService = new(userSevice)

// GetUserByUsername 根据用户名获得用户
func (userSevice *userSevice) GetUserByUsername(username string) (user *models.User, errno common.ErrNo) {
	var err error
	if user, err = dao.UserDao.GetUserByUsername(username); err != nil {
		errno = common.UserNotExisted
		return
	} else {
		errno = common.OK
		return
	}
}

// GetTMember 根据用ID获得用户的TMember
func (userSevice *userSevice) GetTMember(userID int64) (tMember common.TMember, errno common.ErrNo) {
	if user, err := dao.UserDao.GetUser(userID); err != nil {
		errno = common.UserNotExisted
		return
	} else {
		errno = common.OK
		tMember = common.TMember{
			UserID:   strconv.FormatInt(user.ID.ID, 10),
			Nickname: user.Nickname,
			Username: user.Username,
			UserType: common.UserType(user.UserType),
		}
		return
	}
}

func (userSevice *userSevice) UserMD5(userID string) string {
	d := []byte(userID)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
