package converter

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"gopkg.in/yaml.v3"
)

type YamlData struct {
	Currencies []struct {
		Name  string
		Value string
	}
}

func ConverterYamlToPrometheus(w http.ResponseWriter, r *http.Request) {

	ymlFile, errReadFile := ioutil.ReadFile("data.yaml")
	if errReadFile != nil {
		log.Fatal(errReadFile)
	}

	var yamlData YamlData
	errReadYaml := yaml.Unmarshal(ymlFile, &yamlData)

	if errReadFile != nil {
		log.Fatal(errReadYaml)
	}

	isValid := regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString

	var labels string
	if len(yamlData.Currencies) > 0 {
		labels = fmt.Sprintf("%s=\"%s\"", yamlData.Currencies[0].Name, yamlData.Currencies[0].Value)
	}

	for i := 1; i < len(yamlData.Currencies); i++ {
		couple := yamlData.Currencies[i]
		if !isValid(couple.Name) || !isValid(couple.Value) {
			log.Fatal("Error! YAML data not correct!")
		}
		labels = fmt.Sprintf("%s, %s=\"%s\"", labels, couple.Name, couple.Value)
	}

	formatedData := fmt.Sprintf("currencies{%s}", labels)
	w.Write([]byte(formatedData))
}
