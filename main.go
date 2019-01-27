package main

import (
	"github.com/elliotforbes/go-http-monitor/internal/config"
	"github.com/elliotforbes/go-http-monitor/internal/monitor"

	"github.com/elliotforbes/go-http-monitor/internal/reporter"
)

func main() {

	cfg := config.NewConfig()

	consoleReporter := reporter.NewReporter()
	monitor := monitor.Monitor{
		Config:   cfg,
		Reporter: consoleReporter,
	}
	monitor.Start()
}
