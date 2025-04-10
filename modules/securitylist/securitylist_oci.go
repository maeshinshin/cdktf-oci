package securitylist

import (
	"github.com/aws/constructs-go/constructs/v10"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/coresecuritylist"
	"github.com/maeshinshin/cdktf-oci/internal/util"
)

type securitylist struct {
	coresecuritylist.CoreSecurityList
}

type Config struct {
	util.Config
	DisplayName          string
	EgressSecurityRules  *[]*coresecuritylist.CoreSecurityListEgressSecurityRules
	IngressSecurityRules *[]*coresecuritylist.CoreSecurityListIngressSecurityRules
}

func NewPublicAndPrivateSecurityList(stack constructs.Construct, compartmentId *string, vcnId *string, optionsForPublic []Option, optionsForPrivate []Option) (*securitylist, *securitylist) {
	publicSecurityList := NewPublicSecurityList(stack, compartmentId, vcnId, optionsForPublic...)
	privateSecurityList := NewPrivateSecurityList(stack, compartmentId, vcnId, optionsForPrivate...)

	return publicSecurityList, privateSecurityList
}

type Option func(*Config)

func WithDisplayName(name string) Option {
	return func(c *Config) {
		c.DisplayName = name
	}
}

func AddIngressSecurityRule(newIngressSecurityRules ...coresecuritylist.CoreSecurityListIngressSecurityRules) Option {
	return func(c *Config) {
		oldRules := *c.IngressSecurityRules
		newRules := make([]*coresecuritylist.CoreSecurityListIngressSecurityRules, 0, len(oldRules)+len(newIngressSecurityRules))

		for _, rule := range oldRules {
			newRules = append(newRules, rule)
		}

		for i := range newIngressSecurityRules {
			newRule := newIngressSecurityRules[i]
			newRules = append(newRules, &newRule)
		}

		c.IngressSecurityRules = &newRules
	}
}

func AddFreeformTags(tags map[string]*string) Option {
	return func(c *Config) {
		for k, v := range tags {
			(*(c.FreeformTags))[k] = v
		}
	}
}
