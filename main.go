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
	inputLen := len(input)

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

	loc, err := time.LoadLocation(calendarlist.Items[0].TimeZone)
	if err != nil {
		log.Fatalf("Unable to LoadLocation for calendarId's TimeZone: %v", err)
	}
	// Do actual querying
	dates, freetimeMap := QueryAll(qt, srv, calendarId, loc)

	if len(dates) == 0 {
		log.Fatalf("There are no valid time ranges given your constraints :(")
	}
	fmt.Println("HELLO")
	if inputLen == 3 {
		printTimes(dates, freetimeMap)
	} else if inputLen == 4 {
		collatePrintTimes(collateTimes(dates, freetimeMap, qt.duration))
	} else if inputLen == 5 {

	}
}
