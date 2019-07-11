package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Main page")
	w.Write([]byte("!!!"))
}

func main() {
	http.HandleFunc("/page",
		func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintln(writer, "Single page:", request.URL.String())
		})

	http.HandleFunc("/pages/",
		func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintln(writer, "Multiple pages:", request.URL.String())
		})

	http.HandleFunc("/", handler)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
