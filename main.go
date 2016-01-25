package main

import (
	"fmt"
	"log"
	"net/http"

	"time"
)

func main() {
	time.Sleep(5 * time.Second)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from API")
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
