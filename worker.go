package main

import (
	"github.com/sirupsen/logrus"
)

// var wg sync.WaitGroup
var jobs chan int

func init() {
}

func StartWorkers(nbr int) {
	// wg = sync.WaitGroup{}
	jobs = make(chan int, nbr*10)

	for w := 0; w < nbr; w++ {
		go func() {
			worker()
		}()
	}
}

func StopWorkers() {
	// wg.Wait()
	close(jobs)
}

func worker() {
	// consume queue
	for taskNbr := range jobs {
		var eventType SegmentEventType
		var err error

		switch taskNbr % 18 {
		case 0, 6, 10:
			err = BuildRequestIdentify()
			eventType = SegmentEventTypeIdentify
		case 1:
			// err = BuildRequestGroup()
			eventType = SegmentEventTypeGroup
		case 2:
			err = BuildRequestAlias()
			eventType = SegmentEventTypeAlias
		case 3, 7, 11, 14, 16:
			// err = BuildRequestPage()
			eventType = SegmentEventTypePage
		case 4, 8, 12, 15, 17:
			// err = BuildRequestScreen()
			eventType = SegmentEventTypeScreen
		case 5, 9, 13:
			err = BuildRequestTrack()
			eventType = SegmentEventTypeTrack
		}

		if err != nil {
			logrus.Errorf("Event %s: %s", eventType, err.Error())
		} else {
			logrus.Infof("Event %s", eventType)
		}

		// wg.Done()
	}
}

func AddJob(i int) {
	// wg.Add(1)
	jobs <- i
}
