package oas3models

import "encoding/json"

type MediaTypeDoc struct {
	Schema SchemaDoc
	Example *ExampleDoc
	Examples []*ExampleDoc
	Encoding map[string]*EncodingDoc
}
func (m *MediaTypeDoc) MarshalJSON() (_ []byte, err error) {
	x := make(JSON)
	if err = marshalObjIfNotNil(m.Schema, "schema", x); err != nil {return}
	if err = marshalObjIfNotNil(m.Example, "example", x); err != nil {return}
	if err = marshalObjIfNotNil(m.Examples, "examples", x); err != nil {return}
	if err = marshalObjIfNotNil(m.Encoding, "encoding", x); err != nil {return}
	return json.Marshal(x)
}
