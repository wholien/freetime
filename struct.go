package main

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
	duration  int
}
