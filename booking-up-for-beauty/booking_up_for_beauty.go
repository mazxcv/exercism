package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed.
func getDate(date string) time.Time {
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	return t
}
func HasPassed(date string) bool {
	now := time.Now()
	return now.After(getDate(date))
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func getDate2(date string) time.Time {
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return t
}
func IsAfternoonAppointment(date string) bool {
	hour := getDate2(date).Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t := Schedule(date)
	description := fmt.Sprintf("You have an appointment on %v.", t.Format("Monday, January 2, 2006, at 15:04"))
	return description
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	anniversaryDate, _ := time.Parse("2/1/2006", "15/09/2024")
	return anniversaryDate
}
