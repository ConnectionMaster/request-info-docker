package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Request Host: ", r.Host)
	fmt.Println("Request URL: ", r.URL)
	fmt.Println("Context path: ", r.URL.Path[1:])
	fmt.Println("Client IP (RemoteAddr): ", r.RemoteAddr)
	fmt.Println("Request TLS: ", r.TLS)

	fmt.Println("Request Headers:")
	var keys []string
	for k := range r.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println("\t", k, ": ", r.Header[k])
	}

	fmt.Println("Form parameters:")
	r.ParseForm()
	keys = []string{}
	for k := range r.Form {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println("\t", k, ":", r.Form[k])
	}
	fmt.Println("")

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, nil)
}
