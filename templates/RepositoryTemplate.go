package templates

var repositoryTmpl = `// Code generated. DO NOT EDIT.

package entity

import (
	"github.com/jimvmathews/core/application"
)

// {{.Name}}Repository defines the interface for persisting {{.Name}} entities
type {{.Name}}Repository interface {
	Get(publicID string) (*{{.Name}}Entity, error)
	Save(e *{{.Name}}Entity) error
}

// {{.Name}}RepositoryImpl is the default implementation for the {{.Name}}Repository
type {{.Name}}RepositoryImpl struct {
	ds application.DataSource
}

// New{{.Name}}Repository returns a configured {{.Name}}Repository implementation
func New{{.Name}}Repository(datasource application.DataSource) {{.Name}}Repository {
	return &{{.Name}}RepositoryImpl{datasource}
}

// Get returns a {{.Name}}Entity that matches the publicID
func (r *{{.Name}}RepositoryImpl) Get(publicID string) (*{{.Name}}Entity, error) {
	var e {{.Name}}Entity = {{.Name}}Entity{}
	err := r.ds.Get(publicID, &e)
	return &e, err
}

// Save saves the {{.Name}}Entity
func (r *{{.Name}}RepositoryImpl) Save(e *{{.Name}}Entity) error {
	return r.ds.Save(e)
}
`
