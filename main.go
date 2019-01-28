package main

import (
	"github.com/elliotforbes/go-http-monitor/internal/config"
	"github.com/elliotforbes/go-http-monitor/internal/monitor"

	"github.com/elliotforbes/go-http-monitor/internal/reporter"
)

func main() {
	// Sets up the configuration for our project based on passed in flags
	cfg := config.NewConfig()

	// creates a new console reporter which outputs stats + alerts to the console
	// in a structured manner
	consoleReporter := reporter.NewReporter()
	// Creates a new Log monitor which is composed of our consoleReporter and our
	// config
	monitor := monitor.Monitor{
		Config:   cfg,
		Reporter: consoleReporter,
	}
	// kicks off our monitor :)
	monitor.Start()
}
