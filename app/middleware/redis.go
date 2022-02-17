package middleware

import (
	"context"
	"course-system/app/common"
	"course-system/app/dao"
	"course-system/global"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

var ctx = context.Background()

type redisOps struct {
	BookCourseRedisScript     *redis.Script
	GetUserCoursesRedisScript *redis.Script
}

var RedisOps = redisOps{
	BookCourseRedisScript: redis.NewScript(bookCourseLuaScript)}

/**
1. 判断数据是否已经选上了课程，如果抢上了课程，直接返回2
2. 如果数据没有选上课程，则查看课程容量，如果容量不足，则返回0
3. 如果容量足够，则容量--，并将其插入到课程中
*/
const bookCourseLuaScript = `
if tonumber(redis.call('sismember', KEYS[1], ARGV[1]), 10) == 1 then
	return 2
else 
	if tonumber(redis.call('exists',KEYS[s2]), 10) == 0 then
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

// CourseCapTestScript 创建1-20的课程，每个课程容量为100，测试使用
const CourseCapTestScript = `
for i=1,20 do
	redis.call('set', 'course_cap_'..i, 100) 
end
`

func (redisOps *redisOps) BookCourse(keys []string, args ...interface{}) *redis.Cmd {
	return RedisOps.BookCourseRedisScript.Run(context.Background(), global.App.Redis, keys, args)
}

/**
先判断数据是否在redis中
如果数据在redis中，则判断学生是否有效（可能为nil，为防止缓存穿透）
	如果学生有效，先通过学生id（stu_course_stuId）在hash中找到对应的课程id，再通过课程id（course_id_courseId）到hash中找到对应的课程组成课表
	如果学生无效，code = 11
如果数据不在redis中，则访问数据库，判断学生是否存在
	如果学生存在，返回对应课程id的list，然后将此数据set到redis的hash中
	如果学生不存在，code = 11且在redis的hash中插入该学生id，并置为nil
*/
func (redisOps *redisOps) GetUserCourses(keys []string, args ...interface{}) (code common.ErrNo, res common.CourseListStruct) {
	// keys[0]代表redis学生课表hash的key（即stu_course_stuId），key[1]代表学生id（即stuId）
	stuId, _ := strconv.ParseInt(keys[1], 10, 64)
	client := global.App.Redis
	stuCourses, err := client.HGetAll(ctx, keys[0]).Result()
	// 学生数据在redis中，区分是否是空值
	if err == nil {

		if stuCourses == nil { // 为防止缓存穿透的空值
			code = 11
		} else { // 有效学生id
			code = 1
			s := make([]common.TCourse, 0, len(stuCourses))
			// 得到课程map的key，代表该学生的所有课程id
			for _, v := range stuCourses {
				courseId := string("course_id_" + v)
				// 根据key从课程map中得到一条课程的信息
				course, _ := client.HGetAll(ctx, courseId).Result()
				cls := common.TCourse{course["CourseID"], course["Name"], course["TeacherID"]}
				s = append(s, cls)
			}
			res.CourseList = s
		}
	} else { // 学生数据不在redis中，到数据库取，然后放入redis中，区分学生是否存在，可以考虑存nil，因为不会变化，单独测试
		if usercourses, err := dao.UserCourseDao.GetUserCourseList(stuId); err != nil {
			log.Println("数据库获取UserCourse失败", err)
		} else {
			if usercourses == nil { // 学生不存在
				//防止缓存穿透：缓存空值，还可以考虑布隆过滤器
				err := client.HSet(ctx, keys[0], nil).Err()
				if err != nil {
					panic(err)
				}
				code = 11 // 学生不存在
			} else { // 学生存在
				s := make([]common.TCourse, 0, len(usercourses))
				data := make(map[string]interface{})
				//log.Println(os.Stdout, usercourses) // 打印结果
				for i, v := range usercourses {
					courseId := string("course_id_" + string(v.CourseID))
					// 根据key从课程map中得到一条课程的信息
					course, _ := client.HGetAll(ctx, courseId).Result()
					cls := common.TCourse{course["CourseID"], course["Name"], course["TeacherID"]}
					s = append(s, cls)

					//将一条课程的id放入data（map）中，之后用于存入redis
					data["course"+string(i)] = string(v.CourseID)
				}
				res.CourseList = s

				// 一次性保存多个hash字段值
				err := client.HMSet(ctx, keys[0], data).Err()
				if err != nil {
					panic(err)
				}
				code = 1
			}
		}
	}
	return code, res
}
