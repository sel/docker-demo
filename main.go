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
	logger := log.New(os.Stdout, "[STDOUT] ", log.LstdFlags)
	startTicker(logger, 3*time.Second)

	errLogger := log.New(os.Stderr, "[STDERR] ", log.LstdFlags)
	startTicker(errLogger, 7*time.Second)

	http.HandleFunc("/", versionHandler(logger))
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", envInt("LISTEN_PORT", 8000)),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	logger.Printf("Listening on %s", s.Addr)
	logger.Fatal(s.ListenAndServe())
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

// versionHandler responds to requests with a greeting and logs to a given logger.
func versionHandler(logger *log.Logger) http.HandlerFunc {
	msg := fmt.Sprintf("VERSION: %s", VERSION)
	return func(w http.ResponseWriter, req *http.Request) {
		logger.Print(msg)
		fmt.Fprintln(w, msg)
	}
}

// startTicker periodically writes to the given logger.
func startTicker(logger *log.Logger, d time.Duration) {
	ticker := time.NewTicker(d)
	go func() {
		for {
			<-ticker.C
			logger.Println()
		}
	}()
}
