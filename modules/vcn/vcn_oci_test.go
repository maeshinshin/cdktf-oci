package vcn

import (
	"testing"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/corevcn"
	"github.com/maeshinshin/cdktf-oci/modules/compartment"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestNewVcnWithCorrectPropaties(t *testing.T) {
	test := []struct {
		name       string
		properties *map[string]any
		options    []Option
	}{
		{
			name: "no options",
			properties: &map[string]any{
				"DisplayName": "vcn-cdktf-oci",
				"DnsLabel":    "vcncdktfoci",
				"CidrBlocks": []string{
					"10.0.0.0/16",
					"192.168.0.0/16",
				},
				"FreeformTags": map[string]string{
					"provided-by": "cdktf-oci",
				},
			},
		},
		{
			name: "with vcn name",
			properties: &map[string]any{
				"DisplayName": "test",
				"DnsLabel":    "vcncdktfoci",
				"CidrBlocks": []string{
					"10.0.0.0/16",
					"192.168.0.0/16",
				},
				"FreeformTags": map[string]string{
					"provided-by": "cdktf-oci",
				},
			},
			options: []Option{
				WithDisplayName("test"),
			},
		},
		{
			name: "with dns label",
			properties: &map[string]any{
				"DisplayName": "vcn-cdktf-oci",
				"DnsLabel":    "test",
				"CidrBlocks": []string{
					"10.0.0.0/16",
					"192.168.0.0/16",
				},
				"FreeformTags": map[string]string{
					"provided-by": "cdktf-oci",
				},
			},
			options: []Option{
				WithDnsLabel("test"),
			},
		},
		{
			name: "with cidr blocks",
			properties: &map[string]any{
				"DisplayName": "vcn-cdktf-oci",
				"DnsLabel":    "vcncdktfoci",
				"CidrBlocks": []string{
					"10.0.0.0/16",
					"172.16.0.0/16",
				},
				"FreeformTags": map[string]string{
					"provided-by": "cdktf-oci",
				},
			},
			options: []Option{
				WithCidrBlocks([]*string{
					jsii.String("10.0.0.0/16"),
					jsii.String("172.16.0.0/16"),
				}),
			},
		},
		{
			name: "with all options",
			properties: &map[string]any{
				"DisplayName": "test",
				"DnsLabel":    "test",
				"CidrBlocks": []string{
					"10.0.0.0/16",
					"172.16.0.0/16",
				},
				"FreeformTags": map[string]string{
					"provided-by": "cdktf-oci",
				},
			},
			options: []Option{
				WithDisplayName("test"),
				WithDnsLabel("test"),
				WithCidrBlocks([]*string{
					jsii.String("10.0.0.0/16"),
					jsii.String("172.16.0.0/16"),
				}),
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			stack := cdktf.NewTerraformStack(cdktf.Testing_App(nil), jsii.String("test"))
			compartment := compartment.NewCompartment(stack)
			NewVcn(stack, compartment.Id(), tt.options...)
			synth := cdktf.Testing_Synth(stack, jsii.Bool(false))

			if assertion := cdktf.Testing_ToHaveResourceWithProperties(synth, corevcn.CoreVcn_TfResourceType(), tt.properties); !*assertion {
				t.Errorf("%s: Assertion Failed", tt.name)
			}
		})
	}
}
