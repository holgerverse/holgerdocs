package main

import (
	"io/ioutil"
	"log"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
)

type MarkdownContent struct {
	Title        string
	Description  string
	ExampleUsage string
	Resources    []map[string]string
	Dependencies []map[string]string
	Variables    []map[string]string
	Outputs      []map[string]string
}

func parseMarkdown(markdownPath string) map[string]string {
	// Read in the content of the existing markdown file
	fileContent, err := ioutil.ReadFile(markdownPath)
	if err != nil {
		log.Fatal(err)
	}

	results := make(map[string]string)

	// Create custom Markdown parser
	extensions := parser.FencedCode | parser.Tables | parser.DefinitionLists | parser.CommonExtensions
	parser := parser.NewWithExtensions(extensions)

	// Parse Markdown
	temp := markdown.Parse(fileContent, parser)

	for _, child := range temp.AsContainer().Children {
		// Check if the child node is of type Heading
		if _, ok := child.(*ast.Heading); ok {
			/*
			  Copy the heading text and safe it to results if the heading level is 1
			  Also check if the next node is of type paragraph and safe it to results, it
			  is going to be used as description.
			*/
			if (child.(*ast.Heading)).Level == 1 {
				results["title"] = string(child.GetChildren()[0].AsLeaf().Literal)

				description := ast.GetNextNode(child)
				if _, ok := description.(*ast.Paragraph); ok {
					results["description"] = string(description.GetChildren()[0].AsLeaf().Literal)
				}
			}

			/*
			  Check if the heading is level 2 and says 'Example usage'. If so check if the next node is of
			  type *ast.CodeBlock and safe it to results. && string(child.GetChildren()[0].AsLeaf().Literal) == "Example usage"
			*/
			if (child.(*ast.Heading)).Level == 2 && string(child.GetChildren()[0].AsLeaf().Literal) == "Example Usage" {
				exampleUsageCodeBlock := ast.GetNextNode(child)
				if _, ok := exampleUsageCodeBlock.(*ast.CodeBlock); ok {
					results["example_usage"] = string(exampleUsageCodeBlock.AsLeaf().Literal)
				}
			}
		}
	}

	return results
}
