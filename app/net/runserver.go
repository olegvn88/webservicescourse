package main

import (
	"fmt"
	"net/http"
)

func runServer(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/",
		func(write http.ResponseWriter, request *http.Request) {
			fmt.Fprintln(write, "Addr:", addr, "URL:", request.URL.String())
		})

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	fmt.Println("starting server at ", addr)
	server.ListenAndServe()
}

func main() {
	go runServer(":8081")
	runServer(":8080")
}
