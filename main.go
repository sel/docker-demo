package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	http.HandleFunc("/", hello)
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
