package main

import (
	"testing"
	"time"
)

var filterTests = []struct {
	qt  QueryTimes
	exp []Date
}{
	{QueryTimes{TimeRange{0,0}, []Date{Date{10,11}}, float64(1), time.Weekday(2), true, 2}, []Date{Date{10,11}}},
	{QueryTimes{TimeRange{0,0}, []Date{Date{12,25}}, float64(1), time.Weekday(0), true, 4}, []Date{Date{12,25}}},
	{QueryTimes{TimeRange{0,0}, []Date{Date{12,25}, Date{12,26}, Date{12,27}}, float64(1), time.Weekday(0), true, 4}, []Date{Date{12,25}}},
}

func TestFilter(t *testing.T) {
	for _, i := range filterTests {
		actual := filter(i.qt, time.UTC)
		if len(actual) != len(i.exp) {
			t.Errorf("ParseTimeRange(%s): expected %+v || actual %+v", i.qt, i.exp, actual)
		}
	}
}