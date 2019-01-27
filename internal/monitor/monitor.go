package monitor

import (
	"regexp"
	"strings"
	"time"

	"github.com/elliotforbes/go-http-monitor/internal/config"
	"github.com/elliotforbes/go-http-monitor/internal/logger"
	"github.com/elliotforbes/go-http-monitor/internal/reporter"
	"github.com/elliotforbes/go-http-monitor/internal/statistics"
	"github.com/hpcloud/tail"
	"github.com/robfig/cron"
)

const (
	timelayout = "02/Jan/2006:15:04:05 -0700"
)

// Monitor a simple HTTP Log monitoring client
type Monitor struct {
	Config *config.Config
	// Reporter to report any stats collected by the monitor
	Reporter reporter.Reporter
}

// Start is the main entry point for our log monitoring
// system. This takes in a filepath and kicks off monitoring
// and reporting.
func (m *Monitor) Start() {
	logger.Log.Println("Starting HTTP Access Log Monitor")

	// Start the Goroutine that handles dislpaying to the console
	recorder := statistics.NewRecorder()
	go m.Reporter.Output(recorder)
	m.watchAccessLog(m.Config.FilePath, recorder)
}

// watchAccessLog takes in a string path to a file and
// begins tailing that file for new log records. When new log records
// are recorded, it coordinates parsing and the gathering of statistics
func (m *Monitor) watchAccessLog(filepath string, stats *statistics.Recorder) {

	// this tails the log file in question, or throws an error
	t, err := tail.TailFile(filepath, tail.Config{Follow: true})
	if err != nil {
		logger.Log.Println("Error Tailing File: ", err.Error())
		panic(err)
	}

	logger.Log.Println("Starting Cron Job to Flush Old Log Records")
	c := cron.New()
	c.AddFunc("@every 1s", stats.FlushOldRecords)
	logger.Log.Println("Starting Cron Job to Watch for Alerts")
	c.AddFunc("@every 1s", func() {
		stats.CheckAlerts(m.Config.AlertThreshold)
	})
	c.Start()

	logger.Log.Println("Listening for new lines to be added to the file")
	// listen for new lines written to the log file.
	for line := range t.Lines {
		request, err := ParseRequest(line.Text)
		if err != nil {
			logger.Log.Println("Error Parsing Request: ", err.Error())
		}

		// only start tracking requests in the file if it's less than 2 minutes old
		if request.Timestamp.After(time.Now().Add(time.Duration(-2) * time.Minute)) {
			stats.RecordRequest(request, m.Config.AlertThreshold)
		}
	}
}

// parseRequest takes in a HTTP Access Log record and parses
// it into a Request struct which is then returned to the calling
// function
func ParseRequest(input string) (statistics.Request, error) {
	// find the request and trim the quotes
	r, err := regexp.Compile(`"(.*?)"`)
	if err != nil {
		return statistics.Request{}, err
	}
	requestStr := r.FindString(input)
	requestStr = strings.TrimPrefix(requestStr, "\"")
	requestStr = strings.TrimSuffix(requestStr, "\"")
	sections := strings.Split(requestStr, " ")

	section := strings.Split(sections[1], "/")
	requestSection := "/" + section[1]

	// find timestamp and trim the square brackets
	timeRegex, err := regexp.Compile(`\[([^\[\]]*)\]`)
	if err != nil {
		return statistics.Request{}, err
	}
	timestamp := timeRegex.FindString(input)
	timestamp = strings.TrimPrefix(timestamp, "[")
	timestamp = strings.TrimSuffix(timestamp, "]")

	// parse the timestamp to get a time.Time
	logTime, err := time.Parse(timelayout, timestamp)
	if err != nil {
		return statistics.Request{}, err
	}

	request := statistics.Request{
		Verb:      sections[0],
		Section:   requestSection,
		Protocol:  sections[2],
		Timestamp: logTime,
	}
	return request, nil
}
