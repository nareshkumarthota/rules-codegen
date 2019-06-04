package util

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// ExecuteTemplate creates required file from given inputs
func ExecuteTemplate(conversionType, outFilePath string, data interface{}) (bool, error) {

	var t *template.Template

	var fileName string

	fileName = "support.go"
	t = template.Must(template.New("top").Parse(""))

	buf := &bytes.Buffer{}
	if err := t.Execute(buf, data); err != nil {
		log.Fatal("error while rendering template from data: ", err)
	}
	s := buf.String()

	// check for outfilepath if not exist create it
	if strings.Compare(outFilePath, ".") != 0 {
		_, err := os.Stat(outFilePath)
		if err != nil {
			fErr := os.MkdirAll(outFilePath, 0777)
			if fErr != nil {
				log.Fatal("unable to create out folder: ", outFilePath, fErr)
			}
		}
	}

	createFileWithContent(filepath.Join(outFilePath, fileName), s)

	return true, nil
}

func createFileWithContent(filename, content string) error {

	// Create a file on disk
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("error while creating file: %s", err.Error())
		return fmt.Errorf("error while creating file: %s", err.Error())
	}
	defer file.Close()

	// Open the file to write
	file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Printf("error while opening file: %s", err.Error())
		return fmt.Errorf("error while opening file: %s", err.Error())
	}

	// Write the doc to disk
	_, err = file.Write([]byte(content))
	if err != nil {
		log.Printf("error while writing doc to disk: %s", err.Error())
		return fmt.Errorf("error while writing doc to disk: %s", err.Error())
	}

	return nil
}
