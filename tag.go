package oas3models

import "encoding/json"

type TagDoc struct {
	Name, Description string
	ExternalDocs *ExternalDocumentationDoc
}
func (t *TagDoc) MarshalJSON() (_ []byte, err error) {
	x := make(JSON)
	if err = marshalObjIfNotNil(t.ExternalDocs, "externalDocs", x); err != nil {return}
	marshalStrIfLen(t.Name, "name", x)
	marshalStrIfLen(t.Description, "description", x)
	return json.Marshal(x)
}
