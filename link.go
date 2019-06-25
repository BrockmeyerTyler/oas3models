package oas3models

import "encoding/json"

type LinkDoc struct {
	OperationRef, OperationId, Description string
	Parameters map[string]json.Marshaler
	RequestBody json.Marshaler
	Server *ServerDoc
}
func (l *LinkDoc) MarshalJSON() (_ []byte, err error) {
	x := make(JSON)
	if err = marshalObjIfNotNil(l.Parameters, "parameters", x); err != nil {return}
	if err = marshalObjIfNotNil(l.RequestBody, "requestBody", x); err != nil {return}
	if err = marshalObjIfNotNil(l.Server, "server", x); err != nil {return}
	marshalStrIfLen(l.OperationRef, "operationRef", x)
	marshalStrIfLen(l.OperationId, "operationId", x)
	marshalStrIfLen(l.Description, "description", x)
	return json.Marshal(x)
}
