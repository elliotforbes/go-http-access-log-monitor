HTTP Log Monitor
==================



* Display stats every 10s about the traffic during those 10s
    * should display section with the most hits
        i.e. /pages/create -> section == /pages - DONE

* Whenever total traffic for the past 2 minutes exceeds a certain number on average, add a default message saying that:
     "high traffic generated an alert - hits = {value}, triggered at {time}" - DONE
* Whenever total traffic drops again below that value on average for the past 2 minutes, add another message saying that:
    "alert recovered" - DONE

* Make sure all messages showing when alerting thresholds are crossed remain visible on the page for historical reasons - DONE
* Write a test for the alerting logic - DONE

* Make sure a user can keep the app running and monitor the log file continuously - DONE
* Consume an actively written-to HTTP access log
    * default to /tmp/access.log - should be overrideable - DONE
* Explain how you'd improve on this application design - DONE
* Docker build and docker run project - DONE
* There will be something else writing to that file - DONE