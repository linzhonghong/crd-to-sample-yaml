package pkg

import "github.com/Skarlso/crd-to-sample-yaml/v1beta1"

// Rendering provides extra rendering information of this schema.
type Rendering struct {
	// Group defines which group this schema should belong to. If empty
	// the schema's version will be used as grouping information.
	Group string
}

// SchemaType is a wrapper around any kind of object that provide the following:
// - kind
// - group
// - name
// - openAPIV3Schema.
type SchemaType struct {
	Schema     *v1beta1.JSONSchemaProps
	Versions   []*CRDVersion
	Validation *Validation
	Group      string
	Kind       string

	Rendering Rendering
}

// CRDVersion corresponds to a CRD version.
type CRDVersion struct {
	Name   string
	Schema *v1beta1.JSONSchemaProps
}

// Validation is a set of validation rules that should be applied to all versions.
type Validation struct {
	Name   string
	Schema *v1beta1.JSONSchemaProps
}
