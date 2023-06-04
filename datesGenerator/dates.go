package datesGenerator

import (
	"time"
)

func GenerateDates() []string {
	today := time.Now()
	var dates []string
	for i := 0; i < 91; i++ {
		todayDay := today.AddDate(0, 0, -i).Format("02/01/2006")
		dates = append(dates, todayDay)
	}
	return dates
}
