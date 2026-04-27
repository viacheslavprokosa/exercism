package booking

import (
    "time"
    "fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err == nil {
		return time.Now().After(t)
	}
	panic(err)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err == nil {
		hour := t.Hour()
		return hour >= 12 && hour < 18
	}
	panic(err)
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	outputDateFormat := "Monday, January 2, 2006"
	outputTimeFormat := "15:04"

	t, err := time.Parse(layout, date)
	if err == nil {
	return fmt.Sprintf("You have an appointment on %s, at %s.",
		t.Format(outputDateFormat), t.Format(outputTimeFormat))
        }
    panic(err)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(),
		time.September,
		15,
		0, 0, 0, 0,
		time.UTC)
}
