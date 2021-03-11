package templates

var entityTmpl = `// Code generated. DO NOT EDIT.
//go:generate generator . Entity

package entity

import (
	a "github.com/jimvmathews/core/domain/valueobjects"
	"github.com/jimvmathews/core/util"
	"github.com/jinzhu/gorm"
)

// {{.Name}}Entity defines the {{.Name}} entity structure
type {{.Name}}Entity struct {
	a.PublicIdentity
	{{range .Fields}} {{.Name}} {{.DataType}}
	{{end}}
}

// Equals defines the equality for an entity
func (e *{{.Name}}Entity) Equals(a {{.Name}}Entity) bool {
	return e.PublicID == a.PublicID
}

// BeforeCreate defines the before create hook to set the generated Public ID for the {{.Name}} entity
func (e *{{.Name}}Entity) BeforeCreate(scope gorm.Scope) error {
	scope.SetColumn("PublicID", util.GenerateUID("{{.Prefix}}"))
	return nil
}
`
