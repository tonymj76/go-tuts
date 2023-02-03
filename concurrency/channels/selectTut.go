package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	ApiKey            = "f32ee7b348msh230c75aaf106721p1366a6jsn952b266f7ae5"
	ApiGoogleNewsHost = "google-news.p.rapidapi.com"
	ApiFreeNewsHost   = "free-news.p.rapidapi.com"
	GoogleNewsUrl     = "https://google-news.p.rapidapi.com/v1/top_headlines?lang=en&country=US"
	FreeNewsUrl       = "https://free-news.p.rapidapi.com/v1/search?lang=en&q=Elon"
)

var (
	googleStream = make(chan News)
	freeStream   = make(chan News)
)

// Article models holds the structures
type Article struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	Id      string `json:"id"`
	MongoId string `json:"_id"`
}

type News struct {
	Source   string
	Articles []*Article `json:"articles"`
}

type Function struct {
	f       func(news chan<- News)
	channel chan News
}

type functions []*Function

func (f *Function) Run() {
	go f.f(f.channel)
}

func main() {
	fs := functions{
		{f: googleNews, channel: googleStream},
		{f: freeNews, channel: freeStream},
	}
	quickestApiResponse(fs)
}

func quickestApiResponse(fs functions) {
	var articles []*Article
	for _, function := range fs {
		function.Run()
	}

	select {
	case googleNewsResponse := <-googleStream:
		fmt.Printf("Source %s\n", googleNewsResponse.Source)
		articles = googleNewsResponse.Articles
	case freeNewsResponse := <-freeStream:
		fmt.Printf("source %s\n", freeNewsResponse.Source)
		articles = freeNewsResponse.Articles
	case <-time.After(time.Minute * 1):
		fmt.Println("time Out")
	}
	fmt.Printf("Article %v\n", articles)
}

func googleNews(news chan<- News) {
	req, err := http.NewRequest("GET", GoogleNewsUrl, nil)
	if err != nil {
		fmt.Printf("Error initializing request%v\n", err.Error())
		return
	}
	req.Header.Add("X-RapidAPI-Key", ApiKey)
	req.Header.Add("X-RapidAPI-Host", ApiGoogleNewsHost)
	client := &http.Client{Timeout: time.Minute * 2}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request %v\n", err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Google News Response StatusCode %v Status %v\n", resp.StatusCode, resp.Status)
			return
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		fmt.Printf("Google News Response StatusCode %v Status %v\n", resp.StatusCode, resp.Status)
		return
	}
	googleNewsArticles := News{Source: "GoogleNewsApi"}
	if err := json.NewDecoder(resp.Body).Decode(&googleNewsArticles); err != nil {
		fmt.Printf("Error decoding body %v\n", err.Error())
		return
	}
	fmt.Printf("Google Articles %v\n", googleNewsArticles)
	fmt.Printf("Google Articles Size %d\n", len(googleNewsArticles.Articles))
	news <- googleNewsArticles
}

func freeNews(news chan<- News) {
	req, err := http.NewRequest("GET", FreeNewsUrl, nil)
	if err != nil {
		fmt.Printf("Error initializing request%v\n", err.Error())
		return
	}
	req.Header.Add("X-RapidAPI-Key", ApiKey)
	req.Header.Add("X-RapidAPI-Host", ApiFreeNewsHost)
	client := &http.Client{Timeout: time.Minute * 2}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request %v\n", err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Google News Response StatusCode %v Status %v\n", resp.StatusCode, resp.Status)
			return
		}
	}(resp.Body)

	if resp.StatusCode == 200 {
		fmt.Printf("Free News Response StatusCode %v Status %v\n", resp.StatusCode, resp.Status)
		return
	}

	var freeNewsArticles News
	if err := json.NewDecoder(resp.Body).Decode(&freeNewsArticles); err != nil {
		fmt.Printf("Error decoding body %v\n", err.Error())
		return
	}

	freeNewsArticles.Source = "FreeNewsApi"
	fmt.Printf("Free Articles %v\n", freeNewsArticles)
	fmt.Printf("Free Articles Size %d\n", len(freeNewsArticles.Articles))
	news <- freeNewsArticles
}
