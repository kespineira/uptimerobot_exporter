package main

import (
	"log"
	"net/http"

	"github.com/kespineira/uptimerobot_exporter/collector"
	"github.com/kespineira/uptimerobot_exporter/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"github.com/prometheus/exporter-toolkit/web"
)

func main() {
	config.LoadConfig()
	collector.RegisterCollectors()

	http.Handle("/metrics", promhttp.Handler())
	landingConfig := web.LandingConfig{
		Name:        "UpTimeRobot Exporter",
		Description: "Prometheus exporter for UptimeRobot",
		Version:     version.Info(),
		Links: []web.LandingLinks{
			{
				Address: "/metrics",
				Text:    "Metrics",
			},
		},
	}
	landingPage, err := web.NewLandingPage(landingConfig)
	if err != nil {
		log.Fatalf("Error creating landing page: %s", err)
	}
	http.Handle("/", landingPage)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
