package main

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type metrics struct {
	devices prometheus.Gauge
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		devices: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "myapp",
			Name:      "connected_devices",
			Help:      "Number of currently connected devices.",
		}),
	}
	reg.MustRegister(m.devices)
	return m
}

type Device struct {
	ID       int    `json:"id"`
	Mac      string `json:"mac"`
	Firmware string `json:"firmware"`
}

var dvs []Device

func init() {
	dvs = []Device{
		{1, "5F-33-CC-1F-43-82", "2.1.6"},
		{2, "EF-2B-C4-F5-D6-34", "2.1.6"},
	}
}

/*
1- go run main.go
2- curl localhost:8081/metrics
*/
func main() {
	//create a non-global registry without any pre-registered Collectors.
	reg := prometheus.NewRegistry()

	//Optionally, if you want to keep all the golang default metrics, you can use a built-in collector to register it with the custom Prometheus register.
	//reg.MustRegister(collectors.NewGoCollector())

	//Then create metrics using the NewMetrics function
	m := NewMetrics(reg)

	//Now we can use the devices property of the metrics struct and set it to the current number of connected devices.
	//For that, we simply set it to the number of items in the devices slice.
	m.devices.Set(float64(len(dvs)))

	//create a custom prometheus handler with the newly created register.
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	//Also, you can expose the prometheus handler metric as well by adding setting the Registry field.
	//promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})

	http.Handle("/metrics", promHandler)
	//http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/devices", getDevices)

	http.ListenAndServe(":8081", nil)
}

func getDevices(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(dvs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
