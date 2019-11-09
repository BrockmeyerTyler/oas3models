package oasm

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// A container for the expected responses of an operation. The container maps a HTTP response code
// to the expected response.
//
// The documentation is not necessarily expected to cover all possible HTTP response codes because they may not be
// known in advance. However, documentation is expected to cover a successful operation response and any known errors.
//
// The default MAY be used as a default response object for all HTTP codes
// that are not covered individually by the specification.
//
// The Responses Object MUST contain at least one response code,
// and it SHOULD be the response for a successful operation call.
type Responses struct {

	// The documentation of responses other than the ones declared for specific HTTP response codes. Use this field
	// to cover undeclared responses. A Reference Object can link to a response that the OpenAPI Object's
	// components/responses section defines.
	Default *Response

	// Any HTTP status code can be used as the property name, but only one property per code, to describe the expected
	// response for that HTTP status code. A Reference Object can link to a response that is defined in the OpenAPI
	// Object's components/responses section. This field MUST be enclosed in quotation marks (for example, "200") for
	// compatibility between JSON and YAML. To define a range of response codes, this field MAY contain the uppercase
	// wildcard character X. For example, 2XX represents all response codes between [200-299]. The following range
	// definitions are allowed: 1XX, 2XX, 3XX, 4XX, and 5XX. If a response range is defined using an explicit code,
	// the explicit code definition takes precedence over the range definition for that code.
	Codes map[int]Response
}

func (r *Responses) MarshalJSON() ([]byte, error) {
	x := make(JSON)
	var err error
	if r.Default != nil {
		x["default"], err = json.Marshal(r.Default)
		if err != nil {
			return nil, err
		}
	}
	if r.Codes != nil {
		for key, value := range r.Codes {
			x[strconv.Itoa(key)], err = json.Marshal(value)
			if err != nil {
				return nil, err
			}
		}
	}
	return json.Marshal(x)
}

func (r *Responses) UnmarshalJSON(b []byte) error {
	r.Codes = make(map[int]Response)
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		if k == "default" {
			mv, _ := json.Marshal(v)
			err = json.Unmarshal(mv, &r.Default)
			if err != nil {
				return err
			}
		} else {
			i, err := strconv.ParseInt(k, 10, 32)
			if err != nil {
				return fmt.Errorf("valid values for response keys are 'default' and integer numbers")
			}
			var response Response
			mv, _ := json.Marshal(v)
			err = json.Unmarshal(mv, &response)
			if err != nil {
				return err
			}
			r.Codes[int(i)] = response
		}
	}
	return nil
}
