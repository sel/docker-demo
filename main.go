package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

// VERSION is the application version. Its value is injected through the build chain.
var VERSION string = "development"

func main() {
	http.HandleFunc("/", helloHandler)
	port := envInt("LISTEN_PORT", 8000)
	log.Printf("Listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

// envInt returns the integer value of an environment variable or otherwise returns a default.
func envInt(varName string, defaultVal int) int {
	if val, ok := os.LookupEnv(varName); ok {
		if intVal, err := strconv.Atoi(val); err == nil {
			return intVal
		}
	}
	return defaultVal
}

// helloHandler responds to HTTP requests with a greeting.
func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n\nVERSION: %s", VERSION)
}
