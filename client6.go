package main

import (
	"fmt"

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
	_, err := c.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	fmt.Println(r)
	pool.Close()
}
