package main

import (
	"log"
	"net/http"

	"github.com/serg1732/ConverterDataService/pkg/converter"
)

func main() {

	http.HandleFunc("/metrics", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case "GET":
		converter.ConverterYamlToPrometheus(w, r)
	default:
		w.Write([]byte("400 Bad Request"))
	}
}
