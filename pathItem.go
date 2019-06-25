package oas3models

import "encoding/json"

type PathItemDoc struct {
	Ref                  string
	Summary, Description string
	Methods              map[HTTPVerb]*OperationDoc
	Servers              []*ServerDoc
	Parameters           []*ParameterDoc
}

func (p *PathItemDoc) MarshalJSON() (_ []byte, err error) {
	x := make(JSON)
	for key, value := range p.Methods {
		if err = marshalObj(value, string(key), x); err != nil {
			return
		}
	}
	if err = marshalObjIfNotNil(p.Servers, "servers", x); err != nil {
		return
	}
	if err = marshalObjIfNotNil(p.Parameters, "parameters", x); err != nil {
		return
	}
	marshalStrIfLen(p.Ref, "$ref", x)
	marshalStrIfLen(p.Summary, "summary", x)
	marshalStrIfLen(p.Description, "description", x)
	return json.Marshal(x)
}
