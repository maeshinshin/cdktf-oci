package provider

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/provider"
	"github.com/maeshinshin/cdktf-oci/internal/util"
)

func SetOciProvider(stack constructs.Construct) {
	fingerprint := util.GetTerraformVariable(stack, "FingerPrint")

	privateKey := util.GetTerraformVariable(stack, "PrivateKey")

	privateKeyPassword := util.GetTerraformVariable(stack, "PrivateKeyPassword")

	region := util.GetTerraformVariable(stack, "Region")
	tenancyOcid := util.GetTerraformVariable(stack, "TenancyOcid")

	userOcid := util.GetTerraformVariable(stack, "UserOcid")

	ociProviderConfig := &provider.OciProviderConfig{
		Fingerprint:        (*fingerprint).StringValue(),
		PrivateKey:         (*privateKey).StringValue(),
		PrivateKeyPassword: (*privateKeyPassword).StringValue(),
		Region:             (*region).StringValue(),
		TenancyOcid:        (*tenancyOcid).StringValue(),
		UserOcid:           (*userOcid).StringValue(),
	}

	provider.NewOciProvider(stack, jsii.String("oci"), ociProviderConfig)
}
