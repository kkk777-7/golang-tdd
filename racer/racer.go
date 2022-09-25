package racer

import (
	"net/http"
	"time"
)

func Racer(urls ...string) string {
	start0 := time.Now()
	http.Get(urls[0])
	duration0 := time.Since(start0)

	start1 := time.Now()
	http.Get(urls[1])
	duration1 := time.Since(start1)

	if duration0 < duration1 {
		return urls[0]
	}
	return urls[1]
}
