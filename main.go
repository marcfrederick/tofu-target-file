package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

var resourceSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{Type: "resource", LabelNames: []string{"type", "name"}},
	},
}

type Resource struct {
	Type string
	Name string
}

func (r Resource) String() string {
	return fmt.Sprintf("%s.%s", r.Type, r.Name)
}

func findResourcesInFile(filename string) ([]Resource, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	file, diags := hclsyntax.ParseConfig(content, filename, hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		return nil, fmt.Errorf("error parsing HCL: %w", diags)
	}

	body, _, diags := file.Body.PartialContent(resourceSchema)
	if diags.HasErrors() {
		return nil, fmt.Errorf("error reading HCL body: %w", diags)
	}

	var resources []Resource
	for _, block := range body.Blocks {
		resources = append(resources, Resource{Type: block.Labels[0], Name: block.Labels[1]})
	}
	return resources, nil
}

func main() {
	paths := os.Args[1:]
	if len(paths) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file_path> [<file_path> ...]\n", os.Args[0])
		os.Exit(1)
	}

	var resources []Resource
	hasError := false
	for _, path := range paths {
		pathResources, err := findResourcesInFile(path)
		if err != nil {
			log.Printf("error reading resources from %s: %v", path, err)
			hasError = true
			continue
		}
		resources = append(resources, pathResources...)
	}

	for _, resource := range resources {
		fmt.Println(resource)
	}

	if hasError {
		os.Exit(1)
	}
}
