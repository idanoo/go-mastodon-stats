package gomastodonstats

import (
	"log"
	"sort"
)

func Run() {
	log.Println("Fetching counts")

	// Get Counts
	metrics, err := getUserCounts()
	if err != nil {
		log.Fatal(err)
	}

	// Write to DB
	updatedMetrics := persistMetrics(metrics)
	if len(updatedMetrics) > 0 {
		// Sort by counts!
		sort.Slice(updatedMetrics, func(i, j int) bool {
			return updatedMetrics[i].MetricValue > updatedMetrics[j].MetricValue
		})

		// sendToMatrix(updatedMetrics)

		// Only post weekly here
		// localTime, err := time.LoadLocation(TIMEZONE)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// weekday := time.Now().Local().In(localTime).Weekday()
		// if weekday == time.Monday {
		// 	postToMastodon(updatedMetrics)
		// }
	}
}
