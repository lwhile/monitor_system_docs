package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	counter         prometheus.Counter
	timer           *prometheus.Timer
	requestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "api_response_time",
		Help:    "Histogram for the runtime of a simple example function.",
		Buckets: prometheus.LinearBuckets(0.01, 0.01, 10),
	})
)

func sayHello(resp http.ResponseWriter, req *http.Request) {
	counter.Add(1)
	timer = prometheus.NewTimer(requestDuration)
	//timer.ObserveDuration()
	fmt.Fprint(resp, "Hello.")
}

func main() {
	counter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "api_count",
		Help: "test help",
		//ConstLabels: prometheus.Labels{"a": "1", "b": "2"},
	})
	prometheus.MustRegister(counter)
	http.HandleFunc("/", sayHello)
	go http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
