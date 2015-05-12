package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Args struct {
		TemplatePath string
	} `positional-args:"yes" required:"yes"`
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		if err.(*flags.Error).Type == flags.ErrHelp {
			return
		}
		fmt.Fprintln(os.Stderr, err)
		return
	}
	templatePath := opts.Args.TemplatePath
	templateFilename := filepath.Base(templatePath)

	var variables = make(map[string]string)
	for _, variableWithValue := range os.Environ() {
		keyAndValue := strings.SplitN(variableWithValue, "=", 2)
		variables[keyAndValue[0]] = keyAndValue[1]
	}

	t := template.Must(template.New(templateFilename).ParseFiles(templatePath))
	err = t.Execute(os.Stdout, variables)
	if err != nil {
		panic(err)
	}
}
