package templates

import (
	"bytes"
	"go/format"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/jimvmathews/generator/model"
)

// SourceGenerator is an interface that generates sources for an entity model and returns an array of file names
type SourceGenerator interface {
	SourceGenerate(m model.EntityModel) []string
}

// SourceGeneratorImpl defines the structure for the source template
type SourceGeneratorImpl struct {
	packageName   string
	fileExtension string
	Template      *template.Template
}

// SourceGenerate generates the source for a template
func (sg *SourceGeneratorImpl) SourceGenerate(m model.EntityModel) []string {
	var fNames = make([]string, len(m.Entities))
	for i, e := range m.Entities {
		packageDir := filepath.Join(e.AppName, sg.packageName)
		if packageDir == "" {
			log.Fatalln("Package directory name is empty")
		}
		_ = os.Mkdir(packageDir, 666)
		var (
			buf bytes.Buffer
			err error
		)
		err = sg.Template.Execute(&buf, e)
		if err != nil {
			log.Fatalf("Error generating the content using the template & model. %v", err)
		}
		log.Println(buf.String())
		data := buf.Bytes()
		if sg.fileExtension == ".go" {
			data, err = format.Source(buf.Bytes())
			if err != nil {
				log.Fatalf("Error formatting the source: " + err.Error())
			}
		}
		fName := filepath.Join(packageDir, e.Name+sg.Template.Name()+"_gen"+sg.fileExtension)
		err = ioutil.WriteFile(fName, data, 077)
		if err != nil {
			log.Fatalf("Error writing the source file for %v. Error: %v", e.Name, err)
		}
		fNames[i] = fName
	}
	return fNames
}
