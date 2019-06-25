package oas3models

import (
	"encoding/json"
	"fmt"
)

type ServerVariableDoc struct {
	Enum []string
	Default, Description string
}
func (s *ServerVariableDoc) MarshalJSON() (_ []byte, err error) {
	if s.Default == "" {
		return nil, fmt.Errorf("for server variables, 'default' is required")
	}
	x := make(JSON)
	if err = marshalObjIfNotNil(s.Enum, "enum", x); err != nil {return}
	marshalStrIfLen(s.Default, "default", x)
	marshalStrIfLen(s.Description, "description", x)
	return json.Marshal(x)
}
