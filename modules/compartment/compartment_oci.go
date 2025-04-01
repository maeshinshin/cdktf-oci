package compartment

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/identitycompartment"
	"github.com/maeshinshin/cdktf-oci/internal/util"
)

type Compartment struct {
	identitycompartment.IdentityCompartment
}

type Config struct {
	util.Config
	CompartmentName        string
	CompartmentDescription string
}

func NewCompartment(stack constructs.Construct, options ...Option) *Compartment {
	config := &Config{
		CompartmentName:        defaultCompartmentName,
		CompartmentDescription: defaultCompartmentDescription,
		Config: util.Config{
			FreeformTags: util.DefaultTags,
		},
	}

	for _, option := range options {
		option(config)
	}

	compartmentConfig := &identitycompartment.IdentityCompartmentConfig{
		Name:         jsii.String(config.CompartmentName),
		Description:  jsii.String(config.CompartmentDescription),
		FreeformTags: config.FreeformTags,
	}

	compartment := &Compartment{identitycompartment.NewIdentityCompartment(stack, jsii.String("compartment-cdktf-oci"), compartmentConfig)}

	return compartment
}

func WithCompartmentName(name string) Option {
	return func(c *Config) {
		c.CompartmentName = name
	}
}

func WithCompartmentDescription(description string) Option {
	return func(c *Config) {
		c.CompartmentDescription = description
	}
}

func FreeformTags(tags *map[string]*string) Option {
	if tags == nil {
		tags = util.DefaultCompartmentTags
	}
	return func(c *Config) {
		c.FreeformTags = tags
	}
}
