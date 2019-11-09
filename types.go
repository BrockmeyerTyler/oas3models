package oasm

import (
	"encoding/json"
	"strconv"
)

type JSON map[string]json.RawMessage

// Holds the relative paths to the individual endpoints and their operations. The path is appended to the URL from the
// Server Object in order to construct the full URL. The Paths MAY be empty, due to ACL constraints.
type PathsMap map[string]PathItem

// A map of possible out-of band callbacks related to the parent operation. Each value in the map is a Path Item Object
// that describes a set of requests that may be initiated by the API provider and the expected responses.
// The key value used to identify the callback object is an expression, evaluated at runtime,
// that identifies a URL to use for the callback operation.
type CallbacksMap map[string]PathItem

// A map of example names to their Example Objects.
type ExamplesMap map[string]Example

// A map of mime types to Media Type Objects
type MediaTypesMap map[string]MediaType

// Allows configuration of the supported OAuth Flows.
// A map from a OAuth Flow types to OAuth Flow objects.
type OAuthFlowsMap map[string]OAuthFlow

func marshalStrIfLen(value string, property string, object JSON) {
	if value != "" {
		object[property] = json.RawMessage(strconv.Quote(value))
	}
}
