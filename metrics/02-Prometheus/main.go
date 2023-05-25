// https://antonputra.com/monitoring/monitor-golang-with-prometheus/#create-minimal-app
// https://www.youtube.com/watch?v=WUBjlJzI2a0
package main

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ## Define metrics -----------------------------------------------------
type metrics struct {
	devices  prometheus.Gauge
	info     *prometheus.GaugeVec
	upgrades *prometheus.CounterVec
}

// ## Define & Add Gauge to prometheus ------------------------------------
func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{

		devices: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "myapp",
			Name:      "connected_devices",
			Help:      "Number of currently connected devices.",
		}),

		//we need to set a version label with the actual version of the application.
		info: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "myapp",
			Name:      "info",
			Help:      "Information about the My App environment.",
		},
			[]string{"version"}),

		// name it device_upgrade_total and give it a description Number of upgraded devices.
		upgrades: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "myapp",
			Name:      "device_upgrade_total",
			Help:      "Number of upgraded devices.",
		}, []string{"type"}),
	}
	reg.MustRegister(m.devices, m.info, m.upgrades)
	return m
}

// ## App data struct    ------------------------------------
type Device struct {
	ID       int    `json:"id"`
	Mac      string `json:"mac"`
	Firmware string `json:"firmware"`
}

var dvs []Device
var version string

func init() {
	dvs = []Device{
		{1, "5F-33-CC-1F-43-82", "2.1.6"},
		{2, "EF-2B-C4-F5-D6-34", "2.1.6"},
	}

	//Typically this variable will be set using the environment variable or by your CI tool. For the demo, let's just hardcode it in the init() function.
	version = "2.10.5"
}

/*
1- go run main.go
2- curl localhost:8081/metrics
*/
func main() {
	// ## Register prometheus -----------------------------------------------------
	//create a non-global registry without any pre-registered Collectors.
	reg := prometheus.NewRegistry()

	//Optionally, if you want to keep all the golang default metrics, you can use a built-in collector to register it with the custom prometheus register.
	//reg.MustRegister(collectors.NewGoCollector())

	//Then create metrics using the NewMetrics function
	m := NewMetrics(reg)

	// ## Set Metrics Values -----------------------------------------------------
	//Now we can use the devices property of the metrics struct and set it to the current number of connected devices.
	//For that, we simply set it to the number of items in the devices slice.
	m.devices.Set(float64(len(dvs)))
	//we can use the version prometheus label to set the application version and use a constant value of 1.
	m.info.With(prometheus.Labels{"version": version}).Set(1)

	// ## Set Http Handlers -----------------------------------------------------
	pMux := http.NewServeMux()
	//create a custom prometheus handler with the newly created register.
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	//Also, you can expose the prometheus handler metric as well by adding setting the Registry field.
	//promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})

	pMux.Handle("/metrics", promHandler)
	//http.Handle("/metrics", promhttp.Handler())

	go func() {
		log.Fatal(http.ListenAndServe(":8081", pMux))
	}()

	dMux := http.NewServeMux()
	rdh := registerDevicesHandler{metrics: m}
	mdh := manageDevicesHandler{metrics: m}

	dMux.Handle("/devices", rdh)
	dMux.Handle("/devices/", mdh)

	go func() {
		log.Fatal(http.ListenAndServe(":8080", dMux))
	}()

	select {}
}

type registerDevicesHandler struct {
	metrics *metrics
}

// curl -d '{"id": 3, "mac": "96-40-D1-32-D7-1A", "firmware": "3.03.00"}' localhost:8080/devices
// curl -X POST localhost:8080/devices -H "Content-Type: application/json" -d '{"id": 3, "mac": "96-40-D1-32-D7-1A", "firmware": "3.03.00"}'
// curl localhost:8080/devices
func (rdh registerDevicesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getDevices(w, r)
	case "POST":
		createDevice(w, r, rdh.metrics)
	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
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

func createDevice(w http.ResponseWriter, r *http.Request, m *metrics) {
	var dv Device

	err := json.NewDecoder(r.Body).Decode(&dv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dvs = append(dvs, dv)

	// m.devices.Inc()
	m.devices.Set(float64(len(dvs)))

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Device created!"))
}

type manageDevicesHandler struct {
	metrics *metrics
}

// curl -X PUT -d '{"firmware": "2.3.0"}' localhost:8080/devices/1
//
//	hey -n 100000 -c 1 -q 2 -m PUT -d '{"firmware": "2.03.00"}' http://localhost:8080/devices/1
func (mdh manageDevicesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		upgradeDevice(w, r, mdh.metrics)
	default:
		w.Header().Set("Allow", "PUT")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func upgradeDevice(w http.ResponseWriter, r *http.Request, m *metrics) {
	path := strings.TrimPrefix(r.URL.Path, "/devices/")

	id, err := strconv.Atoi(path)
	if err != nil || id < 1 {
		http.NotFound(w, r)
	}

	var dv Device
	err = json.NewDecoder(r.Body).Decode(&dv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := range dvs {
		if dvs[i].ID == id {
			dvs[i].Firmware = dv.Firmware
		}
	}

	sleep(1000)

	m.upgrades.With(prometheus.Labels{"type": "router"}).Inc()

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Upgrading..."))
}

func sleep(ms int) {
	rand.Seed(time.Now().UnixNano())
	now := time.Now()
	n := rand.Intn(ms + now.Second())
	time.Sleep(time.Duration(n) * time.Millisecond)
}
