package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	//"time"
)

func Parse(input []string) QueryTimes {
	//fmt.Println("Parse() ", input, len(input))
	len := len(input)
	switch {
	case len < 3 || len >= 6:
		log.Fatal(errors.New("Incorrect number of args - there needs to be 3 to 5 arguments!"))
	case len == 3:
		return Parse3(input)
	case len == 4:
		Parse4(input)
	case len == 5:
		Parse5(input)
	}
	return QueryTimes{}
}

func Parse3(input []string) QueryTimes {
	//fmt.Println("input:", input, len(input))
	dur := ParseDuration(input[0])
	fmt.Println("duration:", dur)
	timerange := ParseTimeRange(input[1])
	fmt.Println("timerange:", timerange.start, timerange.end)
	dates := ParseDates(input[2])
	fmt.Println("datelength:", len(dates))
	for i, d := range dates {
		fmt.Println("day", i, ":", d.month, d.day)
	}
	return QueryTimes{timerange: timerange, dates: dates, duration: dur}
}

func Parse4(input []string) {
	fmt.Println(input, len(input))
}

func Parse5(input []string) {
	fmt.Println(input, len(input))
}

func ParseDuration(dur string) int {
	//"1hour" || "30minutes"
	if !checkDuration(dur) {
		log.Fatal(errors.New("Duration is not in the correct format:\n" +
			"-- Should be in format '1hour' or '30minutes'!"))
	}
	if strings.Contains(dur, "hour") {
		h, err := strconv.Atoi(strings.TrimFunc(dur, trimToNum))
		if err != nil {
			log.Fatal(err)
		}
		return h
	} else if strings.Contains(dur, "min") {
		m, err := strconv.Atoi(strings.TrimFunc(dur, trimToNum))
		if err != nil {
			log.Fatal(err)
		}
		return m / 60
	}
	return 0
}

func ParseTimeRange(tr string) TimeRange {
	//"1pm-5pm"
	if !checkTimeRange(tr) {
		log.Fatal(errors.New("Time range is not in the correct format:\n" +
			"-- Should be in format '1pm-5pm"))
	}

	timerange := strings.Split(tr, "-")
	if !isNum(timerange[0]) || !isNum(timerange[1]) {
		log.Fatal(errors.New("Hours need to be 1-12!"))
	}
	starttime, err := strconv.Atoi(strings.TrimFunc(timerange[0], trimToNum))
	if err != nil {
		log.Fatal(err)
	}
	endtime, err := strconv.Atoi(strings.TrimFunc(timerange[1], trimToNum))
	if err != nil {
		log.Fatal(err)
	}

	startmeridian, endmeridian := "", ""
	if strings.HasSuffix(timerange[0], "am") {
		startmeridian = "am"
	} else if strings.HasSuffix(timerange[0], "pm") {
		startmeridian = "pm"
	}
	if strings.HasSuffix(timerange[1], "am") {
		endmeridian = "am"
	} else if strings.HasSuffix(timerange[1], "pm") {
		endmeridian = "pm"
	}
	starttime = meridianTo24(starttime, startmeridian)
	endtime = meridianTo24(endtime, endmeridian)
	if starttime >= endtime {
		log.Fatal(errors.New("Starting time should be before the ending time"))
	}

	return TimeRange{start: starttime, end: endtime}
}

func ParseDates(dates string) []Date {
	//11/10 || 11/10,11/12,11/15 || 11/10-11/15
	if strings.Contains(dates, ",") {
		return ParseDatesComma(dates)
	} else if strings.Contains(dates, "-") {
		return ParseDatesDash(dates)
	} else {
		return ParseDate(dates)
	}
}

func ParseDatesComma(dates string) []Date {
	days := strings.Split(dates, ",")
	var datelist []Date
	for _, d := range days {
		datelist = append(datelist, ParseSingleDate(d))
	}
	return datelist
}

func ParseDatesDash(dates string) []Date {
	fmt.Println("parse dates dash")
	daterange := strings.Split(dates, "-")
	if len(daterange) != 2 {
		log.Fatal(errors.New("Incorrect format for dates"))
	}
	startdate := ParseSingleDate(daterange[0])
	enddate := ParseSingleDate(daterange[1])
	var datelist []Date

	datelist = append(datelist, startdate)
	curr := startdate
	for curr.month != enddate.month || curr.day != enddate.day {
		if curr.day == daysInMonth(curr.month) {
			if curr.month == 12 {
				curr.month = 1
			} else {
				curr.month += 1
			}
			curr.day = 1
		} else {
			curr.day += 1
		}
		datelist = append(datelist, curr)
	}
	return datelist
}

func ParseDate(date string) []Date {
	dates := []Date{ParseSingleDate(date)}
	return dates
}

func ParseSingleDate(date string) Date {
	md := strings.Split(date, "/")
	if len(md) != 2 {
		log.Fatal(errors.New("Incorrect format for date!"))
	}
	m, err := strconv.Atoi(md[0])
	if err != nil {
		log.Fatal(err)
	}
	d, err := strconv.Atoi(md[1])
	if err != nil {
		log.Fatal(err)
	}
	if !checkDate(m, d) {
		log.Fatal(errors.New("Incorrect format for date!!"))
	}
	return Date{month: m, day: d}
}

func checkDuration(dur string) bool {
	//"1hour" || "2hours" || "30minutes" || "5mins" || ...
	return strings.HasSuffix(dur, "hour") ||
		strings.HasSuffix(dur, "hours") ||
		strings.HasSuffix(dur, "minutes") ||
		strings.HasSuffix(dur, "minute") ||
		strings.HasSuffix(dur, "mins") ||
		strings.HasSuffix(dur, "min")
}

func checkTimeRange(tr string) bool {
	//"1pm-5pm"
	if !strings.Contains(tr, "-") {
		return false
	}
	timerange := strings.Split(tr, "-")
	if len(timerange) > 2 {
		return false
	}
	if !checkTime(timerange[0]) || !checkTime(timerange[1]) {
		return false
	}

	return true
}

func checkDate(month, day int) bool {
	return !(month < 1 || month > 12) && !(day < 1 || day > daysInMonth(month))
}

func checkTime(t string) bool {
	//"1pm" || "11am"
	return isNum(t) && (strings.HasSuffix(t, "am") || strings.HasSuffix(t, "pm"))
}
