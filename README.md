HTTP Access Monitor
=====================

# Docker

```s
$ docker build -t http-access-monitor .
$ docker run -it -p 8080:8080 http-access-monitor
```

# Running Tests

```s
$ go run test.go
```


# Future Improvements

* File handler changing if access log files get rotated
* Allow graceful shutdown of watching files