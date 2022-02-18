package services

import (
	"course-system/app/common"
	"course-system/app/dao"
	"course-system/app/middleware"
	"course-system/app/models"
	"crypto/md5"
	"encoding/hex"
	"log"
	"strconv"
)

// 创建学生-课程服务对象，以防服务层函数过多，控制层调用函数时，函数重名的情况
type userService struct {
}

var UserService = new(userService)

// GetUserByUsername 根据用户名获得用户
func (userSevice *userService) GetUserByUsername(username string) (user *models.User, errno common.ErrNo) {
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
func (userSevice *userService) GetTMember(userID int64) (tMember common.TMember, errno common.ErrNo) {
	if user, err := dao.UserDao.GetUserByID(userID); err != nil {
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

func (userSevice *userService) UserMD5(userID string) string {
	d := []byte(userID)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

//---------------------------------------------------------------------------------------------------------------

func CreateUseServices(request common.CreateMemberRequest) (common.ErrNo, string) {
	if middleware.RedisOps.IsExistUsername(request.Username + "s") {

		return common.UserHasExisted, ""
	}
	log.Println(request.Username)
	//log.Println(middleware.RedisOps.IsExistUsername(request.Username))
	user := models.User{Username: request.Username, Nickname: request.Nickname, Password: request.Password, UserType: request.UserType}
	if err := dao.UserDao.CreateUser(&user); err != nil {
		return common.UnknownError, ""
	}
	log.Println(user.ID)
	ID := strconv.FormatInt(user.ID.ID, 10)
	userType := strconv.Itoa(int(user.UserType))
	log.Println(ID, userType)
	middleware.RedisOps.AddUsernameInfo(request.Username + "s")
	middleware.RedisOps.AddUserIDInfo(ID)
	middleware.RedisOps.AddUserTypeInfo(userType, ID)
	return common.OK, ID
}

func UpdateServices(request common.UpdateMemberRequest) common.ErrNo {
	user := models.User{Nickname: request.Nickname}
	id, _ := strconv.ParseInt(request.UserID, 10, 64)
	if code := userStatus(request.UserID); code != common.OK {
		return code
	}
	if err := dao.UserDao.UpdateUser(user, id); err != nil {
		return common.UnknownError
	}
	return common.OK
}

func DeleteServices(request common.DeleteMemberRequest) common.ErrNo {
	id, _ := strconv.ParseInt(request.UserID, 10, 64)
	if code := userStatus(request.UserID); code != common.OK {
		return code
	}
	if err := dao.UserDao.DeleteUser(id); err != nil {
		return common.UnknownError
	}
	middleware.RedisOps.AddUserDelInfo(request.UserID)
	middleware.RedisOps.DelUserTypeInfo(request.UserID)
	middleware.RedisOps.DelUserIDInfo(request.UserID)
	return common.OK
}

func GetServices(request common.GetMemberRequest) (common.ErrNo, *models.User) {
	id, _ := strconv.ParseInt(request.UserID, 10, 64)
	if code := userStatus(request.UserID); code != common.OK {
		var temp *models.User
		return code, temp
	}
	temps := middleware.RedisOps.GetMemberInfo(request.UserID)
	if temps.Username != "" {
		return common.OK, temps
	}
	user, err := dao.UserDao.GetUserByID(id)
	if err != nil {
		return common.UnknownError, user
	}
	middleware.RedisOps.AddMemberInfo(request.UserID, user.Nickname, user.Username, strconv.Itoa(int(user.UserType)))
	return common.OK, user
}

func GetsServices(request common.GetMemberListRequest) (common.ErrNo, []*models.User) {
	var users []*models.User
	users, err := dao.UserDao.GetUsers(int(request.Offset), int(request.Limit))
	if err != nil {
		return common.UnknownError, users
	}
	return common.OK, users
}

func userStatus(ID string) common.ErrNo { //用户是否不存在,是否已删除
	if middleware.RedisOps.IsExistDel(ID) {
		return common.UserHasDeleted

	}
	if !middleware.RedisOps.IsExistUserID(ID) {
		return common.UserNotExisted
	}
	return common.OK
}

func GetUserType(ID string) (common.ErrNo, string) {
	userType := middleware.RedisOps.GetUserTypeInfo(ID)
	log.Println("usertype", userType)
	return common.OK, userType

}
