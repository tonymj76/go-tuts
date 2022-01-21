package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var URL = "https://storage.googleapis.com/server-success-rate/hosts/host%v/status"

// WebApplication is used to get the fields from the server
type WebApplication struct {
	RequestsCount int `json:"requests_count"`
	SuccessCount  int `json:"success_count"`
	Application   string
}

// Aggregate calculates the aggregate of success rate with requests count
func (webA *WebApplication) Aggregate() map[string]float64 {
	scoreData := make(map[string]float64)
	scoreData[webA.Application] = float64(webA.SuccessCount) / float64(webA.RequestsCount)
	return scoreData
}

// Walker since we know the amount of the request to be sent (1000 request)
// Walker takes in two channels and compute each request in goroutine
func Walker(streamNumber <-chan int, streamResult chan<- map[string]float64) {
	var webApp WebApplication
	//client := http.Client{
	//	Timeout: time.Duration(20 * time.Second),
	//}
	for number := range streamNumber {
		resp, err := http.Get(fmt.Sprintf(URL, number))
		if err != nil {
			streamResult <- map[string]float64{
				fmt.Sprintf("WebApp%v", number): 0,
			}
			continue
		}
		if err := json.NewDecoder(resp.Body).Decode(&webApp); err != nil {
			streamResult <- map[string]float64{
				fmt.Sprintf("WebApp%v", number): 0,
			}
			continue
		}
		streamResult <- webApp.Aggregate()
	}
}

func main() {
	var collectionData []map[string]float64
	numberOfRequest := 1001
	streamNumber := make(chan int, 100)
	streamResult := make(chan map[string]float64)

	for i := 0; i < cap(streamNumber); i++ {
		go Walker(streamNumber, streamResult)
	}

	go func() {
		for i := 1; i < numberOfRequest; i++ {
			streamNumber <- i
		}
	}()

	for i := 1; i < numberOfRequest; i++ {
		result := <-streamResult
		for _, value := range result {
			if value != 0 {
				collectionData = append(collectionData, result)
			}
		}
	}

	close(streamNumber)
	close(streamResult)
	fmt.Println(collectionData)
}
