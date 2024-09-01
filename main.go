package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var name = "service-1"
	if appName := os.Getenv("app_name"); appName != "" {
		name = appName
	}
	fmt.Fprintf(w, "my app hallo di rubah dari git 20, %s!", name)
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server starting on port 4000...")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
