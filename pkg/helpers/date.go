package helpers

import "time"

func GetStartAndEndDay() (time.Time, time.Time) {
	now := time.Now()
	currentYear, currentMonth, currentDay := now.Date()
	currentLocation := now.Location()

	firstOfDay := time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, currentLocation).Add(-time.Hour * 24)
	lastOfDay := firstOfDay.AddDate(0, 0, 1).Add(-time.Second)
	return firstOfDay, lastOfDay
}

func GetStartAndEndMonth() (time.Time, time.Time) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	return firstOfMonth, lastOfMonth
}
