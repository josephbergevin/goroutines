package main

import "fmt"
import "time"

func main() {
	mySignal := make(chan struct{})
	for i := 0; i < 3; i++ {
		go func(iter int) {
			fmt.Printf("#%d waiting for signal\n", iter)
			<-mySignal
			fmt.Printf("#%d signal received\n", iter)
		}(i)
	}

	<-time.After(2 * time.Second)
	close(mySignal)
	<-time.After(1 * time.Second)
}
