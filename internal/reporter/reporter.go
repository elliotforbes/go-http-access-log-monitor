package reporter

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/elliotforbes/go-http-monitor/internal/statistics"
	"github.com/fatih/color"
)

// Reporter interface dictates the contract
// that all reporters must adhere to.
type Reporter interface {
	// Output does the job of outputting the results
	// in a nice, readible manner
	Output(*statistics.Recorder)
}

// ConsoleReporter is a reporter which outputs the stats
// collected by our monitor to the console/terminal
type ConsoleReporter struct{}

// NewReporter returns a pointer to a new ConsoleReporter struct
func NewReporter() *ConsoleReporter {
	return &ConsoleReporter{}
}

// Output for the Console reporter prints out a structured
// table of the stats collected by our monitoring system
// it uses an ANSI control code to clear the console
func (c ConsoleReporter) Output(stats *statistics.Recorder) {
	for {
		fmt.Println("\033[H\033[2J")
		fmt.Println("Website Traffic")
		fmt.Println(time.Now().Format("15:04:05"))
		fmt.Println("---------------------------")

		w := tabwriter.NewWriter(os.Stdout, 0, 2, 2, ' ', tabwriter.Debug)
		fmt.Fprintln(w, "  Section  \t 10s Average \t Total Hits - Last 2 Minutes  \t High Traffic Alert")

		for _, section := range stats.ToSortedList() {
			// TODO: Break this into an unexported function
			alertColor := color.New(color.FgGreen).SprintFunc()
			if stats.Stats[section].HighTrafficAlert {
				alertColor = color.New(color.Bold, color.FgRed).SprintFunc()
			}

			fmt.Fprintf(w, "  %+v  \t  %+v  \t  %+v  \t  %+v  \n", section, stats.Stats[section].Average, len(stats.Stats[section].Hits), alertColor(stats.Stats[section].HighTrafficAlert))
		}
		w.Flush()
		fmt.Println(" ")
		fmt.Println("All Triggered Alerts")
		fmt.Println("---------------------------")
		for _, alert := range stats.Alerts {
			fmt.Println(alert)
		}

		// Only refresh the count every 10 seconds
		time.Sleep(1000 * time.Millisecond)
	}
}
