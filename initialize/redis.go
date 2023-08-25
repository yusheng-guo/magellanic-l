package initialize

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/yushengguo557/magellanic-l/global"
	"log"
)

// InitRedis initialize Redis
func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "119.91.204.226:6379",
		Password: "8f4a4e91d3", // password set
		DB:       0,            // use default DB
		Protocol: 3,            // specify 2 for RESP 2 or 3 for RESP 3
	})
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		log.Fatalln("connect redis, err:", err)
	}

	log.Println("You successfully connected to Redis!")
	global.App.Redis = rdb
}
