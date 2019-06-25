package oas3models

import (
	"encoding/json"
	"fmt"
)

type SecuritySchemeDoc struct {
	Type SecurityType
	Description string
	Name string  // apiKey
	In SecurityInRequest // apiKey
	Scheme HTTPAuthScheme // http
	BearerFormat string // http (bearer)
	Flows OAuthFlows // oauth2
	OpenIdConnectUrl string // openIdConnect
}
func (s *SecuritySchemeDoc) MarshalJSON() (_ []byte, err error) {
	if s.Type == "" {
		return nil, fmt.Errorf("for security schemes, 'type' is required")
	}
	x := make(JSON)
	marshalStrIfLen(string(s.Type), "type", x)
	marshalStrIfLen(s.Description, "description", x)
	if s.Type == SecurityApiKey {
		if s.Name == "" || s.In == "" {
			return nil, fmt.Errorf("for `%v` security schemes, 'name' and 'in' are required", SecurityApiKey)
		}
		x["name"] = []byte(s.Name)
		x["in"] = []byte(s.In)
	} else if s.Type == SecurityHttp {
		if s.Scheme == "" {
			return nil, fmt.Errorf("for `%v` security schemes, 'scheme' is required", SecurityHttp)
		}
		x["scheme"] = []byte(s.Scheme)
		if s.Scheme == AuthBearer {
			if s.BearerFormat == "" {
				return nil, fmt.Errorf(
					"for `%v` security schemes with a 'scheme' of `%v`, 'bearerFormat' is required",
					SecurityHttp, AuthBearer)
			}
			x["bearerFormat"] = []byte(s.BearerFormat)
		}
	} else if s.Type == SecurityOauth2 {
		if s.Flows == nil {
			return nil, fmt.Errorf("for `%v` security schemes, 'flows' is required", SecurityOauth2)
		}
		for key, value := range s.Flows {
			if value.AuthorizationUrl == "" && (key == OAuthImplicit || key == OAuthAuthorizationCode) {
				return nil, fmt.Errorf(
					"for flows of `%v` security schemes, 'authorizationUrl' is required", SecurityOauth2)
			} else if value.TokenUrl == "" &&
				(key == OAuthPassword || key == OAuthClientCredentials || key == OAuthAuthorizationCode) {
				return nil, fmt.Errorf(
					"for flows of `%v` security schemes, 'tokenUrl' is required", SecurityOauth2)
			}
		}
		var err error
		x["flows"], err = json.Marshal(s.Flows)
		if err != nil {return nil, err}
	} else if s.Type == SecurityOpenIdConnect {
		if s.OpenIdConnectUrl == "" {
			return nil, fmt.Errorf("for `%v` security schemes, 'openIdConnectUrl is required", SecurityOpenIdConnect)
		}
		x["openIdConnectUrl"] = []byte(s.OpenIdConnectUrl)
	}
	return json.Marshal(x)
}
