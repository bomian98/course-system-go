package services

import (
	"course-system/app/common"
	"course-system/app/dao"
	"course-system/app/middleware"
	"course-system/app/models"
	"course-system/global"
	"strconv"
)

// 创建学生-课程服务对象，以防服务层函数过多，控制层调用函数时，函数重名的情况
type userCourseService struct {
}

var UserCourseService = new(userCourseService)

// IsStudentExisted 这个应该属于成员/登录的服务
// 根据ID获得某个成员
func (userCourseService *userCourseService) IsStudentExisted(stuId int64) (result bool) {
	if _, err := dao.UserDao.GetUser(stuId); err != nil {
		return false
	} else {
		return true
	}
}

// BookCourse 根据stu和course进行抢课服务
func (userCourseService *userCourseService) BookCourse(stuId int64, courseId int64) common.ErrNo {
	// 两个key，一个是用户stu对应的课程列表，一个是课程cos对应的容量
	keys := []string{"stu_course_" + strconv.FormatInt(stuId, 10),
		"course_cap_" + strconv.FormatInt(courseId, 10)}
	// 使用 Lua 脚本进行抢课，需要提前将数据存储到 redis 中
	result, _ := middleware.RedisOps.BookCourseRedisScript.Run(global.App.Redis, keys, courseId).Int()
	// 处理结果
	if result == 0 {
		return common.CourseNotAvailable
	} else if result == 1 {
		go userCourseService.InsertUserCourse(models.UserCourse{UserID: stuId, CourseID: courseId})
		return common.OK
	} else {
		return common.CourseHasBooked
	}
}

// InsertUserCourse 插入数据并返回是否出错
func (userCourseService *userCourseService) InsertUserCourse(user_course models.UserCourse) (err error) {
	_, err = dao.UserCourseDao.InsertUserCourse(user_course)
	return
}

/**
获得学生的课程，从压测角度，从内存读取该学生的课程，然后再读取课程的数据库，获得课程信息更好
但是，从真实应用场景的话，应该是从数据库中读取学生的课程，然后再读取课程的数据库
*/

//func (userCourseSevice *userCourseSevice) GetUserCourses(stu_id int64) (courselist common.CourseListStruct) {
//	// 搜一下，如果检索不到的话，是返回错误，还是返回空
//	//courses, err := dao.UserCourseDao.GetUserCourseList(stu_id)
//	//if err != nil {
//	//	courselist.CourseList = nil
//	//	return
//	//}
//	// 获取course表格中的老师和课程的信息
//	//for course := range courses {
//	//
//	//}
//	return nil
//}
