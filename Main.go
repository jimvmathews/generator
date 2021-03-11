package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/jimvmathews/generator/model"
	"github.com/jimvmathews/generator/templates"
)

func generate(t templates.TemplateType, entityModel model.EntityModel) {
	generators, err := t.GetSourceGenerators()
	if err != nil {
		log.Fatalf("Error loading the template")
	}
	log.Printf("Template is %v\n", generators)

	for _, generator := range generators {
		fNames := generator.SourceGenerate(entityModel)
		log.Println("Generated the following files:")
		log.Println(fNames)
	}
}

func main() {
	args := os.Args

	var appNameArg, templateArg string
	if len(args) != 3 {
		log.Fatalln("Incorrect number of argument passed.")
	} else {
		appNameArg = args[1]
		templateArg = args[2]
	}

	templateType, err := templates.GetTemplateType(templateArg)
	if err != nil {
		log.Fatalf("Invalid argument for template type.")
	}
	log.Printf("Template type is %v %T\n", templateType, templateType)

	file, err := ioutil.ReadFile(filepath.Join(appNameArg, "entities.json"))
	if err != nil {
		log.Fatalf("Error loading the entities.json. Error is %v", err)
	}

	entityModel := model.NewEntityModel(file)

	generate(templateType, entityModel)
}
