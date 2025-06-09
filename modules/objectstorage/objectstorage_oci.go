package objectstorage

import (
	"github.com/aws/constructs-go/constructs/v10"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/objectstoragebucket"
	"github.com/maeshinshin/cdktf-oci/internal/util"
)

type objectstorage struct {
	objectstoragebucket.ObjectstorageBucket
}

type Config struct {
	util.Config
	Name        string
	Access_Type string
}

func NewTalosObjectstorage(stack constructs.Construct, compartmentId *string, optionsForArm []Option, optionsForAmd []Option) (*objectstorage, *objectstorage) {
	armTalosObjectstorage := NewArmTalosObjectstorage(stack, compartmentId, optionsForArm...)
	amdTalosObjectstorage := NewAmdTalosObjectstorage(stack, compartmentId, optionsForAmd...)

	return armTalosObjectstorage, amdTalosObjectstorage
}

type Option func(*Config)

func WithName(name string) Option {
	return func(c *Config) {
		c.Name = name
	}
}

func AddFreeformTags(tags map[string]*string) Option {
	return func(c *Config) {
		for k, v := range tags {
			(*(c.FreeformTags))[k] = v
		}
	}
}
