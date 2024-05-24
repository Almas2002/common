package metrics

import (
	"github.com/Almas2002/common/internal/metrics"
	"time"
)

type Metric struct {
	metric *metrics.Metric
}

func New(serviceName string) *Metric {

	return &Metric{metrics.New(serviceName)}
}

func (m *Metric) IncKafkaRequest() {
	m.metric.TotalMessageFromKafka.Inc()
}

func (m *Metric) IncKafkaResponse() {
	m.metric.TotalMessageToKafka.Inc()
}

func (m *Metric) DurationTime(err error, start time.Time) {
	msg := "success"
	if err != nil {
		msg = err.Error()
	}
	since := time.Since(start)
	m.metric.TotalMessageValid.WithLabelValues(msg).Observe(since.Seconds())
}
