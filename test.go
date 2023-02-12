package main

import (
	"log"
	"time"
)

func main() {
	localTime, err := time.LoadLocation("Pacific/Auckland")
	if err != nil {
		log.Fatal(err)
	}
	// Iterate until Monday!
	t := time.Now().In(localTime)
	for t.Weekday() != time.Monday {
		t = t.AddDate(0, 0, -1)
	}
	year, month, day := t.Date()
	log.Println(time.Date(year, month, day, 0, 0, 0, 0, t.Location()).String())
}
