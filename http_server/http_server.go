package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

const vk string = "VERSION"

func main() {
	os.Setenv(vk, "1")
	fmt.Println("Starting http server...")

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Item 1 - Transfer the request header to the response header.
	// Each request header is re-formatted from a string slice to a string.
	for k, vs := range r.Header {
		for _, v := range vs {
			w.Header().Set(k, v)
		}
	}

	// Item 2 - Put the VERSION env var into the resp header
	w.Header().Set(vk, os.Getenv(vk))

	// Item 3 - Log the client IP and the response code
	sc := http.StatusOK
	w.WriteHeader(sc)
	fmt.Println("Client Ip: ", r.RemoteAddr)
	fmt.Println("Status code: ", sc)

}

func healthz(w http.ResponseWriter, r *http.Request) {
	// Item 4 - Return 200 when routing to the healthz handler
	w.WriteHeader(http.StatusOK)
}
