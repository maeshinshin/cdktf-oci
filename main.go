package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/maeshinshin/cdktf-oci/internal/compartment"
	"github.com/maeshinshin/cdktf-oci/internal/provider"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, jsii.String(id))

	provider.SetOciProvider(stack)
	compartment.CreateCompartment(stack)

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
