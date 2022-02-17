package middleware

import (
	"context"
	"course-system/global"
	"github.com/go-redis/redis/v8"
)

type redisOps struct {
	BookCourseRedisScript *redis.Script
}

var RedisOps = redisOps{
	BookCourseRedisScript: redis.NewScript(bookCourseLuaScript)}
var ctx = context.Background()

var redisClient = global.App.Redis

/**
1. 判断数据是否已经选上了课程，如果抢上了课程，直接返回2
2. 如果数据没有选上课程，则查看课程容量，如果容量不足，则返回0
3. 如果容量足够，则容量--，并将其插入到课程中
*/
const bookCourseLuaScript = `
if tonumber(redis.call('sismember', KEYS[1], ARGV[1]), 10) == 1 then
	return 2
else 
	if tonumber(redis.call('exists',KEYS[2]), 10) == 0 then
		return 3
	else 
		if tonumber(redis.call('get',KEYS[2]), 10) == 0 then
				return 0
		else
			redis.call('decr', KEYS[2])
			redis.call('sadd', KEYS[1], ARGV[1])
			return 1
		end
	end
end
`

func (redisOps *redisOps) BookCourse(keys []string, args ...interface{}) *redis.Cmd {
	return RedisOps.BookCourseRedisScript.Run(ctx, redisClient, keys, args)
}

func (redisOps *redisOps) IsStuExist(stuID string) bool {
	return redisClient.SIsMember(ctx, "stu_list", stuID).Val()
}

func (redisOps *redisOps) IsCourseExist(courseID string) bool {
	_, err := redisClient.Get(ctx, "course_cap_"+courseID).Int()
	return err != redis.Nil
}

func (redisOps *redisOps) AddCourse(courseID, name, teacherID string, cap int) {
	redisClient.SetNX(ctx, "course_cap_"+courseID, cap, 0)
	redisOps.AddCourseInfo(courseID, name, teacherID)
}

func (redisOps *redisOps) GetStuCourse(stuId string) []string {
	result, _ := redisClient.SMembers(ctx, "stu_course_"+stuId).Result()
	return result
}

func (redisOps *redisOps) GetCourseInfo(courseId string) []interface{} {
	result, _ := redisClient.HMGet(ctx, "course_info_"+courseId, "CourseID", "Name", "TeacherID").Result()
	return result
}

func (redisOps *redisOps) DelCourseInfo(courseID string) {
	redisClient.Del(ctx, "course_info_"+courseID)
}

func (redisOps *redisOps) AddStuCourse(stuID string, courseID string) {
	redisClient.SAdd(ctx, "stu_course_"+stuID, courseID)
}

func (redisOps *redisOps) AddCourseInfo(courseID, name, teacherID string) {
	redisClient.HMSet(ctx, "course_info_"+courseID, "CourseID", courseID, "Name", name, "TeacherID", teacherID)
}
