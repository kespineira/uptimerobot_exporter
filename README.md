# UptimeRobot Exporter

[![Release](https://img.shields.io/github/release/kespineira/uptimerobot_exporter.svg?style=flat)](https://github.com/kespineira/uptimerobot_exporter/releases/latest)
[![Build](https://github.com/kespineira/uptimerobot_exporter/actions/workflows/build.yaml/badge.svg?branch=main)](https://github.com/kespineira/uptimerobot_exporter/actions/workflows/build.yaml)
[![codecov.io Code Coverage](https://img.shields.io/codecov/c/github/kespineira/uptimerobot_exporter.svg?maxAge=2592000)](https://codecov.io/github/kespineira/uptimerobot_exporter?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/kespineira/uptimerobot_exporter)](https://goreportcard.com/report/github.com/kespineira/uptimerobot_exporter)

The UptimeRobot Exporter is a Prometheus exporter for the UptimeRobot API. It allows you to monitor the status and response times of your websites and services directly from Prometheus. This project is designed to be easy to set up and integrate into your existing Prometheus monitoring environment.

## Quick Start

To get started with the UptimeRobot Exporter quickly, you can use the provided Docker container. Run the following command to start the exporter:

```sh
docker run -d -p 8080:8080 -e UPTIMEROBOT_API_KEY=your_api_key_here kespineira/uptimerobot_exporter
```

Replace `your_api_key_here` with your actual UptimeRobot API key.

## List of Metrics

The UptimeRobot Exporter provides various metrics, including but not limited to:

- **uptimerobot_monitor_status**: Indicates monitored website/service status ( 0 - paused, 1 - not checked yet, 2 - up, 8 - seems down, 9 - down).
- **uptimerobot_monitor_response_time**: The response time of the monitored website/service in milliseconds.

For a complete list of metrics, please refer to the /metrics endpoint of the running exporter.

## License

This project is licensed under the MIT License - see the LICENSE file for details.