package services

import (
	"course-system/app/common"
	"course-system/app/dao"
	"course-system/app/models"
)

// 创建学生-课程服务对象，以防服务层函数过多，控制层调用函数时，函数重名的情况
type userCourseSevice struct {
}

var UserCourseService = new(userCourseSevice)

// 这个应该属于成员/登录的服务
// 根据ID获得某个成员
func (userCourseSevice *userCourseSevice) IsStudentExisted(stu_id int64) (result bool) {
	if _, err := dao.UserDao.GetUser(stu_id); err != nil {
		return false
	} else {
		return true
	}
}

// 插入数据并返回是否出错
func (userCourseSevice *userCourseSevice) InsertUserCourse(usercourse models.UserCourse) (err error) {
	_, err = dao.UserCourseDao.InsertUserCourse(usercourse)
	return
}

func (userCourseSevice *userCourseSevice) GetUserCourses(stu_id int64) (courselist common.CourseListStruct) {
	// 搜一下，如果检索不到的话，是返回错误，还是返回空
	//courses, err := dao.UserCourseDao.GetUserCourseList(stu_id)
	//if err != nil {
	//	courselist.CourseList = nil
	//	return
	//}
	// 获取course表格中的老师和课程的信息
	//for course := range courses {
	//
	//}
}
