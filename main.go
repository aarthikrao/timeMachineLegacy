package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var maxTimeWindow int64 = 10

// Remove the messages in mq routine fast enough
var msgCh = make(chan string, 100)

// Request structure
type Request struct {
	JobName        string
	RunTime        int64
	RepeatAfterSec int64
	DestIP         string
	DestPort       string
	DestExchange   string
}

func main() {
	defer redisClient.Close()
	// Spawn a routine for checking missed jobs
	go feedBack()
	// Spawn 5 Processor routines
	for i := 0; i < 5; i++ {
		go processor(i)
	}
	// Spawn 5 MQ Sender routines
	for i := 0; i < 5; i++ {
		go sendMQMessaage()
	}
	// test() // Uncomment to schedule jobs for testing
	http.HandleFunc("/", createJobHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))

}

// Test cases
func test() {
	// Set 1 lakh jobs at the same time
	for i := 0; i < 1000000; i++ {
		addJob(float64(currentTime()+5), fmt.Sprintf("%d", i))
	}
}

// Returns current time in seconds from epoch
func currentTime() int64 {
	return (time.Now().UnixNano() / (int64(time.Second) / int64(time.Nanosecond)))
}
