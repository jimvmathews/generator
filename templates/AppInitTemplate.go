package templates

var appInitTmpl = `// Code generated. DO NOT EDIT.
//go:generate generator . AppInit

package main

import (
	"log"

	{{if .Entities}}
	"github.com/jimvmathews/common/application"

	"{{.AppName}}/entity"
	"{{.AppName}}/service"
	{{end}}
)

func init() {
	log.Println("In init of App")
	autowire()
	automigrate()
}

// autowire wires up all the dependencies and registers all the services
func autowire() {
	{{if .Entities}}
	log.Println("Autowiring dependencies")
	ds, dsErr := application.App.GetDataSource()
	if dsErr != nil {
		log.Printf("Data source is not configured")
	}
	gs := application.App.GetGRPCServer()
	{{range $i, $e := .Entities}}
	if dsErr == nil {
		r{{$i}} := entity.New{{.Name}}Repository(ds)
		service.Register{{.Name}}ServiceServer(gs.GetServer(), service.New{{.Name}}Service(r{{$i}}))
	}
	{{end}}
	{{end}}
}

// automigrate migrates the data source to the latest schema version required by the application
func automigrate() {
	{{if .Entities}}
	log.Println("Automigrate start")
	ds, dsErr := application.App.GetDataSource()
	if dsErr == nil {
		c := ds.(application.Component)
		c.Start()
		defer c.Close()
	
		{{range $i, $e := .Entities}}
		e{{$i}} := entity.{{.Name}}Entity{}
		ds.AutoMigrate(&e{{$i}})
		{{end}}
	}
	log.Println("Automigrate end")
	{{end}}
}
`
