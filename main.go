package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	myServer := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Println("Listening on port 8080...")
	err := myServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
