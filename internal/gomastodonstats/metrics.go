package gomastodonstats

import (
	"fmt"
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

// persistMetrics - return any updated
func persistMetrics(metrics []metric) []metric {
	var updatedMetrics []metric

	startOfDay := getStartofDay()
	for _, v := range metrics {
		v.MetricTime = startOfDay
		err := insertValues(v)
		if err != nil {
			log.Println(err)
		} else {
			updatedMetrics = append(updatedMetrics, v)
		}
	}

	return updatedMetrics
}

func getUserCounts() ([]metric, error) {
	var metrics []metric

	if PIXELFED_DB_SCHEMA != "" {
		userCount, err := runIntQuery(PIXELFED_DB_SCHEMA, PIXELFED_USER_QUERY)
		if err != nil {
			log.Println(err)
		} else {
			m := metric{
				Service:     PIXELFED_IDENTIFIER,
				MetricName:  METRICNAME_USERCOUNT,
				MetricValue: userCount,
			}
			log.Printf("%s user count: %d", PIXELFED_IDENTIFIER, userCount)
			metrics = append(metrics, m)
		}
	}

	if MATRIX_DB_SCHEMA != "" {
		userCount, err := runIntQuery(MATRIX_DB_SCHEMA, MATRIX_USER_QUERY)
		if err != nil {
			log.Println(err)
		} else {
			m := metric{
				Service:     MATRIX_IDENTIFIDER,
				MetricName:  METRICNAME_USERCOUNT,
				MetricValue: userCount,
			}
			log.Printf("%s user count: %d", MATRIX_IDENTIFIDER, userCount)
			metrics = append(metrics, m)
		}
	}

	if MASTODON_DB_SCHEMA != "" {
		userCount, err := runIntQuery(MASTODON_DB_SCHEMA, MASTODON_USER_QUERY)
		if err != nil {
			log.Println(err)
		} else {
			m := metric{
				Service:     MASTODON_IDENTIFIER,
				MetricName:  METRICNAME_USERCOUNT,
				MetricValue: userCount,
			}
			log.Printf("%s user count: %d", MASTODON_IDENTIFIER, userCount)
			metrics = append(metrics, m)
		}
	}

	if MOBILIZON_DB_SCHEMA != "" {
		userCount, err := runIntQuery(MOBILIZON_DB_SCHEMA, MOBILIZON_USER_QUERY)
		if err != nil {
			log.Println(err)
		} else {
			m := metric{
				Service:     MOBILIZON_IDENTIFIER,
				MetricName:  METRICNAME_USERCOUNT,
				MetricValue: userCount,
			}
			log.Printf("%s user count: %d", MOBILIZON_IDENTIFIER, userCount)
			metrics = append(metrics, m)
		}
	}

	if PEERTUBE_DB_SCHEMA != "" {
		userCount, err := runIntQuery(PEERTUBE_DB_SCHEMA, PEERTUBE_USER_QUERY)
		if err != nil {
			log.Println(err)
		} else {
			m := metric{
				Service:     PEERTUBE_IDENTIFIER,
				MetricName:  METRICNAME_USERCOUNT,
				MetricValue: userCount,
			}
			log.Printf("%s user count: %d", PEERTUBE_IDENTIFIER, userCount)
			metrics = append(metrics, m)
		}
	}

	return metrics, nil
}

func getPrintableString(m metric) string {
	return fmt.Sprintf("%s: %d", m.Service, m.MetricValue)
}
