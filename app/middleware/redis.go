package middleware

import (
	"context"
	"course-system/app/common"
	"course-system/app/models"
	"course-system/global"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

type redisOps struct {
	BookCourseRedisScript *redis.Script
	SetCourseCapScript    *redis.Script
}

var RedisOps = redisOps{
	BookCourseRedisScript: redis.NewScript(bookCourseLuaScript),
	SetCourseCapScript:    redis.NewScript(setCourseCapLuaScript)}
var ctx = context.Background()

/**
 */
const bookCourseLuaScript = `
if tonumber(redis.call('sismember', KEYS[1], ARGV[1]), 10) == 1 then
	return 2
else 
	if tonumber(redis.call('exists',KEYS[2]), 10) == 0 then
		return 3
	else 
		if tonumber(redis.call('get',KEYS[2]), 10) <= 0 then
				return 0
		else
			redis.call('decr', KEYS[2])
			redis.call('sadd', KEYS[1], ARGV[1])
			return 1
		end
	end
end
`

// 因为课程容量是定值，不存在更改的可能性，所有这样设置
// 如果通过访问数据库导致课程容量变化，如从10变到30，则下面一定出错
// 如果课程容量不存在，则直接设置课程容量
// 如果课程容量为-1，则表示之前访问时课程不存在，直接设置课程容量
// 课程容量存在，且不为-1，则表示已经被设置过了，不再设置
const setCourseCapLuaScript = `
if tonumber(redis.call('exists', KEYS[1]), 10) == 0 then
	redis.call('set', KEYS[1], ARGV[1])
else 
	if tonumber(redis.call('get', KEYS[1]), 10) == -1 then
		redis.call('set', KEYS[1], ARGV[1])
	end
end
`

func (redisOps *redisOps) BookCourse(keys []string, args ...interface{}) *redis.Cmd {
	return RedisOps.BookCourseRedisScript.Run(ctx, global.App.Redis, keys, args)
}

func (redisOps *redisOps) IsStuExist(stuID string) int {
	if global.App.Redis.Exists(ctx, "stu_list").Val() == 1 {
		if global.App.Redis.SIsMember(ctx, "stu_list", stuID).Val() {
			return 1
		} else {
			return -1
		}
	} else {
		return 0
	}
}

func (redisOps *redisOps) SetStuList(stuID []string) {
	global.App.Redis.SAdd(ctx, "stu_list", stuID)
}

func (redisOps *redisOps) IsCourseExist(courseID string) bool {
	result, err := global.App.Redis.Get(ctx, "course_cap_"+courseID).Int()
	if err == redis.Nil {
		return false
	} else if result < 0 {
		return false
	}
	return true
}

func (redisOps *redisOps) SetCourseCap(courseID string, cap int) {
	RedisOps.SetCourseCapScript.Run(ctx, global.App.Redis, []string{"course_cap_" + courseID}, cap)
}

func (redisOps redisOps) GetCourseCap(courseID string) (int, error) {
	result, err := global.App.Redis.Get(ctx, "course_cap_"+courseID).Int()
	return result, err
}

func (redisOps *redisOps) AddCourse(courseID, name, teacherID string, cap int) {
	redisOps.SetCourseCap(courseID, cap)
	redisOps.AddCourseInfo(courseID, name, teacherID)
}

func (redisOps *redisOps) AddCourseInfo(courseID, name, teacherID string) {
	global.App.Redis.HMSet(ctx, "course_info_"+courseID, "CourseID", courseID, "Name", name, "TeacherID", teacherID)
}

func (redisOps *redisOps) GetCourseInfo(courseId string) []interface{} {
	result, _ := global.App.Redis.HMGet(ctx, "course_info_"+courseId, "CourseID", "Name", "TeacherID").Result()
	return result
}

func (redisOps *redisOps) DelCourseInfo(courseID string) {
	global.App.Redis.Del(ctx, "course_info_"+courseID)
}

func (redisOps *redisOps) GetStuCourse(stuId string) []string {
	result, _ := global.App.Redis.SMembers(ctx, "stu_course_"+stuId).Result()
	return result
}

func (redisOps *redisOps) AddStuCourse(stuID string, courseID string) {
	global.App.Redis.SAdd(ctx, "stu_course_"+stuID, courseID)
}

//-------------------------------------------------------------------------

func (redisOps *redisOps) AddUserTypeInfo(UserType string, UserID string) {
	global.App.Redis.HMSet(ctx, "UserTypeInfo"+UserID, "UserID", UserID, "UserType", UserType)
}

func (redisOps *redisOps) AddMemberInfo(UserID string, Nickname string, Username string, UserType string) {
	global.App.Redis.HMSet(ctx, "MemberInfo"+UserID, "UserID", UserID, "Nickname", Nickname, "Username", Username, "UserType", UserType)
}

func (redisOps *redisOps) AddUserIDInfo(UserID string) {
	global.App.Redis.Set(ctx, "UserIDInfo"+UserID, 0, 0)
	//global.App.Redis.SAdd(ctx, "UserIDInfo"+UserID)
}

func (redisOps *redisOps) AddUsernameInfo(Username string) {
	global.App.Redis.Set(ctx, Username, 1, 0)
	//global.App.Redis.SAdd(ctx, Username)
}

func (redisOps *redisOps) AddUserDelInfo(UserID string) {
	global.App.Redis.Set(ctx, "UserDelInfo"+UserID, 0, 0)
	//global.App.Redis.SAdd(ctx, "UserDelInfo"+UserID)
}

func (redisOps *redisOps) DelUserIDInfo(UserID string) {
	global.App.Redis.Del(ctx, "UserIDInfo"+UserID)
}

func (redisOps *redisOps) DelUserTypeInfo(UserID string) {
	global.App.Redis.Del(ctx, "UserTypeInfo"+UserID)
}

func (redisOps *redisOps) GetUserTypeInfo(UserID string) string {
	result, _ := global.App.Redis.HGet(ctx, "UserTypeInfo"+UserID, "UserType").Result()
	return result
}

func (redisOps *redisOps) GetMemberInfo(UserID string) *models.User {
	user := models.User{}
	param1, _ := global.App.Redis.HGet(ctx, "MemberInfo"+UserID, "UserID").Result()
	param2, _ := global.App.Redis.HGet(ctx, "MemberInfo"+UserID, "Nickname").Result()
	param3, _ := global.App.Redis.HGet(ctx, "MemberInfo"+UserID, "Username").Result()
	param4, _ := global.App.Redis.HGet(ctx, "MemberInfo"+UserID, "UserType").Result()
	log.Println(param1, param2, param3, param4)
	log.Println(user)
	if param1 == "" {
		return &user
	}
	user.ID.ID, _ = strconv.ParseInt(param1, 10, 64)
	user.Nickname = param2
	user.Username = param3
	temp, _ := strconv.Atoi(param4)
	user.UserType = common.UserType(temp)
	return &user
}

func (redisOps *redisOps) IsExistUserID(UserID string) bool {
	_, err := global.App.Redis.Get(ctx, "UserIDInfo"+UserID).Int()
	log.Println("*", err)
	return err != redis.Nil
}

func (redisOps *redisOps) IsExistUsername(Username string) bool {
	_, err := global.App.Redis.Get(ctx, Username).Int()
	log.Println("////", err)
	return err != redis.Nil
}

func (redisOps *redisOps) IsExistDel(UserID string) bool {
	_, err := global.App.Redis.Get(ctx, "UserDelInfo"+UserID).Int()
	log.Println("#", err)
	return err != redis.Nil
}
