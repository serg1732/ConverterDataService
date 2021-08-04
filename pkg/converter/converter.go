package converter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

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

	currencies := make(map[string]string)

	isValid := regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString

	for _, couple := range yamlData.Currencies {
		if !isValid(couple.Name) || !isValid(couple.Value) {
			log.Fatal("Error! YAML data not correct!")
		}
		currencies[couple.Name] = couple.Value
	}

	jsonData, errJson := json.Marshal(currencies)

	if errJson != nil {
		log.Fatal(errJson)
	}

	formatedData := fmt.Sprintf("currencies%s", string(jsonData))
	formatedData = strings.ReplaceAll(formatedData, ":", "=")
	formatedData = strings.ReplaceAll(formatedData, ",", ", ")
	w.Write([]byte(formatedData))
}
