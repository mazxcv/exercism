package meetup

import "time"

type WeekSchedule int

const (
	First WeekSchedule = iota + 1
	Second
	Third
	Fourth
	Last
	Teenth
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	panic("Please implement the Day function")
}
