# rules-codegen
Rules descriptor to support code geenration tool generates given rules descriptor to support file whihc contains conditions/actions/init methods used in descriptor, this will help user to code less.

Currently this tool accepts below arguments.
```sh
Usage of rules-codegen:
  -input string
        input rules descriptor file (default "rulesdescriptor.json")
  -output string
        path to store generated file (default ".")
```
## Setup
To install the tool, simply open a terminal and enter the below command
```sh
go get github.com/nareshkumarthota/rules-codegen/...
```

## Usage

```sh
cd $GOPATH/src/github.com/nareshkumarthota/rules-codegen/examples

rules-codegen -input rulesdescriptor.json -output flogoRulesApp
```
The resulting output is `support.go` file gets created under flogoRulesApp folder.


## Flogo Plugin Support
This tool can be integrated into [flogocli](https://github.com/project-flogo/cli).
```sh
# Install your plugin
$ flogo plugin install github.com/nareshkumarthota/rules-codegen/cmd

# Run your new plugin command to generate support file
$ flogo rulescodegen -i rulesdescriptor.json -o flogoRulesApp

# Remove your plugin
$ flogo plugin remove github.com/nareshkumarthota/swagger-flogo/cmd
```