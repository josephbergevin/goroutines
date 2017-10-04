package main

import "fmt"

func main() {
	intsToPrint := []int{101, 102, 103}

	for _, toPrint := range intsToPrint {
		fmt.Println(toPrint)
	}
}
