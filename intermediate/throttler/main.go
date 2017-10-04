package main

import (
	"log"
	"net/http"

	"github.com/nozzle/throttler"
)

// This example fetches several URLs concurrently,
// using a Throttler to block until all the fetches are complete.
func main() {
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"https://nozzle.io/",
		"http://www.cnn.com/",
	}

	// Create a new Throttler that will get 2 urls at a time
	th := throttler.New(2, len(urls))
	for i := range urls {
		// Launch a goroutine to fetch the URL.
		go func(workerNum int, url string) {
			// Fetch the URL.
			_, err := http.Get(url)
			log.Printf("worker #%d url:%s\n", workerNum, url)
			// Let Throttler know when the goroutine completes
			// so it can dispatch another worker
			th.Done(err)
		}(i, urls[i])
		// Pauses until a worker is available or all jobs have been completed
		// Returning the total number of goroutines that have errored
		// lets you choose to break out of the loop without starting any more
		if errCount := th.Throttle(); errCount > 0 {
			log.Fatal(th.Err())
		}
	}
}
