package test

import (
	"context"
	"course-system/app/common"
	"course-system/bootstrap"
	"course-system/global"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

// test前先执行
func TestMain(m *testing.M) {
	bootstrap.InitializeConfig()
	bootstrap.InitializeLog()
	global.App.DB = bootstrap.InitializeDB()
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()
	global.App.Redis = bootstrap.InitializeRedis()
	m.Run()
}

func TestGetRedisCap(t *testing.T) {
	ctx := context.Background()
	var courseID string
	var result []interface{}
	var err error
	courseID = "11"
	result, err = global.App.Redis.HMGet(ctx, "course_info_"+courseID, "CourseID", "Name", "TeacherID").Result()
	fmt.Println(result)
	fmt.Println(err)
	if err == redis.Nil {
		fmt.Println("11")
	}
	courseID = "12"
	result, err = global.App.Redis.HMGet(ctx, "course_info_"+courseID, "CourseID", "Name", "TeacherID").Result()
	fmt.Println(result)
	fmt.Println(err)
	if err == redis.Nil {
		fmt.Println("11")
	}
	//fmt.Println(result[1].(string))
	_, ok := result[2].(string)
	if ok {
		fmt.Println(11)
	}
	tt := make([]common.TCourse, 0)
	tt = append(tt, common.TCourse{CourseID: "1"})
	fmt.Println(tt)
	//fmt.Println(reflect.Type(err))
	//_, err := global.App.Redis.Get(context.Background(), "course_cap_1111").Int()
	//fmt.Println(err != redis.Nil)
	//var cap [31]int
	//for i := 1; i <= 30; i += 1 {
	//	key := "course_cap_" + strconv.Itoa(i)
	//	val, _ := global.App.Redis.Get(context.Background(), key).Result()
	//	cap[i], _ = strconv.Atoi(val)
	//}
	//fmt.Println(cap)
	////var capsum int
	//fmt.Println()
}

func TestGetDBCap(t *testing.T) {
	//var cap [31]int
	//for i := 1; i <= 30; i++ {
	//	key := strconv.Itoa(i)
	//	var usecourse []models.UserCourse
	//	result := global.App.DB.Where("course_id=?", key).Find(&usecourse)
	//	row := result.RowsAffected
	//	cap[i] = int(row)
	//}
	//fmt.Println(cap)
}
