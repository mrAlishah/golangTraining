package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

/*
1- go run main.go
2- curl localhost:8081/metrics
*/
func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8081", nil)
}
