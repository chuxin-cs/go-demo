package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// 声明一个全局的redisDb变量
var redisDb *redis.Client

func init() {
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.0:6379", // redis地址
		Password: "123456",         // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})

	_, err := redisDb.Ping().Result()
	if err != nil {
		fmt.Println("连接redis异常", err)
		return
	}
	fmt.Println("连接redis成功")
}

func main() {
	redisDb.Set("aa1", "aa", 0)
	time.Sleep(10 * time.Second)
	fmt.Println(redisDb.Get("aa1"))

}
