// statistics holds the struct definitions for any stats
// we wish to capture.
package statistics

import (
	"sort"
	"strings"
	"time"

	"github.com/elliotforbes/go-http-monitor/internal/logger"
)

// Request is the struct which hold all of the request
// information from a given HTTP Access Log Record
type Request struct {
	Verb      string
	Path      string
	Protocol  string
	Timestamp time.Time
}

// Stats Recorder
type StatsRecorder struct {
	Stats  map[string]*Stats
	Alerts []string
}

// Stats records the number of hits for a given path
type Stats struct {
	Hits             []Request
	Average          int
	HighTrafficAlert bool
}

// NewRecorder returns a pointer to a new StatsRecorder struct
func NewRecorder() *StatsRecorder {
	return &StatsRecorder{
		Stats: make(map[string]*Stats),
	}
}

// CheckAlerts attempts to see if any alerts need to be triggered
// when a new request is recorded.
func (s *StatsRecorder) CheckAlerts(request Request) {
	// if s.Stats[request.Path].Hits / 120
}

func (s *StatsRecorder) FlushOldRecords() {
	for section, _ := range s.Stats {
		for i := len(s.Stats[section].Hits) - 1; i >= 0; i-- {
			if s.Stats[section].Hits[i].Timestamp.Before(time.Now().Add(time.Duration(-2) * time.Minute)) {
				logger.Log.Printf("Flushing Record: %+v\n", s.Stats[section].Hits[i])
				s.Stats[section].Hits = append(s.Stats[section].Hits[:i], s.Stats[section].Hits[i+1:]...)
			} else {
				// TODO: If this list is guaranteed to be ordered, we can break when the
				break
			}
		}
	}
}

func (s *StatsRecorder) ToSortedSlice() []string {
	paths := make([]string, 0, len(s.Stats))
	for key, _ := range s.Stats {
		paths = append(paths, key)
	}
	sort.Strings(paths)
	return paths
}

func (s *StatsRecorder) RecordRequest(request Request, threshold int) {
	// parse section from path
	sections := strings.Split(request.Path, "/")
	section := "/" + sections[1]

	// If section already exists with map then
	// increment total hits for the last 10 seconds
	if _, ok := s.Stats[section]; ok {
		s.Stats[section].Hits = append(s.Stats[section].Hits, request)
		// TODO: calculate average
	} else {
		s.Stats[section] = &Stats{}
		s.Stats[section].Hits = append(s.Stats[section].Hits, request)
	}

}
