package model

import (
	"encoding/json"
	"log"
)

// EntityModel is the structure that defines all the entities of the application
type EntityModel struct {
	AppName  string
	Entities []Entity
}

// Entity defines the structure for generating the Entity and related source code
type Entity struct {
	AppName     string
	PackageName string
	Name        string
	DocComment  string
	Prefix      string
	Fields      []Field
}

// Field isdefines the structure for a field in the entity
type Field struct {
	Name       string
	DataType   string
	ProtoField string
}

// NewEntityModel loads the entity model from an array of bytes
func NewEntityModel(file []byte) EntityModel {
	var model EntityModel
	err := json.Unmarshal([]byte(file), &model)
	if err != nil {
		log.Fatalf("Error unmarshalling model. Error is %v", err)
	}
	log.Printf("Model has %v entities\n", len(model.Entities))

	return model
}
