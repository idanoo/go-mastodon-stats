package gomastodonstats

import (
	"log"
	"time"
)

func getStartofDay() time.Time {
	localTime, err := time.LoadLocation(TIMEZONE)
	if err != nil {
		log.Fatal(err)
	}
	t := time.Now().In(localTime)
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func getStartofDayMonday() time.Time {
	localTime, err := time.LoadLocation(TIMEZONE)
	if err != nil {
		log.Fatal(err)
	}
	// Iterate until Monday!
	t := time.Now().In(localTime)
	t = t.AddDate(0, 0, -1)
	for t.Weekday() != time.Monday {
		t = t.AddDate(0, 0, -1)
	}
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}
