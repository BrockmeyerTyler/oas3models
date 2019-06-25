package oas3models

import "encoding/json"

type SecurityRequirementDoc struct {
	Name string
	Scopes []string
}
func (s *SecurityRequirementDoc) MarshalJSON() (_ []byte, err error) {
	x := map[string][]string {s.Name: s.Scopes}
	return json.Marshal(x)
}
