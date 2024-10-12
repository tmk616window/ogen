package main

import (
	"log"
	"net/http"
	"server/ogen"
)

func main() {
	h, err := ogen.NewServer(nil)
	if err != nil {
		panic(err)
	}

	// Start the server
	if err := http.ListenAndServe(":8080", h); err != nil {
		log.Fatal(err)
	}
}
