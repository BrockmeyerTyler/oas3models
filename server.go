package oas3models

import (
	"encoding/json"
	"fmt"
)

type ServerDoc struct {
	Url, Description string
	Variables map[string]*ServerVariableDoc
}
func (s *ServerDoc) MarshalJSON() (_ []byte, err error) {
	if s.Url == "" {
		return nil, fmt.Errorf("for servers, 'url' is required")
	}
	x := make(JSON)
	if err = marshalObjIfNotNil(s.Variables, "variables", x); err != nil {return}
	marshalStrIfLen(s.Url, "url", x)
	marshalStrIfLen(s.Description, "description", x)
	return json.Marshal(x)
}
