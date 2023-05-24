package gomastodonstats

import (
	"fmt"
	"log"
	"time"
)

// Stores out metric/row data
type metric struct {
	Service                 string    `json:"service"`
	MetricName              string    `json:"metric_name"`
	MetricTime              time.Time `json:"metric_time"`
	MetricValue             int       `json:"metric_value"`
	PreviousDayMetricValue  int       `json:"-"`
	PreviousWeekMetricValue int       `json:"-"`
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
				Service:                 PIXELFED_IDENTIFIER,
				MetricName:              METRICNAME_USERCOUNT,
				MetricValue:             userCount,
				PreviousDayMetricValue:  getLastMetric(PIXELFED_IDENTIFIER),
				PreviousWeekMetricValue: getLastWeekMetric(PIXELFED_IDENTIFIER),
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
				Service:                 MATRIX_IDENTIFIDER,
				MetricName:              METRICNAME_USERCOUNT,
				MetricValue:             userCount,
				PreviousDayMetricValue:  getLastMetric(MATRIX_IDENTIFIDER),
				PreviousWeekMetricValue: getLastWeekMetric(MATRIX_IDENTIFIDER),
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
				Service:                 MASTODON_IDENTIFIER,
				MetricName:              METRICNAME_USERCOUNT,
				MetricValue:             userCount,
				PreviousDayMetricValue:  getLastMetric(MASTODON_IDENTIFIER),
				PreviousWeekMetricValue: getLastWeekMetric(MASTODON_IDENTIFIER),
			}
			log.Printf("%s user count: %d", MASTODON_IDENTIFIER, userCount)
			metrics = append(metrics, m)
		}

		// Extra metrics for DB only
		userCount1W, err := runIntQuery(MASTODON_DB_SCHEMA, MASTODON_1W_ACTIVE_USER_QUERY)
		if err != nil {
			log.Println(err)
		} else {
			customMetrics := []metric{
				{
					Service:     MASTODON_IDENTIFIER,
					MetricName:  METRICNAME_1W_USERCOUNT,
					MetricValue: userCount1W,
				},
			}

			_ = persistMetrics(customMetrics)
		}
	}

	// Second masto instance
	if TINKERNZ_DB_SCHEMA != "" {
		userCount, err := runIntQuery(TINKERNZ_DB_SCHEMA, MASTODON_USER_QUERY)
		if err != nil {
			log.Println(err)
		} else {
			m := metric{
				Service:                 TINKERNZ_IDENTIFIER,
				MetricName:              METRICNAME_USERCOUNT,
				MetricValue:             userCount,
				PreviousDayMetricValue:  getLastMetric(TINKERNZ_IDENTIFIER),
				PreviousWeekMetricValue: getLastWeekMetric(TINKERNZ_IDENTIFIER),
			}
			log.Printf("%s user count: %d", TINKERNZ_IDENTIFIER, userCount)
			metrics = append(metrics, m)
		}
	}

	if MOBILIZON_DB_SCHEMA != "" {
		userCount, err := runIntQuery(MOBILIZON_DB_SCHEMA, MOBILIZON_USER_QUERY)
		if err != nil {
			log.Println(err)
		} else {
			m := metric{
				Service:                 MOBILIZON_IDENTIFIER,
				MetricName:              METRICNAME_USERCOUNT,
				MetricValue:             userCount,
				PreviousDayMetricValue:  getLastMetric(MOBILIZON_IDENTIFIER),
				PreviousWeekMetricValue: getLastWeekMetric(MOBILIZON_IDENTIFIER),
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
				Service:                 PEERTUBE_IDENTIFIER,
				MetricName:              METRICNAME_USERCOUNT,
				MetricValue:             userCount,
				PreviousDayMetricValue:  getLastMetric(PEERTUBE_IDENTIFIER),
				PreviousWeekMetricValue: getLastWeekMetric(PEERTUBE_IDENTIFIER),
			}
			log.Printf("%s user count: %d", PEERTUBE_IDENTIFIER, userCount)
			metrics = append(metrics, m)
		}
	}

	if BOOKWYRM_DB_SCHEMA != "" {
		userCount, err := runIntQuery(BOOKWYRM_DB_SCHEMA, BOOKWYRM_USER_QUERY)
		if err != nil {
			log.Println(err)
		} else {
			m := metric{
				Service:                 BOOKWYRM_IDENTIFIER,
				MetricName:              METRICNAME_USERCOUNT,
				MetricValue:             userCount,
				PreviousDayMetricValue:  getLastMetric(BOOKWYRM_IDENTIFIER),
				PreviousWeekMetricValue: getLastWeekMetric(BOOKWYRM_IDENTIFIER),
			}
			log.Printf("%s user count: %d", BOOKWYRM_IDENTIFIER, userCount)
			metrics = append(metrics, m)
		}
	}

	if CALCKEY_DB_SCHEMA != "" {
		userCount, err := runIntQuery(CALCKEY_DB_SCHEMA, CALCKEY_USER_QUERY)
		if err != nil {
			log.Println(err)
		} else {
			m := metric{
				Service:                 CALCKEY_IDENTIFIER,
				MetricName:              METRICNAME_USERCOUNT,
				MetricValue:             userCount,
				PreviousDayMetricValue:  getLastMetric(CALCKEY_IDENTIFIER),
				PreviousWeekMetricValue: getLastWeekMetric(CALCKEY_IDENTIFIER),
			}
			log.Printf("%s user count: %d", CALCKEY_IDENTIFIER, userCount)
			metrics = append(metrics, m)
		}
	}

	if WRITEAS_DB_SCHEMA != "" {
		userCount, err := runMySqlIntQuery(WRITEAS_DB_SCHEMA, WRITEAS_USER_QUERY)
		if err != nil {
			log.Println(err)
		} else {
			m := metric{
				Service:                 WRITEAS_IDENTIFIER,
				MetricName:              METRICNAME_USERCOUNT,
				MetricValue:             userCount,
				PreviousDayMetricValue:  getLastMetric(WRITEAS_IDENTIFIER),
				PreviousWeekMetricValue: getLastWeekMetric(WRITEAS_IDENTIFIER),
			}
			log.Printf("%s user count: %d", WRITEAS_IDENTIFIER, userCount)
			metrics = append(metrics, m)
		}
	}

	return metrics, nil
}

