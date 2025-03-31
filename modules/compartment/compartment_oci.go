package compartment

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/identitycompartment"
)

type Compartment struct {
	identitycompartment.IdentityCompartment
}

type Config struct {
	CompartmentName        string
	CompartmentDescription string
}

type Option func(*Config)

func NewCompartment(stack constructs.Construct, options ...Option) *Compartment {
	config := &Config{
		CompartmentName:        defaultCompartmentName,
		CompartmentDescription: defaultCompartmentDescription,
	}

	for _, option := range options {
		option(config)
	}

	compartmentConfig := &identitycompartment.IdentityCompartmentConfig{
		Name:        jsii.String(config.CompartmentName),
		Description: jsii.String(config.CompartmentDescription),
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
