package main

import "time"

type Date struct {
	month int
	day   int
}

type TimeRange struct {
	start int
	end   int
}

type QueryTimes struct {
	timerange TimeRange
	dates     []Date
	duration  float64
	weekday   time.Weekday
	careAboutWeekday bool
}

var weekdays = [...]string{
	"sundays",
	"mondays",
	"tuesdays",
	"wednesdays",
	"thursdays",
	"fridays",
	"saturdays",
}
