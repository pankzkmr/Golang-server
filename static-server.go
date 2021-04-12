package main

import (
	"fmt"
	"log"
        "net/http"
        "time"
	"github.com/gorilla/mux"
        "github.com/prometheus/client_golang/prometheus"
        "github.com/prometheus/client_golang/prometheus/promauto"
        "github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
        go func() {
                for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
                }
        }()
}

var (
        opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
                Name: "app_processed_ops_total",
                Help: "The total number of processed events",
        })
)

func main() {
        recordMetrics()
		router := mux.NewRouter()
		router.Handle("/metrics", promhttp.Handler())
		router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

		fmt.Println("Serving requests on port 5005")
		err := http.ListenAndServe(":5005", router)
		log.Fatal(err)
}

