package test

import (
	"encoding/json"
	"github.com/tjbrockmeyer/oas3models"
	"testing"
)

type JSON map[string]interface{}

func testJSON(t *testing.T, item interface{}, f func(JSON, error)) {
	var parsed interface{}
	itemAsJSON, err := json.Marshal(item)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(itemAsJSON))
	err = json.Unmarshal(itemAsJSON, &parsed)
	if err != nil {
		f(nil, err)
		return
	}
	x, ok := (parsed).(map[string]interface{})
	if !ok {
		t.Errorf("could not convert %v to map[string]interface{}", parsed)
		return
	}
	f(x, nil)
}

func TestComponents(t *testing.T) {
	components := &oasm.ComponentsDoc{}
	testJSON(t, components, func(x JSON, err error) {
		if err != nil {
			t.Error("components should be able to serialize with no values")
		}
		if len(x) != 0 {
			t.Error("map should have no keys")
		}
	})
	components.Schemas = map[string]interface{}{"schema": "abc"}
	testJSON(t, components, func(x JSON, err error) {
		if len(x) != 1 {
			t.Error("map should have exactly 1 key")
		}
		if _, ok := x["schemas"]; !ok {
			t.Error("map should have a key of 'schema'")
		}
	})
	components.Parameters = map[string]*oasm.ParameterDoc{"abc": {Name: "abc", In: "query", Schema: "abc"}}
	testJSON(t, components, func(x JSON, err error) {
		if len(x) != 2 {
			t.Error("map should have exactly 2 keys")
		}
		if _, ok := x["parameters"]; !ok {
			t.Error("map should have a key of 'parameters'")
		}
	})
}

func TestContact(t *testing.T) {
	contact := &oasm.ContactDoc{}
	testJSON(t, contact, func(x JSON, err error) {
		if len(x) != 0 {
			t.Error("map should have no keys")
		}
	})
	contact.Name = "Tyler Brockmeyer"
	testJSON(t, contact, func(x JSON, err error) {
		if len(x) != 1 {
			t.Error("map should have exactly 1 key")
		}
	})
}

func TestEncoding(t *testing.T) {
	encoding := &oasm.EncodingDoc{}
	testJSON(t, encoding, func(x JSON, err error) {
		if len(x) != 0 {
			t.Error("map should have no keys")
		}
	})
	encoding.Explode = true
	testJSON(t, encoding, func(x JSON, err error) {
		if len(x) != 1 {
			t.Error("map should have exactly 1 key")
		}
	})
}

func TestOpenApi(t *testing.T) {
	openapi := oasm.OpenAPIDoc{}
	testJSON(t, openapi, func(x JSON, err error) {
		_, ok1 := x["openapi"]
		_, ok2 := x["info"]
		_, ok3 := x["paths"]
		if !ok1 || !ok2 || !ok3 || len(x) != 3 {
			t.Error("map should have exactly three keys: 'openapi', 'info', and 'paths'")
		}
	})
	openapi.Paths = oasm.PathsDoc{}
	testJSON(t, openapi, func(x JSON, err error) {
		if len(x) != 3 {
			t.Error("map should have exactly 3 keys still")
		}
	})
	openapi.Tags = make([]*oasm.TagDoc, 1, 2)
	testJSON(t, openapi, func(x JSON, err error) {
		if len(x) != 4 {
			t.Error("map should have exactly 4 keys")
		}
	})
}

func TestParameter(t *testing.T) {
	parameter := oasm.ParameterDoc{}
	testJSON(t, parameter, func(x JSON, err error) {
		_, ok1 := x["name"]
		_, ok2 := x["in"]
		if !ok1 || !ok2 || len(x) != 2 {
			t.Error("map should have exactly two keys: 'name' and 'in'")
		}
	})
	parameter.Name = "param"
	testJSON(t, parameter, func(x JSON, err error) {
		if len(x) != 2 {
			t.Error("map should have exactly 2 keys still")
		}
	})
	parameter.Description = "The local server"
	testJSON(t, parameter, func(x JSON, err error) {
		if len(x) != 3 {
			t.Error("map should have exactly 3 keys")
		}
	})
}

func TestResponses(t *testing.T) {
	responses := &oasm.ResponsesDoc{}
	testJSON(t, responses, func(x JSON, err error) {
		if len(x) != 0 {
			t.Error("map should have no keys")
		}
	})
	responses.Default = &oasm.ResponseDoc{}
	testJSON(t, responses, func(x JSON, err error) {
		if len(x) != 1 {
			t.Error("map should have exactly 1 key")
		}
	})
	responses.Codes = map[int]*oasm.ResponseDoc{200: {}}
	testJSON(t, responses, func(x JSON, err error) {
		_, ok := x["200"]
		if !ok || len(x) != 2 {
			t.Error("map should have exactly 2 keys, where one is '200'")
		}
	})
}

func TestServer(t *testing.T) {
	server := oasm.ServerDoc{}
	testJSON(t, server, func(x JSON, err error) {
		if _, ok := x["url"]; !ok || len(x) != 1 {
			t.Error("map should have exactly one key: 'url'")
		}
	})
	server.Url = "localhost:5000"
	testJSON(t, server, func(x JSON, err error) {
		if len(x) != 1 {
			t.Error("map should have exactly 1 key")
		}
	})
	server.Description = "The local server"
	testJSON(t, server, func(x JSON, err error) {
		if len(x) != 2 {
			t.Error("map should have exactly 2 keys")
		}
	})
}
