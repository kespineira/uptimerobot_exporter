package main

import (
	"log"
	"net/http"

	"github.com/kespineira/uptimerobot_exporter/collector"
	"github.com/kespineira/uptimerobot_exporter/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	config.LoadConfig()
	collector.RegisterCollectors()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
