package securitylist

import (
	"testing"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/coresecuritylist"
	"github.com/maeshinshin/cdktf-oci/modules/compartment"
	"github.com/maeshinshin/cdktf-oci/modules/vcn"
)

var propaties_private1 = &map[string]any{
	"DisplayName": "securitylist-private-cdktf-oci",
	"EgressSecurityRules": []map[string]any{
		{
			"Description":     "Allow all traffic to the internet",
			"Stateless":       false,
			"Destination":     "0.0.0.0/0",
			"DestinationType": "CIDR_BLOCK",
			"Protocol":        "all",
		},
	},
	"IngressSecurityRules": []map[string]any{
		{
			"Description": "Allow ssh traffic from the VCN CIDR block",
			"Stateless":   false,
			"Source":      "10.0.0.0/16",
			"SourceType":  "CIDR_BLOCK",
			"Protocol":    "6",
			"TcpOptions": map[string]any{
				"Max": 22,
				"Min": 22,
			},
		},
		{
			"Description": "Allow ssh traffic from the VCN CIDR block",
			"Stateless":   false,
			"Source":      "192.168.0.0/16",
			"SourceType":  "CIDR_BLOCK",
			"Protocol":    "6",
			"TcpOptions": map[string]any{
				"Max": 22,
				"Min": 22,
			},
		},
		{
			"Description": "Allow icmp type 3 code 4 traffic from the internet",
			"Stateless":   false,
			"Source":      "0.0.0.0/0",
			"SourceType":  "CIDR_BLOCK",
			"Protocol":    "1",
			"IcmpOptions": map[string]any{
				"Type": 3,
				"Code": 4,
			},
		},
		{
			"Description": "Allow icmp type 3 traffic from the VCN CIDR block",
			"Stateless":   false,
			"Source":      "10.0.0.0/16",
			"SourceType":  "CIDR_BLOCK",
			"Protocol":    "1",
			"IcmpOptions": map[string]any{
				"Type": 3,
			},
		},
		{
			"Description": "Allow icmp type 3 traffic from the VCN CIDR block",
			"Stateless":   false,
			"Source":      "192.168.0.0/16",
			"SourceType":  "CIDR_BLOCK",
			"Protocol":    "1",
			"IcmpOptions": map[string]any{
				"Type": 3,
			},
		},
	},
	"FreeformTags": map[string]string{
		"provided-by": "cdktf-oci",
	},
}

var (
	propaties_private2 = &map[string]any{}
	propaties_private3 = &map[string]any{}
	propaties_private4 = &map[string]any{}
)

func TestNewPrivateSecurityListWithCorrectPropaties(t *testing.T) {
	for k, v := range *propaties_private1 {
		(*propaties_private2)[k] = v
		(*propaties_private3)[k] = v
	}
	(*propaties_private2)["DisplayName"] = "test"
	(*propaties_private3)["IngressSecurityRules"] = append((*propaties_private3)["IngressSecurityRules"].([]map[string]any), map[string]any{
		"Description": "Allow telnet traffic from the VCN CIDR block",
		"Stateless":   false,
		"Source":      "10.0.0.0/16",
		"SourceType":  "CIDR_BLOCK",
		"Protocol":    "6",
		"TcpOptions": map[string]any{
			"Max": 21,
			"Min": 21,
		},
	})
	(*propaties_private3)["IngressSecurityRules"] = append((*propaties_private3)["IngressSecurityRules"].([]map[string]any), map[string]any{
		"Description": "Allow telnet traffic from the VCN CIDR block",
		"Stateless":   false,
		"Source":      "192.168.0.0/16",
		"SourceType":  "CIDR_BLOCK",
		"Protocol":    "6",
		"TcpOptions": map[string]any{
			"Max": 21,
			"Min": 21,
		},
	})
	for k, v := range *propaties_private3 {
		(*propaties_private4)[k] = v
	}
	(*propaties_private4)["DisplayName"] = "test"

	test := []struct {
		name       string
		properties *map[string]any
		options    []Option
	}{
		{
			name:       "no options",
			properties: propaties_private1,
		},
		{
			name:       "with Display Name",
			properties: propaties_private2,
			options: []Option{
				WithDisplayName("test"),
			},
		},
		{
			name:       "add ingress security rule",
			properties: propaties_private3,
			options: []Option{
				AddIngressSecurityRule(coresecuritylist.CoreSecurityListIngressSecurityRules{
					Description: jsii.String("Allow telnet traffic from the VCN CIDR block"),
					Stateless:   jsii.Bool(false),
					Source:      jsii.String("10.0.0.0/16"),
					SourceType:  jsii.String("CIDR_BLOCK"),
					Protocol:    jsii.String("6"),
					TcpOptions: &coresecuritylist.CoreSecurityListIngressSecurityRulesTcpOptions{
						Max: jsii.Number(21),
						Min: jsii.Number(21),
					},
				},
					coresecuritylist.CoreSecurityListIngressSecurityRules{
						Description: jsii.String("Allow telnet traffic from the VCN CIDR block"),
						Stateless:   jsii.Bool(false),
						Source:      jsii.String("192.168.0.0/16"),
						SourceType:  jsii.String("CIDR_BLOCK"),
						Protocol:    jsii.String("6"),
						TcpOptions: &coresecuritylist.CoreSecurityListIngressSecurityRulesTcpOptions{
							Max: jsii.Number(21),
							Min: jsii.Number(21),
						},
					},
				),
			},
		},
		{
			name:       "with all options",
			properties: propaties_private4,
			options: []Option{
				WithDisplayName("test"),
				AddIngressSecurityRule(coresecuritylist.CoreSecurityListIngressSecurityRules{
					Description: jsii.String("Allow telnet traffic from the VCN CIDR block"),
					Stateless:   jsii.Bool(false),
					Source:      jsii.String("10.0.0.0/16"),
					SourceType:  jsii.String("CIDR_BLOCK"),
					Protocol:    jsii.String("6"),
					TcpOptions: &coresecuritylist.CoreSecurityListIngressSecurityRulesTcpOptions{
						Max: jsii.Number(21),
						Min: jsii.Number(21),
					},
				},
					coresecuritylist.CoreSecurityListIngressSecurityRules{
						Description: jsii.String("Allow telnet traffic from the VCN CIDR block"),
						Stateless:   jsii.Bool(false),
						Source:      jsii.String("192.168.0.0/16"),
						SourceType:  jsii.String("CIDR_BLOCK"),
						Protocol:    jsii.String("6"),
						TcpOptions: &coresecuritylist.CoreSecurityListIngressSecurityRulesTcpOptions{
							Max: jsii.Number(21),
							Min: jsii.Number(21),
						},
					},
				),
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			stack := cdktf.NewTerraformStack(cdktf.Testing_App(nil), jsii.String("test"))
			compartment := compartment.NewCompartment(stack)
			vcn := vcn.NewVcn(stack, compartment.Id())
			NewPrivateSecurityList(stack, compartment.Id(), vcn.Id(), tt.options...)
			synth := cdktf.Testing_Synth(stack, jsii.Bool(false))

			if assertion := cdktf.Testing_ToHaveResourceWithProperties(synth, coresecuritylist.CoreSecurityList_TfResourceType(), tt.properties); !*assertion {
				t.Errorf("%s: Assertion Failed", tt.name)
			}
		})
	}
}
