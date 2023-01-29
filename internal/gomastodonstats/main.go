package gomastodonstats

import "log"

func Run() {
	log.Println("Fetching counts")

	// Get Counts
	metrics, err := getUserCounts()
	if err != nil {
		log.Fatal(err)
	}

	// Write to DB
	updatedMetrics := persistMetrics(metrics)

	sendToMatrix(updatedMetrics)
}
