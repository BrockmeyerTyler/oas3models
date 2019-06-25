package oas3models

import "encoding/json"

type ComponentsDoc struct {
	Schemas map[string]SchemaDoc
	Responses map[string]*ResponseDoc
	Parameters map[string]*ParameterDoc
	Examples map[string]*ExampleDoc
	RequestBodies map[string]*RequestBodyDoc
	Headers map[string]*HeaderDoc
	SecuritySchemes map[string]*SecuritySchemeDoc
	Links map[string]*LinkDoc
	Callbacks map[string]CallbackDoc
}
func (c *ComponentsDoc) MarshalJSON() (_ []byte, err error) {
	x := make(JSON)
	if err = marshalObjIfNotNil(c.Schemas, "schemas", x); err != nil {return}
	if err = marshalObjIfNotNil(c.Responses, "responses", x); err != nil {return}
	if err = marshalObjIfNotNil(c.Parameters, "parameters", x); err != nil {return}
	if err = marshalObjIfNotNil(c.Examples, "examples", x); err != nil {return}
	if err = marshalObjIfNotNil(c.RequestBodies, "requestBodies", x); err != nil {return}
	if err = marshalObjIfNotNil(c.Headers, "headers", x); err != nil {return}
	if err = marshalObjIfNotNil(c.SecuritySchemes, "securitySchemes", x); err != nil {return}
	if err = marshalObjIfNotNil(c.Links, "links", x); err != nil {return}
	if err = marshalObjIfNotNil(c.Callbacks, "callbacks", x); err != nil {return}
	return json.Marshal(x)
}
