package oas3models

import (
	"encoding/json"
	"fmt"
)

type InfoDoc struct {
	Title, Description string
	TermsOfService, Version string
	Contact *ContactDoc
	License *LicenseDoc
}
func (i *InfoDoc) MarshalJSON() (_ []byte, err error) {
	if i.Title == "" || i.Version == "" {
		return nil, fmt.Errorf("for info, 'title' and 'version' are required")
	}
	x := make(JSON)
	if err = marshalObjIfNotNil(i.Contact, "contact", x); err != nil {return}
	if err = marshalObjIfNotNil(i.License, "license", x); err != nil {return}
	marshalStrIfLen(i.Title, "title", x)
	marshalStrIfLen(i.Description, "description", x)
	marshalStrIfLen(i.TermsOfService, "termsOfService", x)
	marshalStrIfLen(i.Version, "version", x)
	return json.Marshal(x)
}

