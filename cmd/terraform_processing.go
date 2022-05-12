package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

type Variable struct {
	Name        string         `hcl:",label"`
	Description string         `hcl:"description,optional"`
	Sensitive   bool           `hcl:"sensitive,optional"`
	Type        *hcl.Attribute `hcl:"type,optional"`
	Default     *hcl.Attribute `hcl:"default,optional"`
	Options     hcl.Body       `hcl:",remain"`
}

type Output struct {
	Name        string   `hcl:",label"`
	Description string   `hcl:"description,optional"`
	Sensitive   bool     `hcl:"sensitive,optional"`
	Value       string   `hcl:"value,optional"`
	Options     hcl.Body `hcl:",remain"`
}

type Resource struct {
	Type    string   `hcl:"type,label"`
	Name    string   `hcl:"name,label"`
	Options hcl.Body `hcl:",remain"`
}

type Data struct {
	Type    string   `hcl:"type,label"`
	Name    string   `hcl:"name,label"`
	Options hcl.Body `hcl:",remain"`
}

type Config struct {
	Outputs   []*Output   `hcl:"output,block"`
	Variables []*Variable `hcl:"variable,block"`
	Resources []*Resource `hcl:"resource,block"`
	Data      []*Data     `hcl:"data,block"`
}

func createDocs(hclPath string) map[string][]map[string]string {
	var variables, outputs, resources, data []map[string]string

	parsedConfig := make(map[string][]map[string]string)
	hclConfig := make(map[string][]byte)

	c := &Config{}

	// Iterate all Terraform files and safe the contents in the hclConfig map
	for _, file := range filesInDirectory(hclPath, ".tf") {
		fileContent, err := os.ReadFile(hclPath + "/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		hclConfig[file.Name()] = fileContent
	}

	// Iterate all file contents
	for k, v := range hclConfig {
		parsedConfig, diags := hclsyntax.ParseConfig(v, k, hcl.Pos{Line: 1, Column: 1})
		if diags.HasErrors() {
			log.Fatal(diags)
		}

		diags = gohcl.DecodeBody(parsedConfig.Body, nil, c)
		if diags.HasErrors() {
			log.Fatal(diags)
		}
	}

	for _, v := range c.Variables {
		var variableType string
		var variableDefault string

		if v.Type != nil {
			variableType = (v.Type.Expr).Variables()[0].RootName()
		}

		if v.Default != nil {
			variableDefault = (v.Default.Expr).Variables()[0].RootName()
		}

		variables = append(variables, map[string]string{"name": v.Name, "description": v.Description,
			"sensitive": strconv.FormatBool(v.Sensitive), "type": variableType, "default": variableDefault})
	}

	for _, v := range c.Outputs {
		outputs = append(outputs, map[string]string{"name": v.Name, "description": v.Description,
			"sensitive": strconv.FormatBool(v.Sensitive), "value": v.Value})
	}

	for _, v := range c.Resources {
		resources = append(resources, map[string]string{"type": v.Type, "name": v.Name})
	}

	for _, v := range c.Data {
		data = append(data, map[string]string{"type": v.Type, "name": v.Name})
	}

	parsedConfig["variables"], parsedConfig["outputs"], parsedConfig["resources"], parsedConfig["data"] = variables, outputs, resources, data

	fmt.Println(parsedConfig)

	return parsedConfig
}
