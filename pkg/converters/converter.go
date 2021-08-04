package ConverterData

import (
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v3"
)

func ConverterYamlToPrometheus(w http.ResponseWriter, r *http.Request) {
	ymlFile, errReadFile := ioutil.ReadFile("data.yaml")
	if errReadFile != nil {
		log.Fatal(errReadFile)
	}
	data := make(map[interface{}]interface{})
	errReadYaml := yaml.Unmarshal(ymlFile, &data)
	if errReadFile != nil {
		log.Fatal(errReadYaml)
	}
}
