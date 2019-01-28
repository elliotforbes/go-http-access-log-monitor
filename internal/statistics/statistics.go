// statistics holds the struct definitions for any stats
// we wish to capture.
package statistics

import (
	"sort"
	"strconv"
	"time"

	"github.com/elliotforbes/go-http-monitor/internal/logger"
)

// Request is the struct which hold all of the request
// information from a given HTTP Access Log Record
type Request struct {
	Verb      string
	Section   string
	Protocol  string
	Timestamp time.Time
}

// Stats Recorder
type Recorder struct {
	Stats  map[string]*Stats
	Alerts []string
}

// Stats records the number of hits for a given path
type Stats struct {
	Hits             []Request
	Average          int
	HighTrafficAlert bool
}

// NewRecorder returns a pointer to a new Recorder struct
func NewRecorder() *Recorder {
	return &Recorder{
		Stats: make(map[string]*Stats),
	}
}

// CheckAlerts attempts to see if any alerts need to be triggered
// when a new request is recorded.
func (s *Recorder) CheckAlerts(threshold int) {
	// alertLevel = inputted transactions per second * 2 minutes (120s)
	// if alert level surpasses this, record new alert
	alertLevel := threshold * 120
	for section, _ := range s.Stats {
		if len(s.Stats[section].Hits) > alertLevel {

			if s.Stats[section].HighTrafficAlert != true {
				alertMsg := "High Traffic generated an alert - hits at time of alert = " + strconv.Itoa(len(s.Stats[section].Hits)) + ", triggered at " + time.Now().String()
				s.Alerts = append(s.Alerts, alertMsg)
			}
			// display high traffic alert on that particular section
			s.Stats[section].HighTrafficAlert = true
		} else {
			if s.Stats[section].HighTrafficAlert {
				alertMsg := "Alert Recovered: " + time.Now().String()
				s.Alerts = append(s.Alerts, alertMsg)
				s.Stats[section].HighTrafficAlert = false
			}
		}
	}
}

func (s *Recorder) FlushOldRecords() {
	for section, _ := range s.Stats {
		for i := len(s.Stats[section].Hits) - 1; i >= 0; i-- {
			if s.Stats[section].Hits[i].Timestamp.Before(time.Now().Add(time.Duration(-2) * time.Minute)) {
				logger.Log.Printf("Flushing Record: %+v\n", s.Stats[section].Hits[i])
				// I am shifting all elements left one space as opposed to just deleting the element so I
				// can attempt to preserve order within my array
				s.Stats[section].Hits = append(s.Stats[section].Hits[:i], s.Stats[section].Hits[i+1:]...)
			} else {
				// break here if the last element checked is less than 2 minutes old, we
				break
			}
		}
	}
}

// ToSortedList returns a sorted list of paths which can then
// be used to structure output.
func (s *Recorder) ToSortedList() []string {
	paths := make([]string, 0, len(s.Stats))
	for key, _ := range s.Stats {
		paths = append(paths, key)
	}
	sort.Strings(paths)
	return paths
}

// RecordRequest takes in a new Request and a threshold limit and
// records this request in our Recorder.
func (s *Recorder) RecordRequest(request Request, threshold int) {

	// If section already exists with map then
	// increment total hits for the last 10 seconds
	if _, ok := s.Stats[request.Section]; ok {
		s.Stats[request.Section].Hits = append(s.Stats[request.Section].Hits, request)
	} else {
		s.Stats[request.Section] = &Stats{}
		s.Stats[request.Section].Hits = append(s.Stats[request.Section].Hits, request)
	}

}
