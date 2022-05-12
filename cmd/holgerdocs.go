package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

var readmePath string

func holgerdocs(folderPath string, terraform bool) {
	readmePath = folderPath + "/README.md"

	switch {
	case terraform:
		createDocsForTerraform(folderPath)
	default:
		fmt.Println("No documentation available for this software.")
	}
}

/*
	Collect data from existing README file and Terraform config.
	Call the renderDocs function with specification for Terraform Templates.
*/
func createDocsForTerraform(folderPath string) {
	// Get the contents from the Terraform Configuration
	collectedConfig := createDocs(folderPath)
	// Get the existing content from the README file which is static
	existingContent := parseMarkdown(readmePath)

	// Create finished Markdown object which holds all data to be rendered
	markdownContent := MarkdownContent{
		Title:        existingContent["title"],
		Description:  existingContent["description"],
		ExampleUsage: existingContent["example_usage"],
		Resources:    collectedConfig["resources"],
		Dependencies: collectedConfig["data"],
		Variables:    collectedConfig["variables"],
		Outputs:      collectedConfig["outputs"]}

	renderDocs(folderPath, "templates/holgerdocs_terraform.tmpl", markdownContent)
}

func renderDocs(folderPath string, templatePath string, content MarkdownContent) {
	// Create the absolute path of the template file which is based on the software you want to create your docs from.
	templateFilePath, err := filepath.Abs(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	// Read the content of the template file which is based on the software you want to create your docs from.
	templateContent, err := ioutil.ReadFile(templateFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Create the template object from the template file
	tmpl, err := template.New("holgerdocs").Parse(string(templateContent))
	if err != nil {
		log.Fatal(err)
	}
	// Create the README file, do nothing if it already exists
	f, err := os.Create(readmePath)
	if err != nil {
		log.Fatal(err)
	}

	// Write the populated template to the README file
	err = tmpl.Execute(f, content)
	if err != nil {
		log.Fatal(err)
	}
}
