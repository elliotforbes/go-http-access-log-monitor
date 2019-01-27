package config

import (
	"flag"
)

// Config stores the AlertThreshold as an integer value
// and FilePath which stores the path to the log file we wish
// to monitor
type Config struct {
	AlertThreshold int
	FilePath       string
	TimeFrame      int
}

// NewConfig parses any flags that the program was
// started with and returns a pointer to the populated Config struct
func NewConfig() *Config {
	file := flag.String("file", "/tmp/access.log", "The HTTP Access file you wish to monitor")
	threshold := flag.Int("threshold", 10, "The minimum threshold for accesses triggering a high usage alert")
	time := flag.Int("time", 10, "The timeframe which we want to monitor, defaults to 10s")
	flag.Parse()

	return &Config{
		AlertThreshold: *threshold,
		FilePath:       *file,
		TimeFrame:      *time,
	}
}
