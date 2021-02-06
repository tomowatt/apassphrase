package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"passphrase"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", passphrase.GetPassphrase)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
