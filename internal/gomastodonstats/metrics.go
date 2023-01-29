package gomastodonstats

import (
	"log"
	"time"
)

// Stores out metric/row data
type metric struct {
	Service     string    `json:"service"`
	MetricName  string    `json:"metric_name"`
	MetricTime  time.Time `json:"metric_time"`
	MetricValue int       `json:"metric_value"`
}

func persistMetrics(metrics []metric) {
	startOfDay := getStartofDay()
	for _, v := range metrics {
		v.MetricTime = startOfDay
		err := insertValues(v)
		if err != nil {
			log.Println(err)
		}
	}
}

func getUserCounts() ([]metric, error) {
	var metrics []metric

	if PIXELFED_DB_SCHEMA != "" {
		pfUsers, err := runIntQuery(PIXELFED_DB_SCHEMA, PIXELFED_USER_QUERY)
		if err != nil {
			log.Println(err)
		} else {
			pfMetric := metric{
				Service:     "pixelfed",
				MetricName:  "userCount",
				MetricValue: pfUsers,
			}
			log.Printf("Pixelfed user count: %d", pfUsers)
			metrics = append(metrics, pfMetric)
		}

	}

	return metrics, nil
}
