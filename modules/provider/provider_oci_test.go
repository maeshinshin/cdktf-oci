package provider

import (
	"testing"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/provider"
)

var stack = cdktf.NewTerraformStack(cdktf.Testing_App(nil), jsii.String("test"))

func TestMain(m *testing.M) {
	SetOciProvider(stack)
	m.Run()
}

func TestSetOciProvider(t *testing.T) {
	synth := cdktf.Testing_Synth(stack, jsii.Bool(true))

	if assertion := cdktf.Testing_ToHaveProvider(synth, provider.OciProvider_TfResourceType()); !*assertion {
		t.Error("Assertion Failed")
	}
}
