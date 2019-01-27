package statistics

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	timelayout = "02/Jan/2006:15:04:05 -0700"
)

func TestRecordRequest(t *testing.T) {
	testRecorder := NewRecorder()

	now := time.Now().Add(time.Duration(20) * time.Second)
	dummyRequest := Request{
		Path:      "/api",
		Verb:      "GET",
		Protocol:  "HTTP",
		Timestamp: now,
	}
	testRecorder.RecordRequest(dummyRequest, 10)
	fmt.Printf("%+v\n", testRecorder.Stats["/api"])
	assert.Equal(t, len(testRecorder.Stats["/api"].Hits), 1)
}

func TestFlushOldRecords(t *testing.T) {
	testRecorder := NewRecorder()

	now := time.Now()
	dummyRequest := Request{
		Path:      "/api",
		Verb:      "GET",
		Protocol:  "HTTP",
		Timestamp: now,
	}
	testRecorder.RecordRequest(dummyRequest, 10)
	fmt.Printf("%+v\n", testRecorder.Stats["/api"])
	assert.Equal(t, len(testRecorder.Stats["/api"].Hits), 1)

	// simulate 3 minutes going by
	threeMinsAgo := time.Now().Add(time.Duration(-180) * time.Second)
	testRecorder.Stats["/api"].Hits[0].Timestamp = threeMinsAgo
	testRecorder.FlushOldRecords()
	assert.Equal(t, len(testRecorder.Stats["/api"].Hits), 0)

}
