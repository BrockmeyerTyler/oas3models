package oas3models

import "encoding/json"

type ResponsesDoc struct {
	Default *ResponseDoc
	Responses map[int]*ResponseDoc
}
func (r *ResponsesDoc) MarshalJSON() (_ []byte, err error) {
	x := make(JSON)
	if err = marshalObjIfNotNil(r.Default, "default", x); err != nil {return}
	if r.Responses != nil {
		for key, value := range r.Responses {
			x[string(key)], _ = json.Marshal(value)
		}
	}
	return json.Marshal(x)
}
