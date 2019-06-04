package app

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/nareshkumarthota/rules-codegen/transform"
)

// Generate generates support file to given descriptor
func Generate(config *transform.Config) {
	// Read the descriptor
	data, err := ioutil.ReadFile(config.FileName)
	if err != nil {
		log.Fatal("error occured in reading file: ", err)
	}

	var fileData map[string]interface{}

	json.Unmarshal(data, &fileData)

	config.Data = fileData

	transform.DescriptorToSupportFile(config)

}
