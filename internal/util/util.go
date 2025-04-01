package util

import (
	"github.com/aws/jsii-runtime-go"
)

type FreeFormTags *map[string]*string

type Config struct {
	FreeformTags FreeFormTags
}

var DefaultCompartmentTags = &map[string]*string{
	"provided-by": jsii.String("cdktf-oci"),
}
