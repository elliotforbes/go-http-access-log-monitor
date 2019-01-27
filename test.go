package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	timelayout = "02/Jan/2006:15:04:05 -0700"
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

func main() {
	fmt.Println("Test program that contiunally writes to access.log")

	for {

		f, err := os.OpenFile("access.log", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		fmt.Println("Writing New Log Record")

		timestamp := time.Now().Add(time.Duration(-rand.Intn(119)) * time.Second)
		time.Sleep(200 * time.Millisecond)
		_, err = fmt.Fprintln(f, logPrefixMessages[rand.Intn(10)]+timestamp.Format(timelayout)+logSuffixMessages[rand.Intn(10)])
		if err != nil {
			fmt.Println(err)
		}
	}
}
