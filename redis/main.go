package main

import (
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func initClient() (err error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "118.31.64.83:6379",
		Password: "",
		DB:       0,
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func initClientSentinel()(err error) {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName: "master",
		SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}



func main() {

}
