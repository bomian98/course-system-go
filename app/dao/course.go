package dao

import (
	"course-system/app/common"
	"course-system/app/models"
	"course-system/global"
	"log"
	"strconv"
)

type courseDao struct {
}

var CourseDao = new(courseDao)

func (courseDao *courseDao) GetCourse(courseID string) (tCourse common.TCourse, err error) {
	var tmp models.TCourse

	if result := global.App.DB.First(&tmp, courseID); result.Error != nil {
		err = result.Error
	} else {
		tCourse.CourseID = strconv.Itoa(int(tmp.ID.ID))
		tCourse.TeacherID = tmp.TeacherID
		tCourse.Name = tmp.Name
	}
	return
}

func (courseDao *courseDao) IsCourseExistsByName(name string) (int64, error) {
	var tCourse models.TCourse
	result := global.App.DB.Where("name = ?", name).First(&tCourse)
	if result.RowsAffected == 0 {
		return -1, nil
	} else if result.Error != nil {
		return 0, result.Error
	} else {
		return tCourse.ID.ID, nil
	}
}

func (courseDao *courseDao) InsertCourseByAddress(course *models.TCourse) error {
	err := global.App.DB.Create(course).Error
	if err != nil { // 如果出错了，打印输出情况。调试时使用，之后注释掉
		log.Println(err)
	} else {
		//log.Println(os.Stdout, usercourse) // 打印结果
	}
	return err
}
