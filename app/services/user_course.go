package services

import (
	"course-system/app/common"
	"course-system/app/dao"
	"course-system/app/middleware"
	"course-system/app/models"
	"strconv"
)

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
// 判断是不是学生，判断课程是否存在 --- 插入数据时，如果是学生，则将其加入布隆过滤器中，插入课程同理
// 之后，若这个
func (userCourseService *userCourseService) BookCourse(stuId string, courseId string) common.ErrNo {
	// 两个key，一个是用户stu对应的课程列表，一个是课程cos对应的容量
	stu_id, _ := strconv.ParseInt(stuId, 10, 64)
	cos_id, _ := strconv.ParseInt(courseId, 10, 64)
	keys := []string{"stu_course_" + stuId, "course_cap_" + courseId}
	// 使用 Lua 脚本进行抢课，需要提前将数据存储到 redis 中
	value, _ := middleware.RedisOps.BookCourse(keys, courseId).Int()
	switch value {
	case 0:
		return common.CourseNotAvailable // 课程容量不足
	case 1:
		BookChannel <- &models.UserCourse{UserID: stu_id, CourseID: cos_id}
		return common.OK
	case 2:
		return common.CourseHasBooked // 课程已经绑定过
	default:
		return common.CourseNotExisted // 课程不存在
	}
}

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
