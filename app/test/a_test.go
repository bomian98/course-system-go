package test

import (
	"context"
	"course-system/app/models"
	"course-system/bootstrap"
	"course-system/global"
	"fmt"
	"strconv"
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
	var cap [31]int
	for i := 1; i <= 30; i += 1 {
		key := "course_cap_" + strconv.Itoa(i)
		val, _ := global.App.Redis.Get(context.Background(), key).Result()
		cap[i], _ = strconv.Atoi(val)
	}
	fmt.Println(cap)
	//var capsum int
	fmt.Println()
}

func TestGetDBCap(t *testing.T) {
	var cap [31]int
	for i := 1; i <= 30; i++ {
		key := strconv.Itoa(i)
		var usecourse []models.UserCourse
		result := global.App.DB.Where("course_id=?", key).Find(&usecourse)
		row := result.RowsAffected
		cap[i] = int(row)
	}
	fmt.Println(cap)
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func BenchmarkFib10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}
