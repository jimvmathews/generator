package templates

var serviceTmpl = `// Code generated. DO NOT EDIT.
//go:generate generator . Service

package service

import (
	context "context"
	"{{.AppName}}/entity"
)

// {{.Name}}Service is the service definition for the {{.Name}} entity
type {{.Name}}Service struct {
	repo entity.{{.Name}}Repository
}

// New{{.Name}}Service constructs the {{.Name}}Service
func New{{.Name}}Service(repo entity.{{.Name}}Repository) *{{.Name}}Service {
	return &{{.Name}}Service{repo}
}

// Create{{.Name}} is the service implementation to create a {{.Name}}Entity
func (svc *{{.Name}}Service) Create{{.Name}}(ctx context.Context, pbe *{{.Name}}) (*{{.Name}}Response, error) {
	// Transform {{.Name}} to {{.Name}}Entity
	e := entity.{{.Name}}Entity{}
	e.PublicID = pbe.PublicId
	{{range .Fields}} {{if .ProtoField}}
	e.{{.Name}} = pbe.{{.Name}} {{end}}
	{{end}}

	// Save the {{.Name}}
	err := svc.repo.Save(&e)
	if err != nil {
		return nil, err
	}
	// Set the publicId if it was a new entity
	if pbe.PublicId == "" {
		pbe.PublicId = e.PublicID
	}

	return &{{.Name}}Response{Created: true, Entity: pbe}, nil
}
`
