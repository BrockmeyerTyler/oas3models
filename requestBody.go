package oas3models

import (
	"encoding/json"
	"fmt"
)

type RequestBodyDoc struct {
	Description string
	Required bool
	Content map[MimeType]*MediaTypeDoc
}
func (r *RequestBodyDoc) MarshalJSON() (_ []byte, err error) {
	if r.Content == nil {
		return nil, fmt.Errorf("for request bodies, 'content' is required")
	}
	x := make(JSON)
	if err = marshalObjIfNotNil(r.Content, "content", x); err != nil {return}
	marshalStrIfLen(r.Description, "description", x)
	marshalIfTrue(r.Required, "required", x)
	return json.Marshal(x)
}
