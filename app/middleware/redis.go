package middleware

import "github.com/go-redis/redis"

type redisOps struct {
	BookCourseRedisScript *redis.Script
}

var RedisOps = redisOps{
	BookCourseRedisScript: redis.NewScript(bookCourseLuaScript)}

/**
1. 判断数据是否已经选上了课程，如果抢上了课程，直接返回2
2. 如果数据没有选上课程，则查看课程容量，如果容量不足，则返回0
3. 如果容量足够，则容量--，并将其插入到课程中
*/
const bookCourseLuaScript = `
if redis.call("sismember", KEYS[1], ARGS[1]) == 1 then
	return 2
else 
	if redis.call("get",KEYS[2]) == 0 then
		return 0
	else
		redis.call("decr", KEYS[2])
		redis.call("sadd", KEYS[1], ARGS[1])
		return 1
`
