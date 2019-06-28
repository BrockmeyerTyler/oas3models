package oas3models

import (
	"fmt"
)

// Configuration details for a supported OAuth Flow
type OAuthFlowDoc struct {

	// REQUIRED for (implicit, authorizationCode).
	// The authorization URL to be used for this flow. This MUST be in the form of a URL.
	AuthorizationUrl string `json:"authorizationUrl,omitempty"`

	// REQUIRED for password, clientCredentials, authorizationCode.
	// The token URL to be used for this flow. This MUST be in the form of a URL.
	TokenUrl string `json:"tokenUrl,omitempty"`

	// The URL to be used for obtaining refresh tokens. This MUST be in the form of a URL.
	RefreshUrl string `json:"refreshUrl,omitempty"`

	// REQUIRED. The available scopes for the OAuth2 security scheme.
	// A map between the scope name and a short description for it.
	Scopes map[string]string `json:"scopes"`
}
func (o *OAuthFlowDoc) Validate() error {
	if o.Scopes == nil {
		return fmt.Errorf("for oauth flows, 'scopes' is required")
	}
	return nil
}
