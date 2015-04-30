package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Args struct {
		TemplateFilename string
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
	TemplateFilename := opts.Args.TemplateFilename

	var variables = make(map[string]string)
	for _, variableWithValue := range os.Environ() {
		keyAndValue := strings.SplitN(variableWithValue, "=", 2)
		variables[keyAndValue[0]] = keyAndValue[1]
	}

	t := template.Must(template.New(TemplateFilename).ParseFiles(TemplateFilename))
	err = t.Execute(os.Stdout, variables)
	if err != nil {
		panic(err)
	}
}
