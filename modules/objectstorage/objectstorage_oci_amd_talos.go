package objectstorage

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/objectstoragebucket"
	"github.com/maeshinshin/cdktf-oci/internal/util"
)

func NewAmdTalosObjectstorage(stack constructs.Construct, compartmentId *string, options ...Option) *objectstorage {
	namespace := util.GetTerraformVariable(stack, "ObjectStorageNamespace")

	config := &Config{
		Name: defaultTalosAmdObjectstorageName,
		Config: util.Config{
			FreeformTags: util.DefaultTags,
		},
		Access_Type: defaultTalosObjectstorageAccessType,
	}

	for _, option := range options {
		option(config)
	}

	objectstorageConfig := &objectstoragebucket.ObjectstorageBucketConfig{
		CompartmentId: compartmentId,
		Name:          jsii.String(config.Name),
		Namespace:     (*namespace).StringValue(),
		FreeformTags:  config.FreeformTags,
		AccessType:    jsii.String(config.Access_Type),
	}

	objectstorage := &objectstorage{objectstoragebucket.NewObjectstorageBucket(stack, jsii.String("objectstorage-talos-amd-cdktf-oci"), objectstorageConfig)}

	return objectstorage
}
