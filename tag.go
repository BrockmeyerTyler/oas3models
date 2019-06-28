package oas3models

import "fmt"

// Adds metadata to a single tag that is used by the Operation Object.
// It is not mandatory to have a Tag Object per tag defined in the Operation Object instances.
type TagDoc struct {

	// REQUIRED. The name of the tag.
	Name string `json:"name"`

	// A short description for the tag. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`

	// Additional external documentation for this tag.
	ExternalDocs *ExternalDocumentationDoc `json:"externalDocs,omitempty"`
}
func (t *TagDoc) Validate() error {
	if t.Name == "" {
		return fmt.Errorf("for tags, 'name' is required")
	}
	return nil
}
