package oas3models

import (
	"encoding/json"
	"fmt"
)

type ParameterDoc struct {
	Name, Description        	string
	In                       	InRequest
	Required, Deprecated 		bool
	AllowEmptyValue, Explode 	bool
	AllowReserved 			 	bool
	Schema                   	SchemaDoc
	Example                  	*ExampleDoc
	Examples                 	ExamplesDoc
	Style                    	StyleSetting
	Content                  	map[MimeType]*MediaTypeDoc
}
func (p *ParameterDoc) MarshalJSON() (_ []byte, err error) {
	if p.Name == "" || p.In == "" {
		return nil, fmt.Errorf("for parameters, 'name' and 'in' are required")
	}
	if p.Schema == nil && p.Content == nil || p.Schema != nil && p.Content != nil {
		return nil, fmt.Errorf("for parameters, 'schema' or 'content' are required, but not both")
	}
	x := make(JSON)
	if err = marshalObjIfNotNil(p.Schema, "schema", x); err != nil {return}
	if err = marshalObjIfNotNil(p.Example, "example", x); err != nil {return}
	if err = marshalObjIfNotNil(p.Examples, "examples", x); err != nil {return}
	if err = marshalObjIfNotNil(p.Content, "content", x); err != nil {return}
	marshalStrIfLen(p.Name, "name", x)
	marshalStrIfLen(string(p.In), "in", x)
	marshalStrIfLen(p.Description, "description", x)
	marshalStrIfLen(string(p.Style), "style", x)
	marshalIfTrue(p.Explode, "explode", x)
	marshalIfTrue(p.AllowReserved, "allowReserved", x)
	marshalIfTrue(p.Required, "required", x)
	marshalIfTrue(p.Deprecated, "deprecated", x)
	marshalIfTrue(p.AllowEmptyValue, "allowEmptyValue", x)
	return json.Marshal(x)
}
