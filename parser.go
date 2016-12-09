package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Parse(input []string) (QueryTimes, error) {
	//fmt.Println("Parse() ", input, len(input))
	len := len(input)
	switch {
	case len < 3 || len >= 6:
		return QueryTimes{}, errors.New("Parse: Incorrect number of args - there needs to be 3 to 5 arguments!")
	case len == 3:
		return Parse3(input)
	case len == 4:
		return Parse4(input)
	case len == 5:
		Parse5(input)
	}
	return QueryTimes{}, nil
}

func Parse3(input []string) (QueryTimes, error) {
	dur, err := ParseDuration(input[0])
	if err != nil {
		return QueryTimes{}, err
	}
	fmt.Println("duration:", dur)
	timerange, err := ParseTimeRange(input[1])
	if err != nil {
		return QueryTimes{}, err
	}
	fmt.Println("timerange:", timerange.start, timerange.end)
	dates, err := ParseDates(input[2])
	if err != nil {
		return QueryTimes{}, err
	}
	fmt.Println("datelength:", len(dates))
	for i, d := range dates {
		fmt.Println("day", i, ":", d.month, d.day)
	}
	return QueryTimes{timerange: timerange, dates: dates, duration: dur}, nil
}

func Parse4(input []string) (QueryTimes, error) {
	//1hour 1pm-5pm Mondays 11/10-12/15
	dur, err := ParseDuration(input[0])
	if err != nil {
		return QueryTimes{}, err
	}
	fmt.Println("duration:", dur)
	timerange, err := ParseTimeRange(input[1])
	if err != nil {
		return QueryTimes{}, err
	}
	fmt.Println("timerange:", timerange.start, timerange.end)
	weekday, err := ParseWeekday(input[2])
	if err != nil {
		return QueryTimes{}, err
	}
	fmt.Println("weekday:", weekday)
	dates, err := ParseDates(input[3])
	if err != nil {
		return QueryTimes{}, err
	}
	fmt.Println("datelength:", len(dates))
	for i, d := range dates {
		fmt.Println("day", i, ":", d.month, d.day)
	}
	return QueryTimes{timerange: timerange, dates: dates, duration: dur, weekday: weekday, careAboutWeekday: true, ordinal:0}, nil
}

func Parse5(input []string) (QueryTimes, error) {
	//1hour 1pm-5pm first Monday 11/10-12/15
	dur, err := ParseDuration(input[0])
	if err != nil {
		return QueryTimes{}, err
	}
	fmt.Println("duration:", dur)
	timerange, err := ParseTimeRange(input[1])
	if err != nil {
		return QueryTimes{}, err
	}
	fmt.Println("timerange:", timerange.start, timerange.end)
	ordinal, err := ParseOrdinal(input[2])
	if err != nil {
		return QueryTimes{}, err
	}
	fmt.Println("ordinal:", ordinal)
	weekday, err := ParseWeekday(input[3])
	if err != nil {
		return QueryTimes{}, err
	}
	fmt.Println("weekday:", weekday)
	dates, err := ParseDates(input[4])
	if err != nil {
		return QueryTimes{}, err
	}
	fmt.Println("datelength:", len(dates))
	for i, d := range dates {
		fmt.Println("day", i, ":", d.month, d.day)
	}
	return QueryTimes{timerange: timerange, dates: dates, duration: dur, weekday: weekday, careAboutWeekday: true, ordinal: ordinal}, nil
}

func ParseDuration(dur string) (float64, error) {
	//"1hour" || "30minutes" || 1h12m
	if !checkDuration(dur) {
		return 0, errors.New("Duration is not in the correct format:\n" +
			"-- Should be in format '1hour' or '30minutes'!")
	}
	if containsHour(dur) && containsMin(dur) {
		s := strings.SplitAfter(dur, "h")
		if len(s) != 2 || !isNum(s[0], 23) || !isNum(s[1], 59) || !strings.HasSuffix(s[0], "h") || !strings.HasSuffix(s[1], "m") {
			return 0, errors.New("Duration is not in the correct format:\n" +
				"-- If you want to enter both hours and minutes, it should be in the format '2h30m'")
		}
		h, err := strconv.Atoi(strings.TrimFunc(s[0], trimToNum))
		if err != nil {
			return 0, err
		}
		m, err := strconv.Atoi(strings.TrimFunc(s[1], trimToNum))
		if err != nil {
			return 0, err
		}
		return float64(h) + float64(m)/60, nil
	} else if containsHour(dur) {
		h, err := strconv.Atoi(strings.TrimFunc(dur, trimToNum))
		if err != nil {
			return 0, err
		}
		return float64(h), nil
	} else if containsMin(dur) {
		m, err := strconv.Atoi(strings.TrimFunc(dur, trimToNum))
		if err != nil {
			return 0, err
		}
		return float64(m) / 60, nil
	}
	return 0, nil
}

