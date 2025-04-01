package compartment

import (
	"testing"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/identitycompartment"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestCreateCompartmentWithCorrectPropaties(t *testing.T) {
	test := []struct {
		name       string
		properties *map[string]any
		options    []Option
	}{
		{
			name: "no options",
			properties: &map[string]any{
				"Name":        "cdktf-oci",
				"Description": "CDKTF OCI Compartment",
				"FreeformTags": map[string]string{
					"provided-by": "cdktf-oci",
				},
			},
		},
		{
			name: "with compartment name",
			properties: &map[string]any{
				"Name":        "test",
				"Description": "CDKTF OCI Compartment",
				"FreeformTags": map[string]string{
					"provided-by": "cdktf-oci",
				},
			},
			options: []Option{
				WithCompartmentName("test"),
			},
		},
		{
			name: "with compartment description",
			properties: &map[string]any{
				"Name":        "cdktf-oci",
				"Description": "test",
				"FreeformTags": map[string]string{
					"provided-by": "cdktf-oci",
				},
			},
			options: []Option{
				WithCompartmentDescription("test"),
			},
		},
		{
			name: "with both options",
			properties: &map[string]any{
				"Name":        "test",
				"Description": "test",
				"FreeformTags": map[string]string{
					"provided-by": "cdktf-oci",
				},
			},
			options: []Option{
				WithCompartmentName("test"),
				WithCompartmentDescription("test"),
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			stack := cdktf.NewTerraformStack(cdktf.Testing_App(nil), jsii.String("test"))
			NewCompartment(stack, tt.options...)
			synth := cdktf.Testing_Synth(stack, jsii.Bool(false))

			assertion := cdktf.Testing_ToHaveResourceWithProperties(synth, identitycompartment.IdentityCompartment_TfResourceType(), tt.properties)

			if !*assertion {
				t.Errorf("%s: Assertion Failed", tt.name)
			}
		})
	}
}
