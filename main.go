package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/metrics", ConverterData.ConverterYamlToPrometheus)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
