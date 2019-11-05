package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

//poll是指针
var pool *redis.Pool

//初始化执行函数
func init() {
	pool = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}
func main() {
	c := pool.Get()
	defer c.Close()
	s1 := time.Now().Unix()
	for i := 0; i < 100000; i++ {
		_, err := c.Do("lpush", "redislist", "jiang")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	s2 := time.Now().Unix()
	fmt.Println(s1, s2)
	pool.Close()
}
