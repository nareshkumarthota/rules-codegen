package transform

import (
	"github.com/nareshkumarthota/rules-codegen/util"
)

// DescriptorToSupportFile transforms given swagger to flogo api application
func DescriptorToSupportFile(config *Config) error {
	util.ExecuteTemplate(config.ConversionType, config.OutFilePath, appData(config))
	return nil
}

func appData(config *Config) string {
	return ""
}
