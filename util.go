package oas3models

import "encoding/json"


type JSON map[string][]byte
type PathsDoc map[string]*PathItemDoc
type CallbackDoc map[string]*PathItemDoc
type ExamplesDoc map[string]*ExampleDoc
type SchemaDoc interface {}
type OAuthFlows map[OAuthFlowType]*OAuthFlowDoc


func marshalObj(value interface{}, property string, object JSON) (err error) {
	object[property], err = json.Marshal(value)
	return
}

func marshalObjIfNotNil(value interface{}, property string, object JSON) (err error) {
	if value != nil {
		object[property], err = json.Marshal(value)
	}
	return
}

func marshalIfTrue(value bool, property string, object JSON) {
	if value {
		object[property], _ = json.Marshal(value)
	}
}

func marshalStrIfLen(value string, property string, object JSON) {
	if value != "" {
		object[property] = []byte(value)
	}
}