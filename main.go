package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// CLI flags
	ip := flag.String("ip", "0.0.0.0", "IP address to bind to")
	port := flag.Int("port", 0, "Port to listen on (if 0, random 50000–60000)")
	flag.Parse()

	// Randomize port if not specified
	rand.Seed(time.Now().UnixNano())
	if *port == 0 {
		*port = rand.Intn(10001) + 50000 // 50000–60000 inclusive
	}

	// Define HTTP handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Prepare response
		response := map[string]string{"message": "OK Success"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(response)

		// Log request
		duration := time.Since(start)
		clientIP := r.RemoteAddr
		log.Printf("[%s] %s %s from %s (%v)",
			time.Now().Format("2006-01-02 15:04:05"),
			r.Method,
			r.URL.Path,
			clientIP,
			duration)
	})

	addr := fmt.Sprintf("%s:%d", *ip, *port)
	log.Printf("Server listening on http://%s/", addr)

	// Start HTTP server
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

