package oas3models

import (
	"encoding/json"
	"fmt"
)

type ResponseDoc struct {
	Description string
	Headers map[string]*HeaderDoc
	Content map[MimeType]*MediaTypeDoc
	Links map[string]*LinkDoc
}
func (r *ResponseDoc) MarshalJSON() (_ []byte, err error) {
	if r.Description == "" {
		return nil, fmt.Errorf("for responses, 'description' is required")
	}
	x := make(JSON)
	if err = marshalObjIfNotNil(r.Headers, "headers", x); err != nil {return}
	if err = marshalObjIfNotNil(r.Content, "content", x); err != nil {return}
	if err = marshalObjIfNotNil(r.Links, "links", x); err != nil {return}
	marshalStrIfLen(r.Description, "description", x)
	return json.Marshal(x)
}
