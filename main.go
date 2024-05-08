package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	IP string `json:"ip"`
}

func getPublicIP(w http.ResponseWriter, r *http.Request) {
	ip := strings.Split(r.RemoteAddr, ":")[0]
	response := Response{IP: ip}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", getPublicIP)
	fmt.Printf("Server is running on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error:", err)
	}
}
