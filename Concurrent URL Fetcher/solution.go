package main

import (
	"fmt"
	"net/http"
	"time"
)

type result struct {
	index  int
	url    string
	status string
}

func worker(id int, jobs <-chan int, urls []string, results chan<- result, client *http.Client) {
	for idx := range jobs {
		url := urls[idx]
		resp, err := client.Get(url)
		status := ""
		if err != nil {
			status = "error"
		} else {
			status = fmt.Sprintf("%d", resp.StatusCode)
			resp.Body.Close()
		}
		results <- result{index: idx, url: url, status: status}
	}
}

func main() {
	urls := []string{
		"https://google.com",
		"https://example.com",
		"https://golang.org",
		"https://nonexistent.xyz",
		"https://httpstat.us/500",
		"https://httpstat.us/404",
	}

	const workerCount = 5
	jobs := make(chan int, len(urls))
	resultsChan := make(chan result, len(urls))

	client := &http.Client{Timeout: 5 * time.Second}

	for w := 0; w < workerCount; w++ {
		go worker(w, jobs, urls, resultsChan, client)
	}

	for i := range urls {
		jobs <- i
	}
	close(jobs)

	results := make([]result, len(urls))
	for i := 0; i < len(urls); i++ {
		res := <-resultsChan
		results[res.index] = res
	}

	for _, r := range results {
		fmt.Printf("%s -> %s\n", r.url, r.status)
	}
}
