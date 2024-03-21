package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
