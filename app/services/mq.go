package services

import (
	"course-system/app/dao"
	"course-system/app/models"
	"log"
)

const maxChannel = 65535

// BookChannel 管道，存储学生的选课数据
var BookChannel = make(chan *models.UserCourse, maxChannel)

// BookConsumer 管道的消费者，将学生的选课数据异步写入数据库
func BookConsumer() {
	for {
		//log.Println("等待数据")
		userCourse := <-BookChannel
		//log.Println("获取到数据", *userCourse)
		if err := dao.UserCourseDao.InsertUserCourseByAddress(userCourse); err != nil {
			log.Println("BookCourse插入数据库失败", err)
		}
	}
}
