package services

import (
	"course-system/app/common"
	"course-system/app/dao"
)

// 创建学生-课程服务对象，以防服务层函数过多，控制层调用函数时，函数重名的情况
type userCourseSevice struct {
}

var UserCourseService = new(userCourseSevice)

// 这个应该属于成员/登录的服务
// 根据ID获得某个成员
func (userCourseSevice *userCourseSevice) IsStudentExisted(stu_id uint) (result bool) {
	if _, err := dao.UserDao.GetUser(stu_id); err != nil {
		return false
	} else {
		return true
	}
}

func (userCourseSevice *userCourseSevice) BookCourseService(request common.BookCourseRequest) {

}
