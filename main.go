package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// VERSION is the application version. Its value is injected through the build chain.
var VERSION string = "development"

func main() {
	http.HandleFunc("/", helloHandler)
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", envInt("LISTEN_PORT", 8000)),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Printf("Listening on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
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
