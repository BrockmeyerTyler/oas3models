package oasm

import "encoding/json"

// Describes the operations available on a single path. A Path Item MAY be empty, due to ACL constraints.
// The path itself is still exposed to the documentation viewer but they will not know which operations and
// parameters are available.
type PathItemDoc struct {

	// Allows for an external definition of this path item. The referenced structure MUST be in the format of a Path
	// Item Object. If there are conflicts between the referenced definition and this Path Item's definition,
	// the behavior is undefined.
	Ref string

	// An optional, string summary, intended to apply to all operations in this path.
	Summary string

	// An optional, string description, intended to apply to all operations in this path.
	// CommonMark syntax MAY be used for rich text representation.
	Description string

	// A map of HTTP methods to their operation definitions.
	Methods map[HTTPVerb]*OperationDoc

	// An alternative server array to service all operations in this path.
	Servers []*ServerDoc

	// A list of parameters that are applicable for all the operations described under this path. These parameters can
	// be overridden at the operation level, but cannot be removed there. The list MUST NOT include duplicated
	// parameters. A unique parameter is defined by a combination of a name and location. The list can use the
	// Reference Object to link to parameters that are defined at the OpenAPI Object's components/parameters.
	Parameters []*ParameterDoc
}

func (p *PathItemDoc) MarshalJSON() (_ []byte, err error) {
	x := make(JSON)
	for key, value := range p.Methods {
		if x[string(key)], err = json.Marshal(value); err != nil {
			return
		}
	}
	if p.Servers != nil {
		x["servers"], err = json.Marshal(p.Servers)
		if err != nil {
			return
		}
	}
	if p.Parameters != nil {
		x["parameters"], err = json.Marshal(p.Parameters)
		if err != nil {
			return
		}
	}
	marshalStrIfLen(p.Ref, "$ref", x)
	marshalStrIfLen(p.Summary, "summary", x)
	marshalStrIfLen(p.Description, "description", x)
	return json.Marshal(x)
}
