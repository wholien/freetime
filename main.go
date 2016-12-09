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
	fmt.Printf("Freetime> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input: ", err)
	}
	input := strings.Fields(scanner.Text())
	qt, err := Parse(input)
	if err != nil {
		log.Fatalf("Unable to parse user input: %v", err)
	}
	srv := setup()

	calendarlist, err := srv.CalendarList.List().Do()
	if err != nil {
		log.Fatalf("Unable to retrieve list of calendars: %v", err)
	}

	calendarId := calendarlist.Items[0].Id
	//println(calendarId)
	loc, err := time.LoadLocation(calendarlist.Items[0].TimeZone)
	if err != nil {
		log.Fatalf("Unable to LoadLocation for calendarId's TimeZone: %v", err)
	}
	// Do actual querying
	freetimeMap := QueryAll(qt, srv, calendarId, loc)
	//fmt.Println(len(freetimeMap))

	// Output and or print results
	printTimes(qt, freetimeMap)
}
