package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Главная страница
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		version := os.Getenv("APP_VERSION")
		if version == "" {
			version = "dev"
		}
		fmt.Fprintf(w, "Hello from Kubernetes! Version: %s\n", version)
	})

	// Эндпоинт для проверки здоровья (Liveness/Readiness probes)
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Эндпоинт для метрик (Prometheus будет сюда ходить)
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "# HELP http_requests_total Total number of HTTP requests.\n")
		fmt.Fprintf(w, "# TYPE http_requests_total counter\n")
		fmt.Fprintf(w, "http_requests_total{method=\"get\",code=\"200\"} 1\n")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
