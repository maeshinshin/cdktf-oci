package securitylist

import (
	"testing"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/coresecuritylist"
	"github.com/maeshinshin/cdktf-oci/modules/compartment"
	"github.com/maeshinshin/cdktf-oci/modules/vcn"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestNewPublicAndPrivateSecurityList(t *testing.T) {
	stack := cdktf.NewTerraformStack(cdktf.Testing_App(nil), jsii.String("test"))
	compartment := compartment.NewCompartment(stack)
	vcn := vcn.NewVcn(stack, compartment.Id())
	NewPublicAndPrivateSecurityList(stack, compartment.Id(), vcn.Id(), nil, nil)
	synth := cdktf.Testing_Synth(stack, jsii.Bool(false))

	if assertion := cdktf.Testing_ToHaveResourceWithProperties(synth, coresecuritylist.CoreSecurityList_TfResourceType(), propaties_public1); !*assertion {
		t.Errorf("%s: Assertion Failed", "public security list")
	}

	if assertion := cdktf.Testing_ToHaveResourceWithProperties(synth, coresecuritylist.CoreSecurityList_TfResourceType(), propaties_private1); !*assertion {
		t.Errorf("%s: Assertion Failed", "private security list")
	}
}
