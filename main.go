package main

import (
	"fmt"
	"html/template"
	"os"
)

type templateArguments struct {
	Title string
}

func main() {

	templ, err := template.New("layout").Parse(layout)
	if err != nil {
		panic(err)
	}

	_, err = templ.New("content").Parse(content)
	if err != nil {
		panic(err)
	}

	args := templateArguments{
		Title: "test",
	}

	fmt.Printf("Template names: %s\n", templ.DefinedTemplates())
	err = templ.Execute(os.Stdout, args)
	if err != nil {
		panic(err)
	}

}
