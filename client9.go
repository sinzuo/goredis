/*
ZADD 添加或者修改集合里面的成员

ZCARD获取指定键的成员个数

ZREVRANGE 取出集合，按照升序排列

ZRANGE 取出集合，按照降序排列

ZSCORE 取出集合中的某一个成员的分数

ZREM 移除这个集合里面的成员


*/
package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialPassword("123456"))
	if err != nil {
		panic(err)
	}
	_, err = conn.Do("ZADD", "score", 70, "mrtwenty")
	if err != nil {
		panic(err)
	}
	_, err = conn.Do("ZADD", "score", 80, "dazhaozhao", 85, "xiaoming")
	if err != nil {
		panic(err)
	}

	//获取成员个数
	result, err := conn.Do("ZCARD", "score")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	//取出 升序
	scoreMap, err := redis.StringMap(conn.Do("ZREVRANGE", "score", 0, 2, "withscores"))
	for name := range scoreMap {
		fmt.Println(name, scoreMap[name])
	}

	//取出 降序
	scoreMap, err = redis.StringMap(conn.Do("ZRANGE", "score", 0, 1, "withscores"))
	for name := range scoreMap {
		fmt.Println(name, scoreMap[name])
	}

	//取出 dazhaozhao的分数
	score, err := redis.Int(conn.Do("ZSCORE", "score", "dazhaozhao"))
	if err != nil {
		panic(err)
	}
	fmt.Println(score)

	//移除集合中的某一个或者多个成员
	result, err = conn.Do("ZREM", "score", "dazhaozhao")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
