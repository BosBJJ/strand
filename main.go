package main

import (
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	mux := http.NewServeMux()
	srv := http.Server{
		Handler: mux, 
		Addr: ":8080",
	}
	cfg := apiConfig{}
	mux.Handle("/app/", (http.StripPrefix("/app", cfg.middlewareMetricsInc(http.FileServer(http.Dir("."))))))
	mux.HandleFunc("GET /api/healthz", cfg.handlerReadiness)
	mux.HandleFunc("GET /admin/metrics", cfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", cfg.handlerReset)
	mux.HandleFunc("POST /api/validate_chirp", cfg.handlerValidate)
	srv.ListenAndServe()
	
}