package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

type Metric struct {
	TotalMessageFromKafka prometheus.Counter
	TotalMessageToKafka   prometheus.Counter
	TotalMessageValid     *prometheus.HistogramVec
}

func New(serviceName string) *Metric {
	m := &Metric{
		TotalMessageFromKafka: prometheus.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_total_message_from_kafka", serviceName),
			Help: "Total message from kafka",
		}),
		TotalMessageToKafka: prometheus.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_total_message_to_kafka", serviceName),
			Help: "Total message to kafka",
		}),
		TotalMessageValid: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    fmt.Sprintf("%s_total_message", serviceName),
			Help:    "Total message",
			Buckets: prometheus.DefBuckets,
		},
			[]string{"result"}),
	}

	prometheus.MustRegister(m.TotalMessageFromKafka, m.TotalMessageValid, m.TotalMessageToKafka)

	return m
}
