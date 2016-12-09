package main

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestDaysInMonth31(t *testing.T) {
	var thirty1 []int = []int{1, 3, 5, 7, 8, 10, 12} //len = 7
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if days := daysInMonth(thirty1[r.Intn(7)]); days != 31 {
		t.Errorf("Expected to have 31 days, but it had %d days instead", days)
	}
}

func TestDaysInMonth30(t *testing.T) {
	var thirty []int = []int{4, 6, 9, 11} //len = 4
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if days := daysInMonth(thirty[r.Intn(4)]); days != 30 {
		t.Errorf("Expected to have 30 days, but it had %d days instead", days)
	}
}

func TestDaysInMonth28(t *testing.T) {
	if days := daysInMonth(2); days != 28 {
		t.Errorf("Expected Febuary to have 28 days, but it had %d days instead", days)
	}
}

func TestisNumTrue(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := strconv.Itoa(r.Intn(12) + 1)
	if !isNum(s, 12) {
		t.Errorf("Expected %s to be between 1 and 12 inclusive, but it isnt?", s)
	}
}

func TestisNumFalse(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	d := 0
	for d <= 12 || d >= 1 {
		d = int(r.Uint32())
	}
	s := strconv.Itoa(d)
	if isNum(s, 12) {
		t.Errorf("Expected %s to not be between 1 and 12 inclusive, but it is?", s)
	}
}

func TestTrimToNumT(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	d := r.Int31()
	for d >= 48 && d <= 57 {
		d = r.Int31()
	}
	if !trimToNum(rune(d)) {
		t.Errorf("Expected %v to not be a digit, but it is?", rune(d))
	}
}

func TestTrimToNumF(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	d := r.Int31()
	for d < 48 || d > 57 {
		d = r.Int31()
	}
	if trimToNum(rune(d)) {
		t.Errorf("Expected %v to not be a digit, but it is?", rune(d))
	}
}

func TestMeridianPM(t *testing.T) {
	if tf := meridianTo24(6, "pm"); tf != 18 {
		t.Errorf("Expected to be 18 in 24hour time, but it is %d", tf)
	}
}

func TestMeridian12PM(t *testing.T) {
	if tf := meridianTo24(12, "pm"); tf != 12 {
		t.Errorf("Expected to be 12 in 24hour time, but it is %d", tf)
	}
}

func TestMeridianAM(t *testing.T) {
	if tf := meridianTo24(7, "am"); tf != 7 {
		t.Errorf("Expected to be 7 in 24hour time, but it is %d", tf)
	}
}

func TestPadNumPadded(t *testing.T) {
	s := padNum(8)
	if len(s) != 2 || s != "08" {
		t.Errorf("Expected 8 to be padded to '08', but it is %s", s)
	}
}

func TestPadNumNotPadded(t *testing.T) {
	s := padNum(10)
	if len(s) != 2 || s != "10" {
		t.Errorf("Expected 10 be '10', but it is %s", s)
	}
}

func TestContainsHour1(t *testing.T) {
	s1 := "1hour"
	s2 := "1hours"
	s3 := "1hr"
	s4 := "1hrs"
	s5 := "1h"
	if !containsHour(s1) {
		t.Errorf("Expected '%s' to contain 'hour'", s1)
	}
	if !containsHour(s2) {
		t.Errorf("Expected '%s' to contain 'hours'", s2)
	}
	if !containsHour(s3) {
		t.Errorf("Expected '%s' to contain 'hr'", s3)
	}
	if !containsHour(s4) {
		t.Errorf("Expected '%s' to contain 'hrs'", s4)
	}
	if !containsHour(s5) {
		t.Errorf("Expected '%s' to contain 'hrs'", s4)
	}
}

func TestContainsMin(t *testing.T) {
	s1 := "1min"
	s2 := "1mins"
	s3 := "1minute"
	s4 := "1minutes"
	if !containsMin(s1) {
		t.Errorf("Expected '%s' to contain 'min'", s1)
	}
	if !containsMin(s2) {
		t.Errorf("Expected '%s' to contain 'min'", s2)
	}
	if !containsMin(s3) {
		t.Errorf("Expected '%s' to contain 'min'", s3)
	}
	if !containsMin(s4) {
		t.Errorf("Expected '%s' to contain 'min'", s4)
	}
}
