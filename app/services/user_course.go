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

func (userCourseSevice *userCourseService) GetUserCourses(stuId string) (code common.ErrNo, courselist common.CourseListStruct) {
	// 两个key，一个是redis学生课表hash的key（即stu_course_stuId），一个是所有课程hash的key（即stuId）
	keys := []string{"stu_course_" + stuId, stuId}
	// 使用 Lua 脚本进行抢课，需要提前将数据存储到 redis 中
	code, courselist = middleware.RedisOps.GetUserCourses(keys)

	switch code {
	case 11:
		return common.StudentNotExisted, courselist // 学生不存在
	case 1:
		return common.OK, courselist // 学生有课程
	default:
		return common.StudentHasCourse, courselist // 学生有课程
	}
}
