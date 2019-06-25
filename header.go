package oas3models

import (
	"encoding/json"
	"fmt"
)

type HeaderDoc struct {
	Description                 string
	Required, Deprecated        bool
	AllowEmptyValue, Explode 	bool
	AllowReserved               bool
	Schema                      SchemaDoc
	Example                     *ExampleDoc
	Examples                    ExamplesDoc
	Style                       StyleSetting
	Content                     map[MimeType]*MediaTypeDoc
}
func (h *HeaderDoc) MarshalJSON() (_ []byte, err error) {
	if h.Schema == nil && h.Content == nil || h.Schema != nil && h.Content != nil {
		return nil, fmt.Errorf(
			"for headers, 'schema' or 'content' are required, but not both")
	}
	x := make(JSON)
	if err = marshalObjIfNotNil(h.Schema, "schema", x); err != nil {return}
	if err = marshalObjIfNotNil(h.Example, "example", x); err != nil {return}
	if err = marshalObjIfNotNil(h.Examples, "examples", x); err != nil {return}
	if err = marshalObjIfNotNil(h.Content, "content", x); err != nil {return}
	marshalStrIfLen(h.Description, "description", x)
	marshalStrIfLen(string(h.Style), "style", x)
	marshalIfTrue(h.Required, "required", x)
	marshalIfTrue(h.Deprecated, "deprecated", x)
	marshalIfTrue(h.AllowEmptyValue, "allowEmptyValue", x)
	marshalIfTrue(h.Explode, "explode", x)
	marshalIfTrue(h.AllowReserved, "allowReserved", x)
	return json.Marshal(x)
}
