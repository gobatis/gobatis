package main

import (
	"github.com/gozelle/prometheus/prometheus"
	"github.com/gozelle/prometheus/prometheus/promhttp"
	"net/http"
	"time"
)

func main() {
	var (
		customCounter = prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "gobatis_connections",
				Help: "Description Database Connections",
			},
		)
	)
	prometheus.MustRegister(customCounter)
	
	go func() {
		for {
			customCounter.Inc()
			time.Sleep(3 * time.Second)
		}
	}()
	
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
