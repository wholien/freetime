package main

import (
	"log"
	"strconv"
	"strings"
)

func daysInMonth(m int) int {
	if m < 1 || m > 12 {
		return 0
	}
	if m == 1 || m == 3 || m == 5 || m == 7 || m == 8 || m == 10 || m == 12 {
		return 31
	} else if m == 4 || m == 6 || m == 9 || m == 11 {
		return 30
	} else {
		return 28
	}
}

func isNum(t string, n int) bool {
	num, err := strconv.Atoi(strings.TrimFunc(t, trimToNum))
	if err != nil {
		log.Fatal(err)
	}
	return num >= 1 && num <= n
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

func padNum(num int) string {
	if num < 10 {
		n := strconv.Itoa(num)
		n = "0" + n
		return n
	}
	return strconv.Itoa(num)
}

func containsHour(dur string) bool {
	return strings.Contains(dur, "hour") || strings.Contains(dur, "hours") ||
		strings.Contains(dur, "hr") || strings.Contains(dur, "hrs") ||
		strings.Contains(dur, "h")
}

func containsMin(dur string) bool {
	return strings.Contains(dur, "min") || strings.Contains(dur, "mins") ||
		strings.Contains(dur, "minutes") || strings.Contains(dur, "minute") ||
		strings.Contains(dur, "m")
}
