package oas3models

import (
	"encoding/json"
	"fmt"
)

type OAuthFlowDoc struct {
	AuthorizationUrl string // implicit, authorizationCode
	TokenUrl string // password, clientCredentials, authorizationCode
	RefreshUrl string
	Scopes map[string]string // map["name"]"description"
}
func (o *OAuthFlowDoc) MarshalJSON() (_ []byte, err error) {
	if o.Scopes == nil {
		return nil, fmt.Errorf("for oauth flows, 'scopes' is required")
	}
	x := make(JSON)
	marshalStrIfLen(o.AuthorizationUrl, "authorizationUrl", x)
	marshalStrIfLen(o.TokenUrl, "tokenUrl", x)
	marshalStrIfLen(o.RefreshUrl, "refreshUrl", x)
	x["scopes"], _ = json.Marshal(o.Scopes)
	return json.Marshal(x)
}
