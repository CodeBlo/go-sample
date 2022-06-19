package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Human struct {
	Name string
	Age  int16
}

var humans []Human

func main() {

	http.HandleFunc("/get", get)
	http.HandleFunc("/post", post)
	http.ListenAndServe(":8090", nil)
}

func get(w http.ResponseWriter, req *http.Request) {

	encoder := json.NewEncoder(w)
	encoder.Encode(humans)
}

func post(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		var h Human
		json.NewDecoder(req.Body).Decode(&h)
		humans = append(humans, h)
		fmt.Fprintf(w, "Appended\n")
	} else {
		fmt.Fprintln(w, "Not supported")
	}

}
