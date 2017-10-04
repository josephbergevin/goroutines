package main

import (
	"fmt"
	"time"
)

func main() {
	intervalsToSend := [][]int{
		[]int{1, 2, 3, 4, 5},
		[]int{1, 1, 1, 2},
		[]int{2, 2, 15},
	}

	intervalReportChan := make(chan string)
	quitinTimeChan := make(chan interface{})

	shiftDuration := 10 * time.Second
	overtimeDuration := 15 * time.Second

	for i := range intervalsToSend {
		go sender(i, intervalsToSend[i], intervalReportChan, quitinTimeChan)
	}

	fmt.Printf("waiting %s for sending...\n", shiftDuration.String())
	<-time.After(shiftDuration)

	// start receiving
	receiver(intervalReportChan, quitinTimeChan)

	fmt.Printf("waiting %s to close...\n", shiftDuration.String())
	<-time.After(overtimeDuration)

	// fmt.Println("signaling quitinTime")
	// close(quitinTimeChan)

	// fmt.Println("closing the report chan")
	// close(intervalReportChan)

}

func sender(senderNum int, intervals []int, intervalReportChan chan<- string, quitinTimeChan <-chan interface{}) {
	fmt.Printf("sender #%d reporting for duty\n", senderNum)
	itsQuitinTime := false
	var report string

	for _, interval := range intervals {
		select {
		case <-time.After(time.Duration(interval) * time.Second):
			report = fmt.Sprintf("sender #%d interval: %d", senderNum, interval)
			intervalReportChan <- report
			fmt.Println("report sent:", report)
		case <-quitinTimeChan:
			itsQuitinTime = true
		}
		if itsQuitinTime {
			break
		}
	}
	fmt.Printf("sender #%d returning\n", senderNum)
}

func receiver(intervalReportChan <-chan string, quitinTimeChan <-chan interface{}) {
	fmt.Println("receiver reporting for duty")
	for report := range intervalReportChan {
		fmt.Println("report received:", report)
	}
	fmt.Println("receiver returning")
}
