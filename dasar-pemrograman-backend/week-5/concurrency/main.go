package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

var sites []string = []string{"https://facebook.com", "https://twitter.com", "https://linkedin.com", "https://google.com"}

func main() {
	start := time.Now()
	ch := make(chan string, len(sites))
	for _, site := range sites {
		wg.Add(1)
		go processSite(site, ch)
		wg.Add(1)
		go print(ch)
	}
	wg.Wait()
	end := time.Since(start)
	fmt.Printf("execution time: %.2f\n", end.Seconds())
}

func print(ch chan string) {
	fmt.Println(<-ch)
	wg.Done()
}

func processSite(site string, ch chan string) {
	resp, err := http.Get(site)
	if err != nil {
		ch <- fmt.Sprintf("error while calling %s - %s", site, err.Error())
		wg.Done()
		return
	}
	statusCode := resp.StatusCode
	status := "up"
	if statusCode >= 500 {
		status = "down"
	}
	ch <- fmt.Sprintf("[%d] %s is %s", statusCode, site, status)
	wg.Done()
}
