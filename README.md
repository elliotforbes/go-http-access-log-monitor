HTTP Access Monitor
=====================

# Docker

```s
$ docker build -t http-access-monitor .
$ docker run -it http-access-monitor
```

# Running Tests

To continually write Log access messages with timestamps within the last 2 minutes, you can run `./testWriter`. If you want to unit test the project, then you can do so by running the following:

```s
$ go test ./... -v
```

# Future Improvements

* There isn't any logic covering what happens if a log file gets rotated. In order to make this more production-ready, we would have to build in logic that would resolve this issue.

* I've only implemented a simple ConsoleReporter which meets the `Reporter` interface. Ideally, if we wanted to make this deployable across a fleet of distributed systems, we could possibly implement a WebSocket Reporter which reports the data from all given systems up to an aggregation layer or something similar. I could then leverage a JavaScript frontend and better display the results.

* Further work could be done on optimizing performance. I've done almost everything in a single-threaded fashion, with the exception of the ConsoleReporter running on a separate goroutine. 
    * Having not run this at any great scale, I've not had a chance to test for bottlenecks

* A lot more work should be done on unit testing to ensure that we don't miss any edge-cases. 

* This application feels like a good fit for using event-based libraries such as RxGo or equivalent. I've only recently started playing about with these libraries and didn't yet feel comfortable using them as the base for this project given the time-span.