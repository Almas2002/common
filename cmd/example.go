package main

import (
	"errors"
	"github.com/Almas2002/common/pkg/metrics"
	"github.com/Almas2002/common/pkg/prometheus"
	"time"
)

func main() {
	met := prometheus.New(&prometheus.Config{
		Port: "5455",
		Path: "/metrics",
	})

	// напишите имя сервиса
	m := metrics.New("FOOOOOOOOOOOOOO")
	go func() {
		for {
			m.IncKafkaRequest()
			start := time.Now()
			err := errors.New("bereke bank")
			time.Sleep(time.Millisecond * 500)
			m.DurationTime(err, start)
			m.DurationTime(nil, start)
			m.IncKafkaResponse()
		}
	}()
	met.RunServerMust()

	go func() {
		if err := met.RunServer(); err != nil {

		}
		return
	}()
}
