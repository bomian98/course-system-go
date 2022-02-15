package services

import (
	"course-system/app/dao"
	"course-system/app/models"
	"fmt"
	"os"
)

const maxChannel = 65535

// BookChannel 管道，存储学生的选课数据
// 要不要多整几个管道，毕竟数据库可以支持一次多个连接
var BookChannel = make(chan *models.UserCourse, maxChannel)

// BookConsumer 管道的消费者，将学生的选课数据异步写入数据库
func BookConsumer() {
	for {
		//log.Println("等待数据")
		userCourse := <-BookChannel
		//log.Println("获取到数据", *userCourse)
		// todo: 是否将失败数据压入管道中，应该不能压入同一个管道，大概率死锁
		if err := dao.UserCourseDao.InsertUserCourseByAddress(userCourse); err != nil {
			fmt.Println(os.Stdout, "BookCourse插入数据库失败", err)
		}
	}
}
