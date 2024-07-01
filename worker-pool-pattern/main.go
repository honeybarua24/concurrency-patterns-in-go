package main

import (
	"fmt"
	"log"
	"net/http"
)

type Site struct {
	URL string
}

type Result struct {
	URL, workerIdMsg string
	Status           int
}

func pingWebsite(wId int, jobs <-chan Site, results chan<- Result) {
	for site := range jobs {
		resp, err := http.Get(site.URL)
		if err != nil {
			log.Println(err.Error())
		}
		// sending into result channel
		results <- Result{
			workerIdMsg: fmt.Sprintf("\nThe worker id is %d, and status_code", wId),
			URL:         site.URL,
			Status:      resp.StatusCode,
		}
	}
}

func main() {
	jobs := make(chan Site, 3)
	results := make(chan Result, 3)

	// creating workers
	for w := 1; w <= 4; w++ {
		go pingWebsite(w, jobs, results)
	}

	urls := []string{
		"https://educative.io",
		"https://educative.io/learn",
		"https://educative.io/teach",
		"https://www.educative.io/explore/new",
		"https://www.educative.io/explore/picks",
		"https://www.educative.io/explore/early-access",
		"https://google.com",
	}

	// sending into jobs channel
	for _, url := range urls {
		jobs <- Site{URL: url}
	}
	close(jobs)

	for i := 1; i <= len(urls); i++ {
		result := <-results
		fmt.Println(result)
	}

}
