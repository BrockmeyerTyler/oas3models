package oas3models

import (
	"fmt"
)

// Describes a single operation parameter.
//
// A unique parameter is defined by a combination of a name and location.
type ParameterDoc struct {

	// REQUIRED. The name of the parameter. Parameter names are case sensitive.
	//
	// * If in is "path", the name field MUST correspond to the associated path segment from the path field in the Paths Object. See Path Templating for further information.
	//
	// * If in is "header" and the name field is "Accept", "Content-Type" or "Authorization", the parameter definition SHALL be ignored.
	//
	// * For all other cases, the name corresponds to the parameter name used by the in property.
	Name string `json:"name"`

	// REQUIRED. The location of the parameter. Possible values are "query", "header", "path" or "cookie".
	In   InRequest `json:"in"`

	// A brief description of the parameter. This could contain examples of use. CommonMark syntax MAY be used for rich text representation.
	Description string `json:"description,omitempty"`

	// Determines whether this parameter is mandatory. If the parameter location is "path", this property is REQUIRED
	// and its value MUST be true. Otherwise, the property MAY be included and its default value is false.
	Required bool `json:"required,omitempty"`

	// Specifies that a parameter is deprecated and SHOULD be transitioned out of usage.
	Deprecated bool `json:"deprecated,omitempty"`

	// Sets the ability to pass empty-valued parameters. This is valid only for query parameters and allows sending a
	// parameter with an empty value. Default value is false. If style is used, and if behavior is n/a
	// (cannot be serialized), the value of allowEmptyValue SHALL be ignored.
	AllowEmptyValue bool `json:"allowEmptyValue,omitempty"`

	// Describes how the parameter value will be serialized depending on the type of the parameter value. Default
	// values (based on value of in): for query - form; for path - simple; for header - simple; for cookie - form.
	Style StyleSetting `json:"style,omitempty"`

	// When this is true, parameter values of type array or object generate separate parameters for each value of the
	// array or key-value pair of the map. For other types of parameters this property has no effect. When style is
	// form, the default value is true. For all other styles, the default value is false.
	Explode bool `json:"explode,omitempty"`

	// Determines whether the parameter value SHOULD allow reserved characters, as defined by
	// RFC3986 :/?#[]@!$&'()*+,;= to be included without percent-encoding. This property only applies to parameters
	// with an in value of query. The default value is false.
	AllowReserved bool `json:"allowReserved,omitempty"`

	// The schema defining the type used for the parameter.
	Schema interface{} `json:"schema,omitempty"`

	// Example of the media type. The example SHOULD match the specified schema and encoding properties if present.
	// The example object is mutually exclusive of the examples object. Furthermore, if referencing a schema which
	// contains an example, the example value SHALL override the example provided by the schema. To represent examples
	// of media types that cannot naturally be represented in JSON or YAML, a string value can contain
	// the example with escaping where necessary.
	Example interface{} `json:"example,omitempty"`

	// Examples of the media type. Each example SHOULD contain a value in the correct format as specified in the
	// parameter encoding. The examples object is mutually exclusive of the example object.
	// Furthermore, if referencing a schema which contains an example, the examples value SHALL
	// override the example provided by the schema.
	Examples ExamplesDoc `json:"examples,omitempty"`

	// A map containing the representations for the parameter. The key is the media type and the value describes it.
	// The map MUST only contain one entry.
	Content MediaTypesDoc `json:"content,omitempty"`
}
func (p *ParameterDoc) Validate() error {
	if p.Name == "" || p.In == "" {
		return fmt.Errorf("for parameters, 'name' and 'in' are required")
	}
	if p.Schema == nil && p.Content == nil || p.Schema != nil && p.Content != nil {
		return fmt.Errorf("for parameters, 'schema' or 'content' are required, but not both")
	}
	return nil
}
