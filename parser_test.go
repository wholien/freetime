package main

import (
	"errors"
	"testing"
	"time"
)

var parseDurationTests = []struct {
	d   string
	exp float64
	e   error
}{
	{"1hour", float64(1), nil},
	{"1h2h3h4m", 0, errors.New("Duration is not in the correct format:\n-- If you want to enter both hours and minutes, it should be in the format '2h30m'")},
	{"30minutes", float64(.5), nil},
	{"2h30m", float64(2.5), nil},
	{"yixiaoshi", 0, errors.New("Duration is not in the correct format:\n-- Should be in format '1hour' or '30minutes'!")},
}

func TestParseDuration(t *testing.T) {
	for _, i := range parseDurationTests {
		actual, aerr := ParseDuration(i.d)
		if actual != i.exp {
			t.Errorf("ParseTimeRange(%s): expected %+v || actual %+v", i.d, i.exp, actual)
		} else if aerr != nil && i.e != nil && aerr.Error() != i.e.Error() {
			t.Errorf("ParseTimeRange(%s): expected err %s || actual err %+v", i.d, i.e, aerr)
		}
	}
}

var parseTimeRangeTests = []struct {
	tr  string
	exp TimeRange
	e   error
}{
	{"1pm-5pm", TimeRange{13, 17}, nil},
	{"1pm-5pm-8pm", TimeRange{}, errors.New("Time range is not in the correct format:\n-- Should be in format '1pm-5pm")},
	{"0pm-5pm", TimeRange{}, errors.New("Time range is not in the correct format:\n-- Should be in format '1pm-5pm")},
	{"5pm-1pm", TimeRange{}, errors.New("Starting time should be before the ending time")},
	{"9am-9pm", TimeRange{9, 21}, nil},
}

func TestParseTimeRange(t *testing.T) {
	for _, i := range parseTimeRangeTests {
		actual, aerr := ParseTimeRange(i.tr)
		if actual != i.exp {
			t.Errorf("ParseTimeRange(%s): expected %+v || actual %+v", i.tr, i.exp, actual)
		} else if aerr != nil && i.e != nil && aerr.Error() != i.e.Error() {
			t.Errorf("ParseTimeRange(%s): expected err %s || actual err %+v", i.tr, i.e, aerr)
		}
	}
}

var parseWeekdayTests = []struct {
	wd  string
	exp time.Weekday
	e   error
}{
	{"Mondays", time.Weekday(1), nil},
	{"Tuesdays", time.Weekday(2), nil},
	{"SundaYs", time.Weekday(0), nil},
	{"lolz", 0, errors.New("Not a valid weekday: Please enter a valid weekday")},
	{"day", 0, errors.New("Not a valid weekday: Please enter a valid weekday")},
}

func TestParseWeekday(t *testing.T) {
	for _, i := range parseWeekdayTests {
		actual, aerr := ParseWeekday(i.wd)
		if actual != i.exp {
			t.Errorf("ParseWeekday(%s): expected %v || actual %v", i.wd, i.exp, actual)
		} else if aerr != nil && i.e != nil && aerr.Error() != i.e.Error() {
			t.Errorf("ParseWeekday(%s): expected err %s || actual err %+v", i.wd, i.e, aerr)
		}
	}
}

var parseDatesTests = []struct {
	d   string
	exp []Date
	e   error
}{
	{"10/11", []Date{Date{month: 10, day: 11}}, errors.New("Incorrect format for dates")},
	{"12/25,12/26", []Date{Date{month: 12, day: 25}, Date{month: 12, day: 26}}, nil},
	{"1/1-1/3", []Date{Date{month: 1, day: 1}, Date{month: 1, day: 2}, Date{month: 1, day: 3}}, nil},
	{"2/2/3", []Date{}, errors.New("Incorrect format for date: '/' should separate the month and day")},
	{"10/11-10/12-10/20", nil, errors.New("When using a date range, it should be in the format '10/11-10/15'")},
}

