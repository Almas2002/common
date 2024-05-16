package prometheus

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

type Config struct {
	Port string
	Path string
}

type prometheusClient struct {
	cfg *Config
}

func New(cfg *Config) *prometheusClient {
	return &prometheusClient{cfg}
}

func (p *prometheusClient) RunServer() error {
	http.Handle(p.cfg.Path, promhttp.Handler())
	return http.ListenAndServe(":"+p.cfg.Port, nil)
}

func (p *prometheusClient) RunServerMust() {
	http.Handle(p.cfg.Path, promhttp.Handler())

	log.Fatal(http.ListenAndServe(":"+p.cfg.Port, nil))
}

func (p *prometheusClient) Handler() http.Handler {
	return promhttp.Handler()
}
