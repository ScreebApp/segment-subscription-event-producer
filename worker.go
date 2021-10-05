package main

import (
	"github.com/sirupsen/logrus"
)

var jobs chan int

func StartWorkers(nbr int) {
	jobs = make(chan int, nbr*10)

	for w := 0; w < nbr; w++ {
		go func() {
			worker()
		}()
	}
}

func StopWorkers() {
	close(jobs)
}

func worker() {
	// consume queue
	for taskNbr := range jobs {
		var eventType SegmentEventType
		var err error

		switch taskNbr % 20 {
		case 0, 6, 10, 14:
			err = BuildRequestIdentify()
			eventType = SegmentEventTypeIdentify
		case 1:
			err = BuildRequestGroup()
			eventType = SegmentEventTypeGroup
		case 2:
			err = BuildRequestAlias()
			eventType = SegmentEventTypeAlias
		case 3, 7, 11, 15, 17, 19:
			err = BuildRequestPage()
			eventType = SegmentEventTypePage
		case 4, 8, 12, 16, 18:
			err = BuildRequestScreen()
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
	}
}

func AddJob(i int) {
	jobs <- i
}
