// Command backend is the minimal HTTP service of the application.
package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := ":" + port()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthHandler)

	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("backend listening on %s", addr)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

// port returns the port from the PORT environment variable or the default "8080".
func port() string {
	if p := os.Getenv("PORT"); p != "" {
		return p
	}
	return "8080"
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK\n"))
}
