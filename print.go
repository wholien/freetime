package main

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
)

func printTimes(qt QueryTimes, freetimeMap map[Date][]calendar.TimePeriod) {
	for _, d := range qt.dates {
		tplist := freetimeMap[d]
		fmt.Printf("\n%s/%s: ", padNum(d.month), padNum(d.day))
		for index, tp := range tplist {
			if index > 0 {
				fmt.Printf(", ")
			}
			timestart, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", tp.Start)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("{%s:%s - ", padNum(timestart.Hour()), padNum(timestart.Minute()))
			timeend, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", tp.End)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s:%s}", padNum(timeend.Hour()), padNum(timeend.Minute()))
		}
	}
	fmt.Printf("\n\n")
}
