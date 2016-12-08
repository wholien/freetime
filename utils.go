package main

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

func daysInMonth(m int) int {
	if m < 1 || m > 12 {
		log.Fatal(errors.New("A month needs to be between 1 and 12 inclusive"))
	}
	if m == 1 || m == 3 || m == 5 || m == 7 || m == 8 || m == 10 || m == 12 {
		return 31
	} else if m == 4 || m == 6 || m == 9 || m == 11 {
		return 30
	} else {
		return 28
	}
}

func isNum(t string) bool {
	num, err := strconv.Atoi(strings.TrimFunc(t, trimToNum))
	if err != nil {
		log.Fatal(err)
	}
	return num >= 1 && num <= 12
}

func trimToNum(r rune) bool {
	if n := r - '0'; n >= 0 && n <= 9 {
		return false
	}
	return true
}

func meridianTo24(hour int, meridian string) int {
	if meridian == "am" {
		if hour == 12 {
			return 0
		} else {
			return hour
		}
	} else {
		if hour == 12 {
			return hour
		} else {
			return hour + 12
		}
	}
}
