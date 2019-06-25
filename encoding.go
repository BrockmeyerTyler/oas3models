package oas3models

import "encoding/json"

type EncodingDoc struct {
	ContentType MimeType
	Headers map[string]*HeaderDoc
	Style StyleSetting
	Explode, AllowReserved bool
}
func (e *EncodingDoc) MarshalJSON() (_ []byte, err error) {
	x := make(JSON)
	if err = marshalObjIfNotNil(e.Headers, "headers", x); err != nil {return}
	marshalStrIfLen(string(e.ContentType), "contentType", x)
	marshalStrIfLen(string(e.Style), "style", x)
	marshalIfTrue(e.Explode, "explode", x)
	marshalIfTrue(e.AllowReserved, "allowReserved", x)
	return json.Marshal(x)
}
