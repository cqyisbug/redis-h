package main

import (
	"testing"
	"github.com/go-redis/redis"
	"time"
)

func Test_TTL(t *testing.T){
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:[]string{"172.27.25.35:7000"},
	})
	val ,_ := client.TTL("foo").Result()

	println(GetExpireTime(val))

	println(Nano2Second(val))
}


func GetExpireTime(d time.Duration) string{
	now := time.Now()
	exp := now.Add(d)
	return exp.Format("2006-01-02 15:04:05")
}

func Nano2Second(d time.Duration) int64 {
	return int64(d / time.Second)
}