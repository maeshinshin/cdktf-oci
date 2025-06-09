package objectstorage

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/objectstoragebucket"
	"github.com/maeshinshin/cdktf-oci/internal/util"
)

func NewArmTalosObjectstorage(stack constructs.Construct, compartmentId *string, options ...Option) *objectstorage {
	config := &Config{
		Name: defaultTalosArmObjectstorageName,
		Config: util.Config{
			FreeformTags: util.DefaultTags,
		},
		Namespace:   defaultTalosObjectstorageNamespace,
		Access_Type: defaultTalosObjectstorageAccessType,
	}

	for _, option := range options {
		option(config)
	}

	objectstorageConfig := &objectstoragebucket.ObjectstorageBucketConfig{
		CompartmentId: compartmentId,
		Name:          jsii.String(config.Name),
		Namespace:     jsii.String(config.Namespace),
		FreeformTags:  config.FreeformTags,
		AccessType:    jsii.String(config.Access_Type),
	}

	objectstorage := &objectstorage{objectstoragebucket.NewObjectstorageBucket(stack, jsii.String("objectstorage-talos-arm-cdktf-oci"), objectstorageConfig)}

	return objectstorage
}
