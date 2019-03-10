package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// This processes the scheduled jobs from redis set
func processor(id int) {
	fmt.Println("[INFO] Starting prcessor", id)
	for {
		job := fetchJob()
		if job.Member == nil {
			fmt.Println("Routine:", id, "Waiting for a job")
			continue
		}
		fmt.Println("Routine:", id, "Current time:", currentTime())
		fmt.Println("Time:", job.Score, " Job:", job.Member)
		if int64(job.Score)-currentTime() > maxTimeWindow {
			fmt.Println("Rescheduling job ...")
			reScheduleJob(job.Score, fmt.Sprintf("%v", job.Member))
			time.Sleep(2 * time.Second)
		} else {
			for {
				if int64(job.Score)-currentTime() <= 0 {
					break
				}
				time.Sleep(500 * time.Millisecond)
			}
			fmt.Println("Routine:", id, "Current time:", currentTime())
			runJob(job)
		}
	}
}

// This is a feedback process to check for missed scheduled jobs
func feedBack() {
	for {
		for _, jobWithScore := range findJobByTime(float64(currentTime() - 10)) {
			fmt.Println("[ERROR] Time :", jobWithScore.Score, " Job : ", jobWithScore.Member)
			reScheduleJob(jobWithScore.Score, fmt.Sprintf("%v", jobWithScore.Member))
		}
		fmt.Println("Feedback polling routing running")
		time.Sleep(5 * time.Second)
	}
}

// Add code related to processing your job here
func runJob(job redis.ZWithKey) {
	fmt.Println("Running job :: Time:", job.Score, " Job:", job.Member)
	//Remove from backup set
	removeJobFromRedis(backupRedisKey, fmt.Sprintf("%v", job.Member))
	// msgCh <- fmt.Sprint("Running job :: Time:", job.Score, " Job:", job.Member)
}

// Adds job to main set
func reScheduleJob(timestamp float64, jobName string) {
	addJobToRedis(scheduleRedisKey, timestamp, jobName)
}

// Add jobs to both sets
func addJob(timestamp float64, jobName string) {
	addJobToRedis(scheduleRedisKey, timestamp, jobName)
	addJobToRedis(backupRedisKey, timestamp, jobName)
}

// Remove jobs from both sets
func removeJob(jobName string) {
	removeJobFromRedis(scheduleRedisKey, jobName)
	removeJobFromRedis(backupRedisKey, jobName)
}
