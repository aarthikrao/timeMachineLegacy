# Time Machine

A distributed, fault tolerant job scheduler that uses redis to set jobs. Jobs can be scheduled using a HTTP request

## Getting Started

* Install go-lang
* Install dependencies
```
go get -u github.com/go-redis/redis
go get -u github.com/streadway/amqp
```
* Build the project
```
go build
```
* Run the app
```
./timeMachine
```

## To create a job : 
```
POST / HTTP/1.1
Host: localhost:3000
Content-Type: application/json
{
  "jobName": "FirstJob",
  "runTime": 1552148330,
  "repeatAfterSec": null,
  "destIP": null,
  "destPort": null,
  "destExchange": null
}
```
