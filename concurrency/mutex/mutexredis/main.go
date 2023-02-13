package main

import (
	"base/concurrency/mutex/mutexredis/reentrant_mutex"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	// 初始化连接池
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	max := 2
	for i := 0; i < max; i++ {
		go reentrant_mutex.NewLock(rdb).MockBusiness()
	}
	time.Sleep(time.Second * time.Duration(max/2))
}
