package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	counter prometheus.Counter
)

func sayHello(resp http.ResponseWriter, req *http.Request) {
	counter.Inc()
	fmt.Fprint(resp, "Hello.")
}

func main() {
	counter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "api_count",
		Help: "beego api counter",
	})
	prometheus.MustRegister(counter)
	http.HandleFunc("/", sayHello)
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
