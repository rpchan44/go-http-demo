package http_demo

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
	// CLI flags for IP and port
	ip := flag.String("ip", "0.0.0.0", "IP address to bind to")
	port := flag.Int("port", 0, "Port to listen on (if 0, a random port between 50000–60000 will be used)")
	flag.Parse()

	// Seed RNG
	rand.Seed(time.Now().UnixNano())

	// If no port provided, generate a random one between 50000–60000
	if *port == 0 {
		*port = rand.Intn(10001) + 50000 // 50000–60000 inclusive
	}

	// Handler for "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{"message": "OK Success"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Println("Error encoding JSON:", err)
		}
	})

	addr := fmt.Sprintf("%s:%d", *ip, *port)
	log.Printf("Server listening on http://%s/", addr)

	// Start the HTTP server
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

