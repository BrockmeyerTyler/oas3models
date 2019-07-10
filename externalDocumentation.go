package oasm

import (
	"fmt"
)

// Allows referencing an external resource for extended documentation.
type ExternalDocumentationDoc struct {

	// A short description of the target documentation. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`

	// REQUIRED. The URL for the target documentation. Value MUST be in the format of a URL.
	Url string `json:"url"`
}

func (e *ExternalDocumentationDoc) Validate() error {
	if e.Url == "" {
		return fmt.Errorf("for external documentations, 'url' is required")
	}
	return nil
}
