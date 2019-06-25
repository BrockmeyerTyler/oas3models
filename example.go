package oas3models

import "encoding/json"

type ExampleDoc struct {
	Summary, Description, ExternalValue string
	Value json.Marshaler
}
func (e *ExampleDoc) MarshalJSON() (_ []byte, err error) {
	x := make(JSON)
	if err = marshalObjIfNotNil(e.Value, "value", x); err != nil {return}
	marshalStrIfLen(e.Summary, "summary", x)
	marshalStrIfLen(e.Description, "description", x)
	marshalStrIfLen(e.ExternalValue, "externalValue", x)
	return json.Marshal(x)
}
