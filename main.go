package main

import (
	"flag"

	"github.com/nareshkumarthota/rules-codegen/app"
	"github.com/nareshkumarthota/rules-codegen/transform"
)

func main() {
	input := flag.String("input", "ruledescriptor.json", "input rules descriptor file")
	conversionType := flag.String("type", "flogodescriptor", "conversion type like flogoapiapp or flogodescriptor")
	output := flag.String("output", ".", "path to store generated file")

	flag.Parse()

	config := &transform.Config{}

	config.FileName = *input
	config.ConversionType = *conversionType
	config.OutFilePath = *output

	app.Transform(config)
}
