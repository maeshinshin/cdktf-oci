package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/maeshinshin/cdktf-oci/modules/compartment"
	"github.com/maeshinshin/cdktf-oci/modules/provider"
	"github.com/maeshinshin/cdktf-oci/modules/securitylist"
	"github.com/maeshinshin/cdktf-oci/modules/subnet"
	"github.com/maeshinshin/cdktf-oci/modules/vcn"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, jsii.String(id))

	provider.SetOciProvider(stack)
	compartment := compartment.NewCompartment(stack)
	vcn := vcn.NewVcn(stack, compartment.Id())
	publicSecurityList, privateSecurityList := securitylist.NewPublicAndPrivateSecurityList(stack, compartment.Id(), vcn.Id(), nil, nil)
	subnet.NewPublicAndPrivateSubnet(stack, compartment.Id(), vcn.Id(), "", "", []*string{privateSecurityList.Id()}, []*string{publicSecurityList.Id()}, nil, nil)

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	stack := NewMyStack(app, "cdktf-oci")
	cdktf.NewCloudBackend(stack, &cdktf.CloudBackendConfig{
		Hostname:     jsii.String("app.terraform.io"),
		Organization: jsii.String("maesh-dev"),
		Workspaces:   cdktf.NewNamedCloudWorkspace(jsii.String("cdktf-oci"), nil),
	})

	app.Synth()
}
