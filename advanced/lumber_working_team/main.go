package main

import (
	"context"
	"fmt"
	"time"
)

var backgroundContext = context.Background()

type Timber time.Duration

func main() {
	routeTimeout := 25 * time.Second
	c, cancelFunc := context.WithTimeout(backgroundContext, routeTimeout)
	defer cancelFunc()
	shiftManager(c)
}

func shiftManager(c context.Context) {
	timberRiverChan := make(chan Timber, 3) // for timber to be lumbered
	lumberTruckChan := make(chan string, 3) // for lumber to be delivered back to controller

	shiftDuration := 20 * time.Second
	cleanupDuration := 2 * time.Second
	ctrlCtx, _ := context.WithTimeout(c, shiftDuration+cleanupDuration)
	wrkrCtx, _ := context.WithTimeout(c, shiftDuration)

	go lumberController(ctrlCtx, timberRiverChan, lumberTruckChan)
	go lumberWorker(wrkrCtx, 1, timberRiverChan, lumberTruckChan)
	go lumberWorker(wrkrCtx, 2, timberRiverChan, lumberTruckChan)
	go lumberWorker(wrkrCtx, 3, timberRiverChan, lumberTruckChan)

	fmt.Printf("shiftManager >> waiting %s for shiftDuration...\n", shiftDuration.String())
	<-wrkrCtx.Done()

	fmt.Printf("shiftManager >> quitting time signalled to workers - waiting %s to shut down...\n", cleanupDuration.String())
	<-ctrlCtx.Done()

	fmt.Println("shiftManager >> shut down successfully")
}

func lumberController(c context.Context, timberRiverChan chan<- Timber, lumberTruckChan <-chan string) {
	fmt.Println("lumberController >> reporting for duty")
	timberLogs := pullTimber()
	reportsExpected := len(timberLogs)
	fmt.Printf("lumberController >> pulled %d timberLogs to dispatch\n", reportsExpected)

	// dispatch the timberLogs to the lumberWorkers
	for _, timber := range timberLogs {
		timberRiverChan <- timber
	}
	fmt.Println("lumberController >> sent the timberLogs to the timberRiverChan")

	close(timberRiverChan)
	fmt.Println("lumberController >> closed the timberRiverChan")

	reportsReceived := 0
	for {
		select {

		case report := <-lumberTruckChan:
			fmt.Println("lumberController >> received lumber shipment:", report)
			reportsReceived++
			if reportsReceived == reportsExpected {
				fmt.Printf("lumberController >> received %d of %d expected reports\n", reportsReceived, reportsExpected)
				fmt.Println("lumberController >> returning")
				return
			}

		case <-c.Done():
			fmt.Println("lumberController >> context cancelled")
			return

		}
	}
}

func lumberWorker(c context.Context, lumberWorkerNum int, timberRiverChan <-chan Timber, lumberTruckChan chan<- string) {
	fmt.Printf("lumberWorker #%d >> reporting for duty\n", lumberWorkerNum)
	errChan := make(chan error)

	for timber := range timberRiverChan {
		fmt.Printf("lumberWorker #%d >> received timber | job duration: %s\n", lumberWorkerNum, timber.JobDuration().String())

		go func() {
			errChan <- timber.ProduceLumber(c)
		}()

		err := <-errChan
		if err != nil {
			fmt.Printf("lumberWorker #%d >> context cancelled", lumberWorkerNum)
			lumberTruckChan <- timber.ReportError(lumberWorkerNum, err)
			break
		}
		lumberTruckChan <- timber.ShipLumber(lumberWorkerNum)
	}
	fmt.Printf("lumberWorker #%d >> returning\n", lumberWorkerNum)
}

func (timber Timber) ShipLumber(lumberWorkerNum int) string {
	return fmt.Sprintf("lumberWorker #%d | bundle size: %d", lumberWorkerNum, timber.BundleSize())
}

func (timber Timber) ReportError(lumberWorkerNum int, err error) string {
	return fmt.Sprintf("lumberWorker #%d | bundle size: %d | error: %s", lumberWorkerNum, timber.BundleSize(), err.Error())
}

func (timber Timber) ProduceLumber(c context.Context) error {
	select {
	case <-time.After(timber.JobDuration()):
		return nil

	case <-c.Done():
		fmt.Println("worker context cancelled")
		return c.Err()
	}
}

func (timber Timber) JobDuration() time.Duration {
	return time.Duration(timber)
}

func (timber Timber) BundleSize() int {
	return int(time.Duration(timber) / time.Second)
}

func pullTimber() []Timber {
	return []Timber{
		Timber(19 * time.Second),
		Timber(25 * time.Second),
		Timber(10 * time.Second),
		Timber(3 * time.Second),
		Timber(4 * time.Second),
	}
}
