package provider

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/provider"
)

func SetOciProvider(stack constructs.Construct) {
	fingerprint := cdktf.NewTerraformVariable(stack, jsii.String("FingerPrint"), &cdktf.TerraformVariableConfig{
		Type:      jsii.String("string"),
		Default:   jsii.String(""),
		Sensitive: jsii.Bool(true),
	})

	privateKey := cdktf.NewTerraformVariable(stack, jsii.String("PrivateKey"), &cdktf.TerraformVariableConfig{
		Type:      jsii.String("string"),
		Default:   jsii.String(""),
		Sensitive: jsii.Bool(true),
	})

	privateKeyPassword := cdktf.NewTerraformVariable(stack, jsii.String("PrivateKeyPassword"), &cdktf.TerraformVariableConfig{
		Type:      jsii.String("string"),
		Default:   jsii.String(""),
		Sensitive: jsii.Bool(true),
	})

	region := cdktf.NewTerraformVariable(stack, jsii.String("Region"), &cdktf.TerraformVariableConfig{
		Type:      jsii.String("string"),
		Default:   jsii.String(""),
		Sensitive: jsii.Bool(true),
	})

	tenancyOcid := cdktf.NewTerraformVariable(stack, jsii.String("TenancyOcid"), &cdktf.TerraformVariableConfig{
		Type:      jsii.String("string"),
		Default:   jsii.String(""),
		Sensitive: jsii.Bool(true),
	})

	userOcid := cdktf.NewTerraformVariable(stack, jsii.String("UserOcid"), &cdktf.TerraformVariableConfig{
		Type:      jsii.String("string"),
		Default:   jsii.String(""),
		Sensitive: jsii.Bool(true),
	})

	ociProviderConfig := &provider.OciProviderConfig{
		Fingerprint:        fingerprint.StringValue(),
		PrivateKey:         privateKey.StringValue(),
		PrivateKeyPassword: privateKeyPassword.StringValue(),
		Region:             region.StringValue(),
		TenancyOcid:        tenancyOcid.StringValue(),
		UserOcid:           userOcid.StringValue(),
	}

	provider.NewOciProvider(stack, jsii.String("oci"), ociProviderConfig)
}
