package main

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
)

func printTimes(dates []Date, freetimeMap map[Date][]calendar.TimePeriod) {
	for _, d := range dates {
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

func collatePrintTimes(timeperiods []calendar.TimePeriod) {
	fmt.Println("len(tp): ", len(timeperiods))
	for _, tp := range timeperiods {
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
	fmt.Printf("\n\n")
}

func collateTimes(dates []Date, freetimeMap map[Date][]calendar.TimePeriod, dur float64) []calendar.TimePeriod {
	if len(dates) == 1 {
		return freetimeMap[dates[0]]
	}
	tplist := freetimeMap[dates[0]]
	for _, d := range dates[1:] {
		tplistcurr := freetimeMap[d]

		fmt.Println(tplist)
		fmt.Println(tplistcurr)

		var result []calendar.TimePeriod
		i, j := 0, 0
		for i < len(tplist) && j < len(tplistcurr) {
			curr := calendar.TimePeriod{}
			i_start, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", tplist[i].Start)

			fmt.Println("i_start:", i_start)
			i_end, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", tplist[i].End)
			j_oldstart, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", tplistcurr[j].Start)
			fmt.Println("j_start:", j_oldstart)
			j_oldend, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", tplistcurr[j].End)

			j_start := time.Date(i_start.Year(), i_start.Month(), i_start.Day(), j_oldstart.Hour(), j_oldstart.Minute(), 0, 0, j_oldstart.Location())
			j_end := time.Date(i_start.Year(), i_start.Month(), i_start.Day(), j_oldend.Hour(), j_oldend.Minute(), 0, 0, j_oldend.Location())
			
			if i_start.After(j_end) {
				j++
			} else if j_start.After(i_end) {
				i++
			} else {
				if i_start.After(j_start) {
					fmt.Println("hey")
					curr.Start = tplist[i].Start
				} else {
					fmt.Println("hey2")
					curr.Start = tplistcurr[j].Start
				}
				if i_end.After(j_end) {
					curr.End = tplistcurr[j].End
					j++
				} else {
					curr.End = tplist[i].End
					i++
				}
				st, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", curr.Start)
				en, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", curr.End)
				newen := time.Date(st.Year(), st.Month(), st.Day(), en.Hour(), en.Minute(), 0, 0, en.Location())
				curr.Start = st.String()
				curr.End = newen.String()
				fmt.Println("curr:", curr)
				fmt.Println("hours:",st.Sub(newen).Hours())
				fmt.Println("duration:",dur)
				if st.Sub(newen).Hours() <= (-dur) {
					fmt.Println("append")
					result = append(result, curr)
				}
			}
		}
		tplist = result
	}

	return tplist
}
