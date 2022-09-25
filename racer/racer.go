package racer

import (
	"net/http"
	"time"
)

func Racer(urls ...string) string {
	duration0 := measureResponseTime(urls[0])
	duration1 := measureResponseTime(urls[1])

	if duration0 < duration1 {
		return urls[0]
	}
	return urls[1]
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
