package main

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	devices = make(map[string]string) // Store device ID and IP address
	mu      sync.Mutex                // To handle concurrent access to devices map
)

// Handler to register devices
func registerDevice(w http.ResponseWriter, r *http.Request) {
	deviceID := r.URL.Query().Get("id")
	ip := r.RemoteAddr

	mu.Lock()
	devices[deviceID] = ip
	mu.Unlock()

	fmt.Fprintf(w, "Registered device %s with IP %s\n", deviceID, ip)
}

// Handler to list registered devices
func getDevices(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	for id, ip := range devices {
		fmt.Fprintf(w, "Device %s: %s\n", id, ip)
	}
}

func main() {
	http.HandleFunc("/register", registerDevice)
	http.HandleFunc("/devices", getDevices)

	fmt.Println("Signaling server running on :8080")
	http.ListenAndServe(":8080", nil)
}
