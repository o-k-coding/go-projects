package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type Kata struct {
	PackageName string
	TestPackageName string
	FunctionName string
}


func main() {
	packageNamePtr := flag.String("package", "example", "a string representing the package name for the kata")
	flag.Parse()
	packageName := *packageNamePtr
	testPackageName := packageName + "_test"
	packageNameParts := strings.Split(packageName, "_")
	var partsBuilder strings.Builder
	for _, part := range packageNameParts {
		partsBuilder.WriteString(strings.Title(part))
	}
	functionName := partsBuilder.String()
	kata := Kata{
		PackageName: packageName,
		TestPackageName: testPackageName,
		FunctionName: functionName,
	}
	t_root := template.Must(template.New("package.tmpl").ParseGlob("templates/*.tmpl"))

	for _, t_node := range t_root.Templates() {
		name := t_node.Name()
		new_name := strings.ReplaceAll(name, "package", packageName)
		new_name = strings.ReplaceAll(new_name, ".tmpl", "")

		file, err := os.Create("output/" + new_name)
		if err != nil {
				fmt.Println("Error creating file output/" + new_name, err)
				continue
		}

		err = t_node.Execute(file, kata)
		if err != nil {
			panic(err)
		}
	}

}
