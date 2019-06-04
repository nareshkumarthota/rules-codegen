package cmd

import (
	"github.com/nareshkumarthota/rules-codegen/app"
	"github.com/nareshkumarthota/rules-codegen/transform"
	"github.com/project-flogo/cli/common" // Flogo CLI support code
	"github.com/spf13/cobra"
)

func init() {
	appgen.Flags().StringVarP(&input, "input", "i", "swagger.json", "path to input swagger file")
	appgen.Flags().StringVarP(&output, "output", "o", ".", "path to generated file")
	common.RegisterPlugin(appgen)
}

var input, output string
var appgen = &cobra.Command{
	Use:              "rulescodegen",
	Short:            "generates support file for rules app",
	Long:             "generates support file for flogo based rules application for given descriptor",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {
		config := &transform.Config{}

		config.FileName = input
		config.OutFilePath = output

		app.Generate(config)
	},
}
