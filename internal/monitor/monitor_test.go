package monitor

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var logPrefixMessages = []string{
	`127.0.0.1 - james [`,
	`127.0.0.1 - james [`,
	`127.0.0.1 - jill [`,
	`127.0.0.1 - frank [`,
	`127.0.0.1 - mary [`,
	`127.0.0.1 - mary [`,
	`127.0.0.1 - mary [`,
	`127.0.0.1 - mary [`,
	`127.0.0.1 - mary [`,
	`127.0.0.1 - mary [`,
}

var logSuffixMessages = []string{
	`] "GET /report HTTP/1.0" 200 123`,
	`] "GET /report HTTP/1.0" 200 123`,
	`] "GET /api/user HTTP/1.0" 200 23`,
	`] "POST /api/user HTTP/1.0" 200 34`,
	`] "POST /api/user HTTP/1.0" 503 12`,
	`] "POST /topics HTTP/1.0" 503 12`,
	`] "POST /categories HTTP/1.0" 503 12`,
	`] "POST /categories/news HTTP/1.0" 503 12`,
	`] "POST /categories/politics HTTP/1.0" 503 12`,
	`] "POST /api/report HTTP/1.0" 503 12`,
}

func TestMonitor(t *testing.T) {
	fmt.Println("Test program that contiunally writes to access.log")

	for i := 0; i < 5; i++ {

		f, err := os.OpenFile("access.log", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		fmt.Println("Writing New Log Record")

		timestamp := time.Now().Add(time.Duration(-rand.Intn(119)) * time.Second)

		_, err = fmt.Fprintln(f, logPrefixMessages[rand.Intn(10)]+timestamp.Format(timelayout)+logSuffixMessages[rand.Intn(10)])
		if err != nil {
			fmt.Println(err)
		}
	}
}

func TestParseRequest(t *testing.T) {
	testLogs := []string{
		`127.0.0.1 - james [27/Jan/2019:17:25:48 +0000] "POST /topics HTTP/1.0" 503 12`,
		`127.0.0.1 - mary [27/Jan/2019:17:25:29 +0000] "POST /categories/politics HTTP/1.0" 503 12`,
		``,
	}

	request1 := ParseRequest(testLogs[0])
	assert.Equal(t, request1.Path, "/topics")

	request2 := ParseRequest(testLogs[1])
	assert.Equal(t, request2.Path, "/categories")

	request3 := ParseRequest(testLogs[1])
	assert.Equal(t, request2.Path, "/categories")

}
