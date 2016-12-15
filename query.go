package main

import (
	//"fmt"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
)

func QueryAll(qt QueryTimes, srv *calendar.Service, calendarId string, loc *time.Location) ([]Date, map[Date][]calendar.TimePeriod) {
	resp := make(map[Date][]calendar.TimePeriod)
	//var querydates []Date
	qt.dates = filter(qt, loc)
	for _, d := range qt.dates {
		tp := QueryOne(qt.timerange, d, qt.duration, srv, calendarId, loc)
		//querydates = append(querydates, d)
		resp[d] = tp
	}
	return qt.dates, resp
}

func QueryOne(timerange TimeRange, date Date, dur float64, srv *calendar.Service, calendarId string, loc *time.Location) []calendar.TimePeriod {
	fbri := &calendar.FreeBusyRequestItem{Id: calendarId}
	fbriarr := []*calendar.FreeBusyRequestItem{fbri}
	year := time.Now().Year()

	tstart := time.Date(year, time.Month(date.month), date.day, timerange.start, 0, 0, 0, loc)
	tend := time.Date(year, time.Month(date.month), date.day, timerange.end, 0, 0, 0, loc)

	freebusyResp, err := srv.Freebusy.Query(&calendar.FreeBusyRequest{Items: fbriarr, TimeMin: tstart.Format(time.RFC3339), TimeMax: tend.Format(time.RFC3339)}).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve freebusy. %v", err)
	}
	var freetimes []calendar.TimePeriod
	if len(freebusyResp.Calendars) > 0 {
		freest := tstart
		for _, v := range freebusyResp.Calendars {
			for _, tp := range v.Busy {
				st, err := time.Parse(time.RFC3339, tp.Start)
				if err != nil {
					log.Fatal(err)
				}
				if freest.Sub(st).Hours() <= (-dur) {
					freetimes = append(freetimes, calendar.TimePeriod{Start: freest.In(loc).String(), End: st.In(loc).String()})
				}
				en, err := time.Parse(time.RFC3339, tp.End)
				if err != nil {
					log.Fatal(err)
				}
				freest = en
			}
			if freest.Sub(tend).Hours() <= (-dur) {
				freetimes = append(freetimes, calendar.TimePeriod{Start: freest.In(loc).String(), End: tend.In(loc).String()})
			}
		}
	}
	//fmt.Println("freetimes size: ", len(freetimes))
	// for _, ft := range freetimes {
	// 	fmt.Println(ft.Start, "<--->", ft.End)
	// }
	return freetimes
}

func filter(qt QueryTimes, loc *time.Location) []Date {
	if qt.careAboutWeekday {
		var newdates []Date
		for _, d := range qt.dates {
			if time.Date(time.Now().Year(), time.Month(d.month), d.day, 0, 0, 0, 0, loc).Weekday() == qt.weekday {
				newdates = append(newdates, d)
			}
		}
		if qt.ordinal != 0 {
			var ordinalnewdates []Date
			for _, d := range newdates {
				thisday := time.Date(time.Now().Year(), time.Month(d.month), d.day, 0, 0, 0, 0, loc)
				counter := 0
				countday := time.Date(time.Now().Year(), time.Month(d.month), 1, 0, 0, 0, 0, loc)
				for !thisday.Before(countday) {
					if countday.Weekday() == qt.weekday {
						counter++
					}
					countday = countday.AddDate(0, 0, 1)
				}
				if counter == qt.ordinal {
					ordinalnewdates = append(ordinalnewdates, d)
				}
			}
			return ordinalnewdates
		}
		return newdates
	} else {
		return qt.dates
	}
}
