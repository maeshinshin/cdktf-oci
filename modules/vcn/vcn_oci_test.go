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
				"DnsLabel":    "vcn-cdktf-oci",
				"CidrBlocks": []string{
					"10.0.0.0/8",
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
				"DnsLabel":    "test",
				"CidrBlocks": []string{
					"10.0.0.0/8",
					"192.168.0.0/16",
				},
				"FreeformTags": map[string]string{
					"provided-by": "cdktf-oci",
				},
			},
			options: []Option{
				WithVcnName("test"),
			},
		},
		{
			name: "with cidr blocks",
			properties: &map[string]any{
				"DisplayName": "vcn-cdktf-oci",
				"DnsLabel":    "vcn-cdktf-oci",
				"CidrBlocks": []string{
					"10.0.0.0/8",
					"172.16.0.0/12",
				},
				"FreeformTags": map[string]string{
					"provided-by": "cdktf-oci",
				},
			},
			options: []Option{
				WithCidrBlocks([]*string{
					jsii.String("10.0.0.0/8"),
					jsii.String("172.16.0.0/12"),
				}),
			},
		},
		{
			name: "with both options",
			properties: &map[string]any{
				"DisplayName": "test",
				"DnsLabel":    "test",
				"CidrBlocks": []string{
					"10.0.0.0/8",
					"172.16.0.0/12",
				},
				"FreeformTags": map[string]string{
					"provided-by": "cdktf-oci",
				},
			},
			options: []Option{
				WithVcnName("test"),
				WithCidrBlocks([]*string{
					jsii.String("10.0.0.0/8"),
					jsii.String("172.16.0.0/12"),
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

			assertion := cdktf.Testing_ToHaveResourceWithProperties(synth, corevcn.CoreVcn_TfResourceType(), tt.properties)

			if !*assertion {
				t.Errorf("%s: Assertion Failed", tt.name)
			}
		})
	}
}
