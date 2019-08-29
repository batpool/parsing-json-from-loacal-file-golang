package main

import (
	"net/http"
	"play/handler"
)

func main() {

	
	http.HandleFunc("/", handler.IndexHandler)
	http.ListenAndServe(":80", nil)
}
