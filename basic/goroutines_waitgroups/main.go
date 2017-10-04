package main

import (
	"fmt"
	"sync"
)

func main() {
	intsToPrint := []int{101, 102, 103}
	wg := &sync.WaitGroup{}
	for _, toPrint := range intsToPrint {
		wg.Add(1)
		go func(toPrint int) {
			fmt.Println(toPrint)
			wg.Done()
		}(toPrint)
	}
	wg.Wait()
}