func ParseTimeRange(tr string) (TimeRange, error) {
	//"1pm-5pm"
	if !checkTimeRange(tr) {
		return TimeRange{}, errors.New("Time range is not in the correct format:\n-- Should be in format '1pm-5pm")
	}

	timerange := strings.Split(tr, "-")
	starttime, err := strconv.Atoi(strings.TrimFunc(timerange[0], trimToNum))
	if err != nil {
		return TimeRange{}, err
	}
	endtime, err := strconv.Atoi(strings.TrimFunc(timerange[1], trimToNum))
	if err != nil {
		return TimeRange{}, err
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
		return TimeRange{}, errors.New("Starting time should be before the ending time")
	}

	return TimeRange{start: starttime, end: endtime}, nil
}

func ParseWeekday(wd string) (time.Weekday, error) {
	if !strings.HasSuffix(wd, "s") {
		wd = wd + "s"
	}
	if !strings.Contains(wd, "day") {
		return 0, errors.New("Not a valid weekday: Please enter a valid weekday")
	} else {
		for i, d := range weekdays {
			if strings.ToLower(wd) == d {
				return time.Weekday(i), nil
			}
		}
	}
	return 0, errors.New("Not a valid weekday: Please enter a valid weekday")
}

func ParseOrdinal(ord string) (int, error) {
	for i, d := range ordinals {
		if strings.ToLower(ord) == d {
			return i, nil
		}
	}
	return 0, errors.New("Not a valid ordinal number: Ordinal numbers are words like 'first' and 'second', and only go up to 'fifth'")
}

func ParseDates(dates string) ([]Date, error) {
	//11/10 || 11/10,11/12,11/15 || 11/10-11/15
	if strings.Contains(dates, ",") {
		return ParseDatesComma(dates)
	} else if strings.Contains(dates, "-") {
		return ParseDatesDash(dates)
	} else {
		return ParseDate(dates)
	}
}

func ParseDatesComma(dates string) ([]Date, error) {
	days := strings.Split(dates, ",")
	var datelist []Date
	for _, d := range days {
		currd, err := ParseSingleDate(d)
		if err != nil {
			return nil, err
		}
		datelist = append(datelist, currd)
	}
	return datelist, nil
}

func ParseDatesDash(dates string) ([]Date, error) {
	daterange := strings.Split(dates, "-")
	if len(daterange) != 2 {
		return nil, errors.New("When using a date range, it should be in the format '10/11-10/15'")
	}
	startdate, err := ParseSingleDate(daterange[0])
	if err != nil {
		return nil, err
	}
	enddate, err := ParseSingleDate(daterange[1])
	if err != nil {
		return nil, err
	}
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
	return datelist, nil
}

func ParseDate(date string) ([]Date, error) {
	d, err := ParseSingleDate(date)
	if err != nil {
		return []Date{}, err
	}
	dates := []Date{d}
	return dates, nil
}

func ParseSingleDate(date string) (Date, error) {
	md := strings.Split(date, "/")
	if len(md) != 2 {
		return Date{}, errors.New("Incorrect format for date: '/' should separate the month and day")
	}
	m, err := strconv.Atoi(md[0])
	if err != nil {
		return Date{}, err
	}
	d, err := strconv.Atoi(md[1])
	if err != nil {
		return Date{}, err
	}
	if !checkDate(m, d) {
		return Date{}, errors.New("Incorrect format for date: month and/or day out of range")
	}
	return Date{month: m, day: d}, nil
}

func checkDuration(dur string) bool {
	//"1hour" || "2hours" || "30minutes" || "5mins" || "1h30m" || ...
	return strings.HasSuffix(dur, "hour") || strings.HasSuffix(dur, "hours") ||
		strings.HasSuffix(dur, "minutes") || strings.HasSuffix(dur, "minute") ||
		strings.HasSuffix(dur, "mins") || strings.HasSuffix(dur, "min") ||
		strings.HasSuffix(dur, "m") || strings.HasSuffix(dur, "h") ||
		strings.HasSuffix(dur, "hrs") || strings.HasSuffix(dur, "hr")
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
	return isNum(t, 12) && (strings.HasSuffix(t, "am") || strings.HasSuffix(t, "pm"))
}
