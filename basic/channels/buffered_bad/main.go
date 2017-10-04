package main

import "fmt"

func main() {
	intChan := make(chan int, 2)

	intChan <- 1
	intChan <- 2
	intChan <- 3 // blocks here

	fmt.Println(<-intChan, <-intChan)
}
