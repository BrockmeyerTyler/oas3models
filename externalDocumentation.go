package oas3models

import (
	"encoding/json"
	"fmt"
)

type ExternalDocumentationDoc struct {
	Description, Url string
}
func (e *ExternalDocumentationDoc) MarshalJSON() (_ []byte, err error) {
	if e.Url == "" {
		return nil, fmt.Errorf("for external documentations, 'url' is required")
	}
	x := make(JSON)
	marshalStrIfLen(e.Description, "description", x)
	marshalStrIfLen(e.Url, "url", x)
	return json.Marshal(x)
}
