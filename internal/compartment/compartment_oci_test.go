package compartment

import (
	"testing"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/identitycompartment"
)

var stack = cdktf.NewTerraformStack(cdktf.Testing_App(nil), jsii.String("test"))

func TestMain(m *testing.M) {
	CreateCompartment(stack)
	m.Run()
}

func TestCreateCompartment(t *testing.T) {
	synth := cdktf.Testing_Synth(stack, jsii.Bool(false))

	assertion := cdktf.Testing_ToHaveResource(synth, identitycompartment.IdentityCompartment_TfResourceType())

	if !*assertion {
		t.Error("Assertion Failed")
	}
}

func TestCreateCompartmentWithCorrectPropaties(t *testing.T) {
	synth := cdktf.Testing_Synth(stack, jsii.Bool(false))

	properties := &map[string]any{
		"Description": "CDKTF OCI Compartment",
		"Name":        "cdktf-oci",
	}

	assertion := cdktf.Testing_ToHaveResourceWithProperties(synth, identitycompartment.IdentityCompartment_TfResourceType(), properties)

	if !*assertion {
		t.Error("Assertion Failed")
	}
}
