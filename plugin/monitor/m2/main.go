package main

import (
	"github.com/rcrowley/go-metrics"
	"log"
	"os"
	"time"
)

func main() {
	c := metrics.NewCounter()
	metrics.MustRegister("foo", c)
	c.Inc(47)
	
	//http.Handle("/metrics", metrics.DefaultRegistry)
	//http.ListenAndServe(":8080", nil)
	
	go metrics.Log(metrics.DefaultRegistry, 5*time.Second, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))
	
	time.Sleep(10 * time.Second)
}
