package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", userHandleFunc)
	http.ListenAndServe(":8080", nil)
}
func userHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "webserver test")
}
