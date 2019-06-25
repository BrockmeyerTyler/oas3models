package oas3models

import "encoding/json"

type ContactDoc struct {
	Name, Url, Email string
}
func (c *ContactDoc) MarshalJSON() (_ []byte, err error) {
	x := make(JSON)
	marshalStrIfLen(c.Name, "name", x)
	marshalStrIfLen(c.Url, "url", x)
	marshalStrIfLen(c.Email, "email", x)
	return json.Marshal(x)
}
