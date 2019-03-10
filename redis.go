package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var scheduleRedisKey = "jobSet"
var backupRedisKey = "mySet"

var redisClient = redis.NewClient(&redis.Options{
	Addr:       "localhost:6379",
	PoolSize:   2,
	MaxRetries: 2,
	// Password:   "",
	DB: 0,
})

func removeJobFromRedis(setName string, jobName string) {
	err := redisClient.ZRem(setName, jobName)
	if err != nil {
		fmt.Println("Error in removing job from redis : " + jobName)
		// panic(err)
	}
}

func addJobToRedis(setName string, timestamp float64, jobName string) {
	err := redisClient.ZAdd(setName, redis.Z{
		Score:  float64(timestamp),
		Member: jobName,
	}).Err()
	if err != nil {
		fmt.Println("Error in adding job to redis : " + jobName)
	}
}

func findJobByTime(timestamp float64) []redis.Z {
	val, err := redisClient.ZRangeByScoreWithScores(backupRedisKey, redis.ZRangeBy{
		Min: string(0),
		Max: fmt.Sprintf("%f", timestamp),
	}).Result()
	if err != nil {
		fmt.Println("Error in fetching backup jobs from redis")
	}
	return val
}

func fetchJob() redis.ZWithKey {
	val, err := redisClient.BZPopMin(3*time.Second, scheduleRedisKey).Result()
	if err != nil {
		// fmt.Println("Error in getting job :", err.Error())
	}
	return val
}
