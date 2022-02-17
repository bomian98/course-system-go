package services

import (
	"course-system/app/common"
	"course-system/app/dao"
	"course-system/app/middleware"
	"course-system/app/models"
	"course-system/global"
	"log"
	"strconv"
)

type courseService struct {
}

var CourseService = new(courseService)

func (courseService *courseService) CreateCourse(name string, cap int) (courseID string, no common.ErrNo) {
	if cap < 0 {
		return "", common.UnknownError
	}
	ID, err := dao.CourseDao.IsCourseExistsByName(name) // 判断课程是否存在
	if err != nil {
		return "", common.UnknownError
	} else if ID > 0 {
		return strconv.Itoa(int(ID)), common.UnknownError
	}
	course := models.TCourse{Name: name, Cap: cap}
	if err = dao.CourseDao.InsertCourseByAddress(&course); err != nil { // 插入课程
		log.Println(err)
		return "", common.UnknownError // 插入时，数据库出现错误
	} else {
		courseID = strconv.Itoa(int(course.ID.ID))
		middleware.RedisOps.AddCourse(courseID, name, "", cap) // 插入缓存
		return courseID, common.OK
	}
}

func (courseService *courseService) GetCourse(courseID string) (common.TCourse, common.ErrNo) {
	if !middleware.RedisOps.IsCourseExist(courseID) {
		return common.TCourse{}, common.CourseNotExisted
	}
	tCourse := GetCourseInfo(courseID)
	return tCourse, common.OK
	//if tCourse, err := dao.CourseDao.GetCourse(courseID); err != nil {
	//	return tCourse, common.UnknownError
	//} else {
	//	return tCourse, common.OK
	//}
}

func (courseService *courseService) BindCourse(courseID, teacherID string) common.ErrNo {
	//if !middleware.RedisOps.IsCourseExist(courseID) {return common.CourseNotExisted}
	var tmp models.TCourse
	result := global.App.DB.Model(&models.TCourse{}).Where("id=?", courseID).First(&tmp)
	if result.Error != nil {
		return common.UnknownError // 数据不存在
	} else {
		if tmp.TeacherID == "" {
			tmp.TeacherID = teacherID
			global.App.DB.Save(&tmp)
			middleware.RedisOps.DelCourseInfo(courseID) // 先删除数据库，再删除缓存，降低问题可能性
			return common.OK                            // 绑定成功
		} else if tmp.TeacherID == teacherID {
			return common.CourseHasBound // 之前已经绑定过
		} else {
			return common.CourseHasBound // 之前已经绑定过，但是不是这个老师
		}
	}
}

func (courseService *courseService) UnBindCourse(courseID, teacherID string) common.ErrNo {
	id, _ := strconv.Atoi(courseID)
	tmp := models.TCourse{TeacherID: teacherID, ID: models.ID{ID: int64(id)}}
	result := global.App.DB.Where(&tmp).First(&tmp)
	if result.Error != nil {
		return common.CourseNotBind //当前课程和教师没有绑定在一起
	} else {
		tmp.TeacherID = ""
		global.App.DB.Save(tmp)
		middleware.RedisOps.DelCourseInfo(courseID) // 先删除数据库，再删除缓存，降低问题可能性
		return common.OK
	}
}

func (courseService *courseService) GetTeacherCourse(teacherID string) ([]*common.TCourse, common.ErrNo) {
	var tmp []models.TCourse
	var res []*common.TCourse
	result := global.App.DB.Where("teacher_id = ?", teacherID).Find(&tmp)
	if result.Error != nil {
		return res, common.UnknownError
	} else {
		for _, value := range tmp {
			var tcourse common.TCourse
			tcourse.Name = value.Name
			tcourse.CourseID = strconv.Itoa(int(value.ID.ID))
			tcourse.TeacherID = value.TeacherID
			res = append(res, &tcourse)
		}
		return res, common.OK
	}
}
