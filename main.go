package main

import (
	"bufio"
	"fmt"
    "log"
	"os"
	"strings"
	"time"
)

func main() {
	//prompt and parse user input
	fmt.Printf("Freetime: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input: ", err)
	}
	input := strings.Fields(scanner.Text())
	qt := Parse(input)
    srv := setup()

	calendarlist, err := srv.CalendarList.List().Do()
    if err != nil {
       log.Fatalf("Unable to retrieve list of calendars: %v", err)
    }
	// println(len(calendarlist.Items))
	// for _, id := range calendarlist.Items {
	// 	fmt.Println(id.Id, "loc: ", id.TimeZone)
	// }
	calendarId := calendarlist.Items[0].Id
	//println(calendarId)
	loc, _ := time.LoadLocation(calendarlist.Items[0].TimeZone)

	// Do actual querying
	freetimeMap := QueryAll(qt, srv, calendarId, loc)
	//fmt.Println(len(freetimeMap))

	// Output and or print results
	printTimes(qt, freetimeMap)
}
