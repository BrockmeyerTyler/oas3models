package oas3models

import "encoding/json"

type PathItemDoc struct {
	Ref string
	Summary, Description string
	Get, Put, Post, Delete *OperationDoc
	Options, Head, Patch, Trace *OperationDoc
	Servers []*ServerDoc
	Parameters []*ParameterDoc
}
func (p *PathItemDoc) MarshalJSON() (_ []byte, err error) {
	x := make(JSON)
	if err = marshalObjIfNotNil(p.Get, "get", x); err != nil {return}
	if err = marshalObjIfNotNil(p.Put, "put", x); err != nil {return}
	if err = marshalObjIfNotNil(p.Post, "post", x); err != nil {return}
	if err = marshalObjIfNotNil(p.Delete, "delete", x); err != nil {return}
	if err = marshalObjIfNotNil(p.Options, "options", x); err != nil {return}
	if err = marshalObjIfNotNil(p.Head, "head", x); err != nil {return}
	if err = marshalObjIfNotNil(p.Patch, "patch", x); err != nil {return}
	if err = marshalObjIfNotNil(p.Trace, "trace", x); err != nil {return}
	if err = marshalObjIfNotNil(p.Servers, "servers", x); err != nil {return}
	if err = marshalObjIfNotNil(p.Parameters, "parameters", x); err != nil {return}
	marshalStrIfLen(p.Ref, "$ref", x)
	marshalStrIfLen(p.Summary, "summary", x)
	marshalStrIfLen(p.Description, "description", x)
	return json.Marshal(x)
}
