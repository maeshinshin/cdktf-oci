package util_test

import (
	"testing"

	"github.com/aws/jsii-runtime-go"
	"github.com/stretchr/testify/assert"

	"github.com/maeshinshin/cdktf-oci/internal/util"
)

func TestDefaultCompartmentTags(t *testing.T) {
	defaultTags := util.DefaultCompartmentTags

	assert.NotNil(t, defaultTags, "DefaultCompartmentTags should not be nil")
	assert.Contains(t, *defaultTags, "provided-by", "DefaultCompartmentTags should contain 'provided-by'")
	assert.Equal(t, jsii.String("cdktf-oci"), (*defaultTags)["provided-by"], "Tag value should be 'cdktf-oci'")
}
