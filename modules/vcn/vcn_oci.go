package vcn

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/corevcn"
	"github.com/maeshinshin/cdktf-oci/internal/util"
)

type Vcn struct {
	corevcn.CoreVcn
}

type Config struct {
	util.Config
	CidrBlocks  *[]*string
	DisplayName string
	DnsLabel    string
}

func NewVcn(stack constructs.Construct, compartmentId *string, options ...Option) *Vcn {
	config := &Config{
		DisplayName: defaultVcnName,
		DnsLabel:    defaultDnsLabel,
		Config: util.Config{
			FreeformTags: util.DefaultTags,
		},
		CidrBlocks: defaultCidrBlocks,
	}

	for _, option := range options {
		option(config)
	}

	vcnConfig := &corevcn.CoreVcnConfig{
		CompartmentId: compartmentId,
		CidrBlocks:    config.CidrBlocks,
		DisplayName:   jsii.String(config.DisplayName),
		DnsLabel:      jsii.String(config.DnsLabel),
		FreeformTags:  config.FreeformTags,
	}

	vcn := &Vcn{corevcn.NewCoreVcn(stack, jsii.String("vcn-cdktf-oci"), vcnConfig)}

	return vcn
}

type Option func(*Config)

func WithDisplayName(name string) Option {
	return func(c *Config) {
		c.DisplayName = name
	}
}

func WithDnsLabel(label string) Option {
	return func(c *Config) {
		c.DnsLabel = label
	}
}

func WithCidrBlocks(cidrBlocks []*string) Option {
	return func(c *Config) {
		c.CidrBlocks = &cidrBlocks
	}
}

func AddFreeformTags(tags map[string]*string) Option {
	return func(c *Config) {
		for k, v := range tags {
			(*(c.FreeformTags))[k] = v
		}
	}
}
