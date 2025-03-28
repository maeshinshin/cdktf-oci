package main

import (
	"testing"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/provider"
)

var (
	stack = NewMyStack(cdktf.Testing_App(nil), "cdktf-oci")
	synth = cdktf.Testing_Synth(stack, jsii.Bool(true))
)

func TestShuldContainOciProvider(t *testing.T) {
	assertion := cdktf.Testing_ToHaveProvider(synth, provider.OciProvider_TfResourceType())

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
