package main

import (
	"log"
	"net/http"
	"sync"
)

// This example fetches several URLs concurrently,
// using a Throttler to block until all the fetches are complete.
// Compare to http://golang.org/pkg/sync/#example_WaitGroup
func main() {
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"https://nozzle.io/",
		"http://www.cnn.com/",
	}

	wg := sync.WaitGroup{}
	// create a buffered channel to send any errors through
	errChan := make(chan error, len(urls))

	for i := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(workerNum int, url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			_, err := http.Get(url)
			if err != nil {
				errChan <- err
			}
			log.Printf("worker #%d url:%s\n", workerNum, url)
		}(i, urls[i])
	}

	// Wait for all HTTP fetches to complete.
	wg.Wait()
	close(errChan)

	// check for errors on the errChan
	for err := range errChan {
		log.Fatal(err)
	}

}
