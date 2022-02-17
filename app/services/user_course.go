package services

import (
	"course-system/app/common"
	"course-system/app/dao"
	"course-system/app/middleware"
	"course-system/app/models"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type userCourseService struct {
}

var UserCourseService = new(userCourseService)

// BookCourse 根据stu和course进行抢课服务
func (userCourseService *userCourseService) BookCourse(stuId string, courseId string) common.ErrNo {
	// 判断课程是否存在，通过course_cap来判断，
	// cap < 0，表示课程不存在；cap == 0，课程已满；cap > 0，课程存在
	cap, err := middleware.RedisOps.GetCourseCap(courseId)
	if cap < 0 { // 如果小于0，则说明课程不存在
		return common.CourseNotExisted
	} else if cap == 0 { // 课程容量不足，无法抢该门课程
		return common.CourseNotAvailable
	} else if err == redis.Nil { // 缓存中没有数据
		course, err := dao.CourseDao.GetCourse(courseId)
		if err != nil { // 没有这个课程
			middleware.RedisOps.SetCourseCap(courseId, -1)
			return common.CourseNotExisted
		} else { // 有这个课程，将这个课程加载到内存中
			middleware.RedisOps.AddCourse(courseId, course.Name, course.TeacherID, course.Cap)
		}
	}

	isExist := StudentExist(stuId)
	if !isExist {
		return common.StudentNotExisted
	}

	// 两个key，一个是用户stu对应的课程列表，一个是课程cos对应的容量
	keys := []string{"stu_course_" + stuId, "course_cap_" + courseId}
	// 使用 Lua 脚本进行抢课，需要提前将数据存储到 redis 中
	value, _ := middleware.RedisOps.BookCourse(keys, courseId).Int()
	switch value {
	case 0:
		return common.CourseNotAvailable // 课程容量不足
	case 1:
		stuId, _ := strconv.ParseInt(stuId, 10, 64)
		cosId, _ := strconv.ParseInt(courseId, 10, 64)
		BookChannel <- &models.UserCourse{UserID: stuId, CourseID: cosId}
		return common.OK
	case 2:
		return common.CourseHasBooked // 课程已经绑定过
	default:
		return common.CourseNotExisted // 课程不存在
	}
}

func (userCourseService *userCourseService) GetUserCourses(stuId string) ([]common.TCourse, common.ErrNo) {
	courseList := make([]common.TCourse, 0)
	// 学生是否存在
	if !StudentExist(stuId) {
		return courseList, common.StudentHasNoCourse
	}
	var courseIDList []string
	courseIDList = middleware.RedisOps.GetStuCourse(stuId)
	if len(courseIDList) == 0 { // 缓存中没有数据，从数据库访问
		stuIDInt, _ := strconv.Atoi(stuId)
		usercourses, err := dao.UserCourseDao.GetUserCourseList(stuIDInt)
		if err != nil { // 数据库出现故障
			return courseList, common.UnknownError
		}
		middleware.RedisOps.AddStuCourse(stuId, "") // 设置空值，避免多次访问数据库
		for _, usercourse := range usercourses {    // 遍历数据库中所有的数据
			courseIDInt := usercourse.CourseID
			courseID := strconv.Itoa(int(courseIDInt))
			courseList = append(courseList, GetCourseInfo(courseID)) // 获取课程信息
			middleware.RedisOps.AddStuCourse(stuId, courseID)        // 将该课程放到学生课程列表缓存中
		}
	} else { // courseIDList 并不为空
		// 第一个，即最后一个为""，则没有课程
		if len(courseList) == 1 && courseIDList[0] == "" {
			return courseList, common.StudentHasNoCourse
		}
		for _, courseID := range courseIDList { // 遍历所有的课程ID
			if courseID != "" {
				courseList = append(courseList, GetCourseInfo(courseID)) // 将课程信息添加到list中
			}
		}
	}
	if len(courseList) == 0 { // 最终获得的列表为空，学生没有课程
		return courseList, common.StudentHasNoCourse
	} else { // 返回学生课程
		return courseList, common.StudentHasCourse
	}
}

func GetCourseInfo(courseID string) common.TCourse {
	var tCourse common.TCourse
	var ok bool
	info := middleware.RedisOps.GetCourseInfo(courseID) //从缓存中读取课程信息
	tCourse.CourseID, ok = info[0].(string)             // 尝试将其转换为string类型
	if !ok {                                            //转换失败，即缓存中不存在该信息
		course, err := dao.CourseDao.GetCourse(courseID) // 从数据库中读取
		if err != nil {
			return common.TCourse{}
		}
		tCourse.CourseID = courseID
		tCourse.Name = course.Name
		tCourse.TeacherID = course.TeacherID
		middleware.RedisOps.AddCourseInfo(courseID, tCourse.Name, tCourse.TeacherID) // 写入缓存
	} else {
		tCourse.Name, _ = info[1].(string)
		tCourse.TeacherID, _ = info[2].(string)
	}
	return tCourse
}

func StudentExist(stuId string) bool {
	// 判断学生ID是否是学生，或学生是否存在
	// 判断stu_list是否存在，如果不存在，则读取数据库中所有数据
	// 从stu_course_stuID 读取，如果不存在，则读取数据库数据
	isStu := middleware.RedisOps.IsStuExist(stuId)
	isExist := false
	if isStu == -1 {
		return isExist // 学生不存在，对应ID不是学生或ID不存在
	} else if isStu == 0 {
		stus := dao.UserDao.GetAllStuList()
		var stuIDs []string
		for _, stu := range stus {
			stuID := strconv.Itoa(int(stu.ID.ID))
			stuIDs = append(stuIDs, stuID)
			if stuID == stuId { // 判断学生是否存在
				isExist = true
			}
		}
		middleware.RedisOps.SetStuList(stuIDs) // 添加到缓存中
	} else {
		isExist = true
	}
	return isExist
}
