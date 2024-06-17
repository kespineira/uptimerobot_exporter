package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Collector interface
type Collector interface {
	Describe(ch chan<- *prometheus.Desc)
	Collect(ch chan<- prometheus.Metric)
}

// RegisterCollectors registers all the collectors
func RegisterCollectors() {
	prometheus.MustRegister(NewUptimeRobotCollector())
}
