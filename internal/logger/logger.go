package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	logFile = "agent.log"
)

// Log will log our output to a file as we are currently using the
// terminal to display ongoing results. This will be useful for debugging
var Log = logrus.New()

func init() {
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
	}
	Log.SetOutput(f)
}
