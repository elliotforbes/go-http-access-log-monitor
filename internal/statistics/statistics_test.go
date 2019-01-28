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
		Section:   "/api",
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
		Section:   "/api",
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

func TestCheckAlerts(t *testing.T) {
	testRecorder := NewRecorder()

	now := time.Now()
	dummyRequest := Request{
		Section:   "/api",
		Verb:      "GET",
		Protocol:  "HTTP",
		Timestamp: now,
	}
	for i := 0; i < 1201; i++ {
		testRecorder.RecordRequest(dummyRequest, 10)
	}
	assert.Equal(t, len(testRecorder.Stats["/api"].Hits), 1201)

	testRecorder.CheckAlerts(10)
	assert.Equal(t, len(testRecorder.Alerts), 1)

	threeMinsAgo := time.Now().Add(time.Duration(-180) * time.Second)
	for i := 0; i < len(testRecorder.Stats["/api"].Hits); i++ {
		testRecorder.Stats["/api"].Hits[i].Timestamp = threeMinsAgo
	}
	testRecorder.FlushOldRecords()
	testRecorder.CheckAlerts(10)

	assert.Equal(t, 2, len(testRecorder.Alerts))

}
