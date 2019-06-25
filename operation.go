package oas3models

import (
	"encoding/json"
	"fmt"
)

type OperationDoc struct {
	Tags []string
	Summary, Description, OperationId string
	ExternalDocs *ExternalDocumentationDoc
	Parameters []*ParameterDoc
	RequestBody *RequestBodyDoc
	Responses map[int]*ResponseDoc
	Callbacks json.Marshaler
	Deprecated bool
	Security []*SecurityRequirementDoc
	Server *ServerDoc
}
func (o *OperationDoc) MarshalJSON() (_ []byte, err error) {
	if o.Responses == nil {
		return nil, fmt.Errorf("for operations, 'responses' is required")
	}
	x := make(JSON)
	if err = marshalObj(o.Responses, "responses", x); err != nil {return}
	if err = marshalObjIfNotNil(o.Tags, "tags", x); err != nil {return}
	if err = marshalObjIfNotNil(o.ExternalDocs, "externalDocs", x); err != nil {return}
	if err = marshalObjIfNotNil(o.Parameters, "parameters", x); err != nil {return}
	if err = marshalObjIfNotNil(o.RequestBody, "requestBody", x); err != nil {return}
	if err = marshalObjIfNotNil(o.Callbacks, "callbacks", x); err != nil {return}
	if err = marshalObjIfNotNil(o.Security, "security", x); err != nil {return}
	if err = marshalObjIfNotNil(o.Server, "server", x); err != nil {return}
	marshalStrIfLen(o.Summary, "summary", x)
	marshalStrIfLen(o.Description, "description", x)
	marshalStrIfLen(o.OperationId, "operationId", x)
	marshalIfTrue(o.Deprecated, "deprecated", x)
	return json.Marshal(x)
}
