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
	go feedBack()
	go processor(1)
	// sendMQMessaage()
	test()
	http.HandleFunc("/", createJobHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))

}

// Test cases
func test() {
	for i := 0; i < 10; i++ {
		addJob(float64(currentTime()+int64(i*2)), fmt.Sprintf("%d", i))
	}
	for i := 10; i < 15; i++ {
		addJob(float64(currentTime()+10), fmt.Sprintf("%d", i))
	}
}

// Returns current time in seconds from epoch
func currentTime() int64 {
	return (time.Now().UnixNano() / (int64(time.Second) / int64(time.Nanosecond)))
}
