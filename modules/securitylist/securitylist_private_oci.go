package securitylist

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/coresecuritylist"
	"github.com/maeshinshin/cdktf-oci/internal/util"
)

func NewPrivateSecurityList(stack constructs.Construct, compartmentId *string, vcnId *string, options ...Option) *securitylist {
	config := &Config{
		DisplayName: defaultPrivateSecurityListName,
		Config: util.Config{
			FreeformTags: util.DefaultTags,
		},
		EgressSecurityRules:  defaultEgressSecurityRules,
		IngressSecurityRules: defaultPrivateIngressSecurityRules,
	}

	for _, option := range options {
		option(config)
	}

	securitylistConfig := &coresecuritylist.CoreSecurityListConfig{
		CompartmentId:        compartmentId,
		VcnId:                vcnId,
		DisplayName:          jsii.String(config.DisplayName),
		FreeformTags:         config.FreeformTags,
		EgressSecurityRules:  config.EgressSecurityRules,
		IngressSecurityRules: config.IngressSecurityRules,
	}

	securitylist := &securitylist{coresecuritylist.NewCoreSecurityList(stack, jsii.String("securitylist-private-cdktf-oci"), securitylistConfig)}

	return securitylist
}
