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
	Output(statistics.StatsRecorder)
}

// ConsoleReporter is a reporter which outputs the stats
// collected by our monitor to the console/terminal
type ConsoleReporter struct{}

func NewReporter() *ConsoleReporter {
	return &ConsoleReporter{}
}

// Output for the Console reporter prints out a structured
// table of the stats collected by our monitoring system
// it uses an ANSI control code to clear the console
func (c ConsoleReporter) Output(stats statistics.StatsRecorder) {
	for {
		fmt.Println("\033[H\033[2J")
		fmt.Println("Displaying Traffic for last 10 seconds...")
		fmt.Println("---------------------------")

		w := tabwriter.NewWriter(os.Stdout, 0, 2, 2, ' ', tabwriter.Debug)
		fmt.Fprintln(w, "  Section  \t 10s Average \t Total Hits - Last 2 Minutes  \t High Traffic Alert")

		for _, section := range stats.ToSortedSlice() {
			alertColor := color.New(color.FgGreen).SprintFunc()
			if stats.Stats[section].HighTrafficAlert {
				alertColor = color.New(color.Bold, color.FgRed).SprintFunc()
			}

			fmt.Fprintf(w, "  %+v  \t  %+v  \t  %+v  \t  %+v  \n", section, stats.Stats[section].Average, len(stats.Stats[section].Hits), alertColor(stats.Stats[section].HighTrafficAlert))
		}
		fmt.Println("All Triggered Alerts")
		fmt.Println("---------------------------")
		for _, alert := range stats.Alerts {
			fmt.Println(alert)
		}
		w.Flush()

		// Only refresh the count every 10 seconds
		time.Sleep(1000 * time.Millisecond)
	}
}