func getPrintableString(m metric, useWeek bool) string {
	output := fmt.Sprintf("%s: %d", SERVICE_LINKS[m.Service], m.MetricValue)
	diff := m.MetricValue - m.PreviousDayMetricValue
	if useWeek {
		diff = m.MetricValue - m.PreviousWeekMetricValue
	}

	if diff < 0 {
		output = fmt.Sprintf("%s (%d)", output, diff)
	} else if diff > 0 {
		output = fmt.Sprintf("%s (+%d)", output, diff)
	}

	return output
}

func getLastMetric(serviceName string) int {
	val, err := runIntQuery(
		POSTGRESQL_STATS_DB,
		fmt.Sprintf(
			"SELECT metric_value FROM %s WHERE metric_name = '%s' AND service = '%s' ORDER BY metric_time DESC LIMIT 1",
			POSTGRESQL_STATS_TABLE,
			METRICNAME_USERCOUNT,
			serviceName,
		),
	)

	if err != nil {
		log.Println(err)
		return 0
	}

	return val
}

func getLastWeekMetric(serviceName string) int {
	monday := getStartofDayMonday()
	val, err := runIntQueryWithTime(
		POSTGRESQL_STATS_DB,
		fmt.Sprintf(
			"SELECT metric_value FROM %s WHERE metric_name = '%s' AND service = '%s' AND metric_time = $1 LIMIT 1",
			POSTGRESQL_STATS_TABLE,
			METRICNAME_USERCOUNT,
			serviceName,
		),
		monday,
	)

	if err != nil {
		log.Println(err)
		return 0
	}

	return val
}
