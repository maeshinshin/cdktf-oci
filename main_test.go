package main

import (
	"testing"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/corevcn"
	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/identitycompartment"
	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/provider"
)

var (
	stack = NewMyStack(cdktf.Testing_App(nil), "cdktf-oci")
	synth = cdktf.Testing_Synth(stack, jsii.Bool(true))
)

func TestShouldContainOciProvider(t *testing.T) {
	assertion := cdktf.Testing_ToHaveProvider(synth, provider.OciProvider_TfResourceType())

	if !*assertion {
		t.Error("Assertion Failed")
	}
}

func TestShouldContainIdentityCompartment(t *testing.T) {
	assertion := cdktf.Testing_ToHaveResource(synth, identitycompartment.IdentityCompartment_TfResourceType())

	if !*assertion {
		t.Error("Assertion Failed")
	}
}

func TestShouldContainIdentityCompartmentWithCorrectPropaties(t *testing.T) {
	properties := &map[string]any{
		"Name":        "compartment-cdktf-oci",
		"Description": "CDKTF OCI Compartment",
	}

	assertion := cdktf.Testing_ToHaveResourceWithProperties(synth, identitycompartment.IdentityCompartment_TfResourceType(), properties)

	if !*assertion {
		t.Error("Assertion Failed")
	}
}

func TestShouldVnc(t *testing.T) {
	assertion := cdktf.Testing_ToHaveResource(synth, corevcn.CoreVcn_TfResourceType())

	if !*assertion {
		t.Error("Assertion Failed")
	}
}

func TestShouldVncWithCorrectPropaties(t *testing.T) {
	properties := &map[string]any{
		"DisplayName": "vcn-cdktf-oci",
		"DnsLabel":    "vcncdktfoci",
		"CidrBlocks": []string{
			"10.0.0.0/16",
			"192.168.0.0/16",
		},
		"FreeformTags": map[string]string{
			"provided-by": "cdktf-oci",
		},
	}

	assertion := cdktf.Testing_ToHaveResourceWithProperties(synth, corevcn.CoreVcn_TfResourceType(), properties)
	if !*assertion {
		t.Error("Assertion Failed")
	}
}

func TestCheckValidity(t *testing.T) {
	assertion := cdktf.Testing_ToBeValidTerraform(cdktf.Testing_FullSynth(stack))

	if !*assertion {
		t.Error("Assertion Failed")
	}
}
