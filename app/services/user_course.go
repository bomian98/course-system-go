package services

import (
	"System/app/common"
	"System/app/dao"
	"System/app/middleware"
	"System/app/models"
	"fmt"
	"strconv"
)

type userCourseService struct {
}

var UserCourseService = new(userCourseService)

// IsStudentExisted 这个应该属于成员/登录的服务
// 根据ID获得某个成员
func (userCourseService *userCourseService) IsStudentExisted(stuId int64) (result bool) {
	if _, err := dao.UserDao.GetUserByID(stuId); err != nil {
		return false
	} else {
		return true
	}
}

// BookCourse 根据stu和course进行抢课服务
// 判断是不是学生，判断课程是否存在 --- 插入数据时，如果是学生，则将其加入布隆过滤器中，插入课程同理
// 之后，若这个
func (userCourseService *userCourseService) BookCourse(stuId int64, courseId int64) common.ErrNo {
	// 两个key，一个是用户stu对应的课程列表，一个是课程cos对应的容量
	keys := []string{"stu_course_" + strconv.FormatInt(stuId, 10),
		"course_cap_" + strconv.FormatInt(courseId, 10)}
	// 使用 Lua 脚本进行抢课，需要提前将数据存储到 redis 中
	value, err := middleware.RedisOps.BookCourse(keys, courseId).Int()
	// 处理抢课结果
	fmt.Println(value, err)
	if value == 0 { // 结果为0，课程容量不足
		return common.CourseNotAvailable
	} else if value == 1 { // 有容量，学生也没有选择该门课，将数据放入管道中，异步写入数据库
		BookChannel <- &models.UserCourse{UserID: stuId, CourseID: courseId}
		return common.OK
	} else { // 学生重复选择同一门课程
		return common.CourseHasBound
	}
	//	StudentID
}
