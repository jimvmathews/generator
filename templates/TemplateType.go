package templates

import (
	"fmt"
	"html/template"
	"log"
)

// TemplateType defines a supported template for generation
type TemplateType int

const (
	// EntityTemplate defines the Entity template
	EntityTemplate TemplateType = iota + 1
	// ServiceTemplate defines the template for the Service component
	ServiceTemplate
	// ProtoTemplate defines the template for the Protobuf files
	ProtoTemplate
	// AppInitTemplate defines the template for the applcation initializing code
	AppInitTemplate
	// All defines all templates
	All
)

func (t TemplateType) String() string {
	return [...]string{"Invalid", "Entity", "Service", "Proto", "AppInit", "All"}[t]
}

// GetSourceGenerators returns the set of templates based on the type
func (t TemplateType) GetSourceGenerators() ([]SourceGenerator, error) {
	var (
		generators []SourceGenerator
		err        error
	)
	switch t {
	case EntityTemplate:
		generators, err = getEntityGenerators()
	case ServiceTemplate:
		generators, err = getServiceGenerators()
	case ProtoTemplate:
		generators, err = getProtoGenerators()
	case AppInitTemplate:
		generators, err = getAppInitGenerators()
	}
	return generators, err
}

var templateLookup = map[string]TemplateType{
	"Entity":  EntityTemplate,
	"Service": ServiceTemplate,
	"Proto":   ProtoTemplate,
	"AppInit": AppInitTemplate,
}

// GetTemplateType returns the TemplateType given a string
func GetTemplateType(t string) (TemplateType, error) {
	result, found := templateLookup[t]
	if !found {
		return 0, fmt.Errorf("%s not a valid value", t)
	}
	return result, nil
}

// getEntityGenerators returns the template for generating the entity
func getEntityGenerators() ([]SourceGenerator, error) {
	var (
		generators = make([]SourceGenerator, 2)
		err        error
	)
	entityGenerator := SourceGeneratorImpl{packageName: "entity", fileExtension: ".go"}
	repoGenerator := SourceGeneratorImpl{packageName: "entity", fileExtension: ".go"}

	entityGenerator.Template, err = template.New("Entity").Parse(entityTmpl)
	if err != nil {
		log.Fatalf("Error parsing & creating the template. %v", err)
	}
	repoGenerator.Template, err = template.New("Repository").Parse(repositoryTmpl)
	if err != nil {
		log.Fatalf("Error parsing & creating the template. %v", err)
	}
	generators[0] = &entityGenerator
	generators[1] = &repoGenerator
	return generators, nil
}

func getServiceGenerators() ([]SourceGenerator, error) {
	var (
		generators = make([]SourceGenerator, 1)
		err        error
	)
	svcGenerator := SourceGeneratorImpl{packageName: "service", fileExtension: ".go"}

	svcGenerator.Template, err = template.New("Service").Parse(serviceTmpl)
	if err != nil {
		log.Fatalf("Error parsing & creating the template. %v", err)
	}
	generators[0] = &svcGenerator

	return generators, nil
}

func getProtoGenerators() ([]SourceGenerator, error) {
	var (
		generators = make([]SourceGenerator, 1)
		err        error
	)
	protoGenerator := SourceGeneratorImpl{packageName: "protos", fileExtension: ".proto"}

	protoGenerator.Template = template.New("Proto")
	funcMap := template.FuncMap{"inc": func(i int) int { return i + 2 }}
	protoGenerator.Template.Funcs(funcMap)
	_, err = protoGenerator.Template.Parse(protoTmpl)
	if err != nil {
		log.Fatalf("Error parsing & creating the template. %v", err)
	}
	generators[0] = &protoGenerator

	return generators, nil
}

func getAppInitGenerators() ([]SourceGenerator, error) {
	var (
		generators = make([]SourceGenerator, 1)
		err        error
	)
	appInitGenerator := AppInitGenerator{}

	appInitGenerator.Template = template.New("AppInit")
	funcMap := template.FuncMap{"inc": func(i int) int { return i + 1 }}
	appInitGenerator.Template.Funcs(funcMap)
	_, err = appInitGenerator.Template.Parse(appInitTmpl)
	if err != nil {
		log.Fatalf("Error parsing & creating the template. %v", err)
	}
	generators[0] = &appInitGenerator

	return generators, nil
}
