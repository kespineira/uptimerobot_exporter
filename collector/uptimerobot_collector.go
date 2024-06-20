package collector

import (
	"log"
	"strconv"

	"github.com/kespineira/uptimerobot_exporter/internal/uptimerobot"
	"github.com/prometheus/client_golang/prometheus"
)

// UptimeRobotCollector is a custom Prometheus collector for UptimeRobot metrics
type UptimeRobotCollector struct {
	responseTime   *prometheus.Desc
	status         *prometheus.Desc
	accountDetails *prometheus.Desc
	client         *uptimerobot.Client
}

// NewUptimeRobotCollector creates a new UptimeRobotCollector
func NewUptimeRobotCollector() *UptimeRobotCollector {
	client := uptimerobot.NewClient()
	return &UptimeRobotCollector{
		responseTime: prometheus.NewDesc(
			"uptimerobot_monitor_response_time",
			"Response time of the monitor",
			[]string{"fiendly_name", "type", "url"},
			nil,
		),
		status: prometheus.NewDesc(
			"uptimerobot_monitor_status",
			"Status of the monitor",
			[]string{"friendly_name", "http_username", "interval", "keyword_type", "keyword_value", "port", "sub_type", "type", "url"},
			nil,
		),
		accountDetails: prometheus.NewDesc(
			"uptimerobot_accountdetails",
			"Details of the account",
			[]string{"monitor_limit", "up_monitors", "down_monitors", "paused_monitors"},
			nil,
		),
		client: client,
	}
}

// Describe sends the super-set of all possible descriptors of metrics collected by this Collector.
func (c *UptimeRobotCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.responseTime
	ch <- c.status
	ch <- c.accountDetails
}

// Collect is called by the Prometheus registry when collecting metrics.
func (c *UptimeRobotCollector) Collect(ch chan<- prometheus.Metric) {
	monitors, err := c.client.GetMonitors()
	if err != nil {
		log.Println("Error fetching monitors:", err)
		return
	}

	for _, monitor := range monitors {
		ch <- prometheus.MustNewConstMetric(
			c.responseTime,
			prometheus.GaugeValue,
			float64(monitor.ResponseTimes[0].Value),
			monitor.FriendlyName,
			strconv.Itoa(monitor.Type),
			monitor.URL,
		)
		ch <- prometheus.MustNewConstMetric(
			c.status,
			prometheus.GaugeValue,
			float64(monitor.Status),
			monitor.FriendlyName,
			monitor.HttpUsername,
			strconv.Itoa(monitor.Interval),
			monitor.KeywordType,
			monitor.KeywordValue,
			monitor.Port,
			monitor.SubType,
			strconv.Itoa(monitor.Type),
			monitor.URL,
		)
	}

	accountDetails, err := c.client.GetAccountDetails()
	if err != nil {
		log.Println("Error fetching account details:", err)
		return
	}

	ch <- prometheus.MustNewConstMetric(
		c.accountDetails,
		prometheus.GaugeValue,
		1,
		strconv.Itoa(accountDetails.MonitorLimit),
		strconv.Itoa(accountDetails.UpMonitors),
		strconv.Itoa(accountDetails.DownMonitors),
		strconv.Itoa(accountDetails.PausedMonitors),
	)
}
