package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/data", data)

	fmt.Println("Server start")
	http.ListenAndServe(":1111", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./index.html")
}

func data(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Method: %s\n", req.Method)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if req.Method == "OPTIONS" {
		return
	}

	body, _ := ioutil.ReadAll(req.Body)
	w.Write(body)
}
