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
