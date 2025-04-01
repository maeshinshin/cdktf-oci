package securitylist

import (
	"github.com/aws/jsii-runtime-go"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/coresecuritylist"
)

const (
	defaultPublicSecurityListName = "securitylist-public-cdktf-oci"
)

var defaultPublicIngressSecurityRules = &[]*coresecuritylist.CoreSecurityListIngressSecurityRules{
	{
		Description: jsii.String("Allow ssh traffic from the internet"),
		Stateless:   jsii.Bool(false),
		Source:      jsii.String("0.0.0.0/0"),
		SourceType:  jsii.String("CIDR_BLOCK"),
		Protocol:    jsii.String("6"),
		TcpOptions: &coresecuritylist.CoreSecurityListIngressSecurityRulesTcpOptions{
			Max: jsii.Number(22),
			Min: jsii.Number(22),
		},
	},
	{
		Description: jsii.String("Allow icmp type 3 code 4 traffic from the internet"),
		Stateless:   jsii.Bool(false),
		Source:      jsii.String("0.0.0.0/0"),
		SourceType:  jsii.String("CIDR_BLOCK"),
		Protocol:    jsii.String("1"),
		IcmpOptions: &coresecuritylist.CoreSecurityListIngressSecurityRulesIcmpOptions{
			Type: jsii.Number(3),
			Code: jsii.Number(4),
		},
	},
	{
		Description: jsii.String("Allow icmp type 3 traffic from the VCN CIDR block"),
		Stateless:   jsii.Bool(false),
		Source:      jsii.String("10.0.0.0/16"),
		SourceType:  jsii.String("CIDR_BLOCK"),
		Protocol:    jsii.String("1"),
		IcmpOptions: &coresecuritylist.CoreSecurityListIngressSecurityRulesIcmpOptions{
			Type: jsii.Number(3),
		},
	},
	{
		Description: jsii.String("Allow icmp type 3 traffic from the VCN CIDR block"),
		Stateless:   jsii.Bool(false),
		Source:      jsii.String("192.168.0.0/16"),
		SourceType:  jsii.String("CIDR_BLOCK"),
		Protocol:    jsii.String("1"),
		IcmpOptions: &coresecuritylist.CoreSecurityListIngressSecurityRulesIcmpOptions{
			Type: jsii.Number(3),
		},
	},
}
