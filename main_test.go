package main

import (
	"testing"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/coresecuritylist"
	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/corevcn"
	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/identitycompartment"
	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/provider"
)

var (
	stack = NewMyStack(cdktf.Testing_App(nil), "cdktf-oci")
	synth = cdktf.Testing_Synth(stack, jsii.Bool(true))
)

func TestShouldContainOciProvider(t *testing.T) {
	if assertion := cdktf.Testing_ToHaveProvider(synth, provider.OciProvider_TfResourceType()); !*assertion {
		t.Error("Assertion Failed")
	}
}

func TestShouldContainIdentityCompartment(t *testing.T) {
	if assertion := cdktf.Testing_ToHaveResource(synth, identitycompartment.IdentityCompartment_TfResourceType()); !*assertion {
		t.Error("Assertion Failed")
	}
}

func TestShouldContainIdentityCompartmentWithCorrectPropaties(t *testing.T) {
	properties := &map[string]any{
		"Name":        "compartment-cdktf-oci",
		"Description": "CDKTF OCI Compartment",
	}

	if assertion := cdktf.Testing_ToHaveResourceWithProperties(synth, identitycompartment.IdentityCompartment_TfResourceType(), properties); !*assertion {
		t.Error("Assertion Failed")
	}
}

func TestShouldContainVnc(t *testing.T) {
	if assertion := cdktf.Testing_ToHaveResource(synth, corevcn.CoreVcn_TfResourceType()); !*assertion {
		t.Error("Assertion Failed")
	}
}

func TestShouldContainVncWithCorrectPropaties(t *testing.T) {
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

	if assertion := cdktf.Testing_ToHaveResourceWithProperties(synth, corevcn.CoreVcn_TfResourceType(), properties); !*assertion {
		t.Error("Assertion Failed")
	}
}

func TestShouldContainPublicSecurityList(t *testing.T) {
	properties := &map[string]any{
		"DisplayName": "securitylist-public-cdktf-oci",
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
				"Description": "Allow ssh traffic from the internet",
				"Stateless":   false,
				"Source":      "0.0.0.0/0",
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

	if assertion := cdktf.Testing_ToHaveResourceWithProperties(synth, coresecuritylist.CoreSecurityList_TfResourceType(), properties); !*assertion {
		t.Errorf("%s: Assertion Failed", "public security list")
	}
}

func TestShouldContainPrivateSecurityList(t *testing.T) {
	properties := &map[string]any{
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

	if assertion := cdktf.Testing_ToHaveResourceWithProperties(synth, coresecuritylist.CoreSecurityList_TfResourceType(), properties); !*assertion {
		t.Error("Assertion Failed")
	}
}

func TestCheckValidity(t *testing.T) {
	if assertion := cdktf.Testing_ToBeValidTerraform(cdktf.Testing_FullSynth(stack)); !*assertion {
		t.Error("Assertion Failed")
	}
}
