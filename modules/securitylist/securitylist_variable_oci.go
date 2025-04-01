package securitylist

import (
	"github.com/aws/jsii-runtime-go"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/coresecuritylist"
)

var defaultEgressSecurityRules = &[]*coresecuritylist.CoreSecurityListEgressSecurityRules{
	{
		Description:     jsii.String("Allow all traffic to the internet"),
		Stateless:       jsii.Bool(false),
		Destination:     jsii.String("0.0.0.0/0"),
		DestinationType: jsii.String("CIDR_BLOCK"),
		Protocol:        jsii.String("all"),
	},
}