func TestParseDates(t *testing.T) {
	for _, i := range parseDatesTests {
		actual, aerr := ParseDates(i.d)
		if len(actual) != len(i.exp) {
			t.Errorf("ParseDates(%s): expected %+v || actual %+v", i.d, i.exp, actual)
		}
		if len(actual) > 0 && len(actual) > 0 {
			if actual[0].month != i.exp[0].month || actual[0].day != i.exp[0].day {
				t.Errorf("ParseDates(%s): expected %+v || actual %+v", i.d, i.exp, actual)
			}
		} else if aerr != nil && i.e != nil && aerr.Error() != i.e.Error() {
			t.Errorf("ParseDates(%s): expected err %s || actual err %+v", i.d, i.e, aerr)
		}
	}
}

var parseDatesCommaTests = []struct {
	d   string
	exp []Date
	e   error
}{
	{"10/11", []Date{Date{month: 10, day: 11}}, nil},
	{"12/25,12/26", []Date{Date{month: 12, day: 25}, Date{month: 12, day: 26}}, nil},
	{"1/1,1/3", []Date{Date{month: 1, day: 1}, Date{month: 1, day: 3}}, nil}, //same as below
}

func TestParseDatesComma(t *testing.T) {
	for _, i := range parseDatesCommaTests {
		actual, aerr := ParseDatesComma(i.d)
		if len(actual) != len(i.exp) {
			t.Errorf("ParseDatesComma(%s): expected %+v || actual %+v", i.d, i.exp, actual)
		}
		if len(actual) > 0 && len(actual) > 0 {
			if actual[0].month != i.exp[0].month || actual[0].day != i.exp[0].day {
				t.Errorf("ParseDatesComma(%s): expected %+v || actual %+v", i.d, i.exp, actual)
			}
		} else if aerr != nil && i.e != nil && aerr.Error() != i.e.Error() {
			t.Errorf("ParseDatesComma(%s): expected err %s || actual err %+v", i.d, i.e, aerr)
		}
	}
}

var parseDatesDashTests = []struct {
	d   string
	exp []Date
	e   error
}{
	{"10/11", nil, errors.New("When using a date range, it should be in the format '10/11-10/15'")},
	{"12/25-12/26", []Date{Date{month: 12, day: 25}, Date{month: 12, day: 26}}, nil},
	{"1/1-1/3", []Date{Date{month: 1, day: 1}, Date{month: 1, day: 2}, Date{month: 1, day: 3}}, nil}, //same as below
}

func TestParseDatesDash(t *testing.T) {
	for _, i := range parseDatesDashTests {
		actual, aerr := ParseDatesDash(i.d)
		if len(actual) != len(i.exp) {
			t.Errorf("ParseDatesDash(%s): expected %+v || actual %+v", i.d, i.exp, actual)
		}
		if len(actual) > 0 {
			if actual[0].month != i.exp[0].month || actual[0].day != i.exp[0].day {
				t.Errorf("ParseDatesDash(%s): expected %+v || actual %+v", i.d, i.exp, actual)
			}
		} else if aerr != nil && i.e != nil && aerr.Error() != i.e.Error() {
			t.Errorf("ParseDatesDash(%s): expected err %s || actual err %+v", i.d, i.e, aerr)
		}
	}
}

var parseDateTests = []struct {
	d   string
	exp []Date
	e   error
}{
	{"10/11", []Date{Date{month: 10, day: 11}}, nil},
	{"12/25", []Date{Date{month: 12, day: 25}}, nil},
	{"1/1", []Date{Date{month: 1, day: 1}}, nil}, //same as below
	{"2/2/3", []Date{}, errors.New("Incorrect format for date: '/' should separate the month and day")},
	{"0/15", []Date{}, errors.New("Incorrect format for date: month and/or day out of range")},
}

