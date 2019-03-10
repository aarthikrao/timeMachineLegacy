package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// HTTP handler to schedule a job
func createJobHandler(rw http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	log.Println(string(body))
	var jobR Request
	err := json.Unmarshal(body, &jobR)
	if err != nil {
		log.Println("Error in parsing json")
	}
	//Validate job request
	log.Println("Scheduling Job:", jobR.JobName, " Time:", jobR.RunTime, "kjhdf:", 10)
	addJob(float64(jobR.RunTime), jobR.JobName)
}

// HTTP handler to delete a job
func deleteJobHandler(rw http.ResponseWriter, req *http.Request) {
	// TODO
}
