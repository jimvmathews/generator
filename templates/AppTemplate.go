package templates

var appTmpl = `// Code generated. DO NOT EDIT.
//go:generate generator . App

package main

import (
	"log"

	this "github.com/jimvmathews/common/application"
)

func main() {
	log.Println("In main function of {{.AppName}}")
	this.App.Start()
}
`
