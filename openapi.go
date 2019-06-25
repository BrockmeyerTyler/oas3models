package oas3models

import (
	"encoding/json"
	"fmt"
)

type OpenAPIDoc struct {
	OpenApi string
	Info *InfoDoc
	Servers []*ServerDoc
	Paths PathsDoc
	Components *ComponentsDoc
	Security []*SecurityRequirementDoc
	Tags []*TagDoc
	ExternalDocs *ExternalDocumentationDoc
}
func (o *OpenAPIDoc) MarshalJSON() (_ []byte, err error) {
	if o.OpenApi == "" || o.Info == nil || o.Paths == nil {
		return nil, fmt.Errorf("for open apis, 'openapi' is required")
	}
	x := make(JSON)
	if err = marshalObj(o.Info, "info", x); err != nil {return}
	if err = marshalObj(o.Paths, "paths", x); err != nil {return}
	if err = marshalObjIfNotNil(o.Servers, "servers", x); err != nil {return}
	if err = marshalObjIfNotNil(o.Components, "components", x); err != nil {return}
	if err = marshalObjIfNotNil(o.Security, "security", x); err != nil {return}
	if err = marshalObjIfNotNil(o.Tags, "tags", x); err != nil {return}
	if err = marshalObjIfNotNil(o.ExternalDocs, "externalDocs", x); err != nil {return}
	marshalStrIfLen(o.OpenApi, "openapi", x)
	return json.Marshal(x)
}
