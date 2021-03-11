package templates

import (
	"bytes"
	"go/format"
	"html/template"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/jimvmathews/generator/model"
)

// AppInitGenerator generates the source for autowiring the dependencies & auto migrating the data source of the application
type AppInitGenerator struct {
	Template *template.Template
}

// SourceGenerate generates the source for the application initialization
func (ag *AppInitGenerator) SourceGenerate(m model.EntityModel) []string {
	var fNames = make([]string, 1)
	packageDir := "."
	if len(m.Entities) > 0 {
		packageDir = m.Entities[0].AppName
	}
	var (
		buf bytes.Buffer
		err error
	)
	err = ag.Template.Execute(&buf, m)
	if err != nil {
		log.Fatalf("Error generating the content using the template & model. %v", err)
	}
	log.Println(buf.String())
	data := buf.Bytes()
	data, err = format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("Error formatting the source: " + err.Error())
	}
	fName := filepath.Join(packageDir, "AppInit_gen.go")
	err = ioutil.WriteFile(fName, data, 666)
	if err != nil {
		log.Fatalf("Error writing the app init source file. Error: %v", err)
	}
	fNames[0] = fName

	return fNames
}
