package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"play/utility"
)

// IndexHandler data
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "GO BAT 🦇🦇🦇🦇🦇🦇🦇")
	w.Header().Set("Content-Type", "application/json")
	mm, _ := json.Marshal(utility.ReadJson("database/data.json"))
	fmt.Fprint(w, string(mm))
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":80", nil)
}
