package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	counter prometheus.Counter
	guage   prometheus.Gauge
)

func sayHello(resp http.ResponseWriter, req *http.Request) {
	counter.Inc()
	guage.SetToCurrentTime()
	fmt.Fprint(resp, "Hello.")
}

func test() {
	counter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "api_count",
		Help: "beego api counter",
	})
	prometheus.MustRegister(counter)
	http.HandleFunc("/", sayHello)
	http.Handle("/metrics", promhttp.Handler())

	guage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "api_gauge",
		Help: "help",
	})
	prometheus.MustRegister(guage)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}

}
