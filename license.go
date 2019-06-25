package oas3models

import (
	"encoding/json"
	"fmt"
)

type LicenseDoc struct {
	Name, Url string
}
func (l *LicenseDoc) MarshalJSON() (_ []byte, err error) {
	if l.Name == "" {
		return nil, fmt.Errorf("for licenses, 'name' is required")
	}
	x := make(JSON)
	marshalStrIfLen(l.Name, "name", x)
	marshalStrIfLen(l.Url, "url", x)
	return json.Marshal(x)
}