func TestParseDate(t *testing.T) {
	for _, i := range parseDateTests {
		actual, aerr := ParseDate(i.d)
		if len(actual) != len(i.exp) {
			t.Errorf("ParseDate(%s): expected %+v || actual %+v", i.d, i.exp, actual)
		}
		if len(actual) > 0 {
			if actual[0].month != i.exp[0].month || actual[0].day != i.exp[0].day {
				t.Errorf("ParseDate(%s): expected %+v || actual %+v", i.d, i.exp, actual)
			}
		} else if aerr != nil && i.e != nil && aerr.Error() != i.e.Error() {
			t.Errorf("ParseDate(%s): expected err %s || actual err %+v", i.d, i.e, aerr)
		}
	}
}

var parseSingleDateTests = []struct {
	d   string
	exp Date
	e   error
}{
	{"10/11", Date{month: 10, day: 11}, nil},
	{"12/25", Date{month: 12, day: 25}, nil},
	{"1/1", Date{month: 1, day: 1}, nil}, //same as below
	{"2/2/3", Date{}, errors.New("Incorrect format for date: '/' should separate the month and day")},
	{"0/15", Date{}, errors.New("Incorrect format for date: month and/or day out of range")},
}

func TestParseSingleDate(t *testing.T) {
	for _, i := range parseSingleDateTests {
		actual, aerr := ParseSingleDate(i.d)
		if actual != i.exp {
			t.Errorf("ParseSingleDate(%s): expected %+v || actual %+v", i.d, i.exp, actual)
		} else if aerr != nil && i.e != nil && aerr.Error() != i.e.Error() {
			t.Errorf("ParseSingleDate(%s): expected err %s || actual err %+v", i.d, i.e, aerr)
		}
	}
}

var checkDurationTests = []struct {
	dur string
	exp bool
}{
	{"1hour", true},
	{"2hours", true},
	{"3hours", true}, //same as below
	{"30mins", true},
	{"15minutes", true},
	{"5epochs", false}, //checkTimeRange does not check order
	{"ruo", false},     //that is in ParseTimeRange()
	{"12", false},
}

func TestCheckDuration(t *testing.T) {
	for _, d := range checkDurationTests {
		actual := checkDuration(d.dur)
		if actual != d.exp {
			t.Errorf("CheckDuration(%s): expected %t || actual %t", d.dur, d.exp, actual)
		}
	}
}

var checkTimeRangeTests = []struct {
	s   string
	exp bool
}{
	{"1pm-5pm", true},
	{"1pm-0pm", false},
	{"2pm-1pm", true}, //same as below
	{"9am-5pm", true},
	{"4am-11am", true},
	{"4am-2am", true},  //checkTimeRange does not check order
	{"5pm-10pm", true}, //that is in ParseTimeRange()
	{"12pm-5pm", true},
}

func TestCheckTimeRange(t *testing.T) {
	for _, d := range checkTimeRangeTests {
		actual := checkTimeRange(d.s)
		if actual != d.exp {
			t.Errorf("CheckTimeRange(%s): expected %t || actual %t", d.s, d.exp, actual)
		}
	}
}

var checkDateTests = []struct {
	m   int
	d   int
	exp bool
}{
	{12, 1, true},
	{13, 1, false},
	{12, 0, false},
	{6, 31, false},
	{2, 28, true},
	{1, 15, true},
	{4, 29, true},
	{1, 1, true},
}

func TestCheckDate(t *testing.T) {
	for _, d := range checkDateTests {
		actual := checkDate(d.m, d.d)
		if actual != d.exp {
			t.Errorf("CheckDate(%s, %s): expected %t || actual %t", d.m, d.d, d.exp, actual)
		}
	}
}

var checkTimeTests = []struct {
	time     string //input
	expected bool   //expected result
}{
	{"1pm", true},
	{"11am", true},
	{"12am", true},
	{"13am", false},
	{"0pm", false},
	{"24vm", false},
	{"6pm", true},
	{"100mp", false},
}

func TestCheckTime(t *testing.T) {
	for _, d := range checkTimeTests {
		actual := checkTime(d.time)
		if actual != d.expected {
			t.Errorf("CheckTime(%s): expected %t || actual %t", d.time, d.expected, actual)
		}
	}
}
