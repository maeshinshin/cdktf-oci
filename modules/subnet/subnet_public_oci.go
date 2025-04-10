package subnet

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/coresubnet"
	"github.com/maeshinshin/cdktf-oci/internal/util"
)

func NewPublicsubnet(stack constructs.Construct, compartmentId *string, vcnId *string, cidrBlock string, securityListIds []string, options ...Option) *subnet {
	if cidrBlock == "" {
		cidrBlock = defaultPublicSubnetCidrBlock
	}

	config := &Config{
		DisplayName: defaultPublicSubnetName,
		Config: util.Config{
			FreeformTags: util.DefaultTags,
		},
	}

	for _, option := range options {
		option(config)
	}

	subnetConfig := &coresubnet.CoreSubnetConfig{
		CompartmentId:   compartmentId,
		VcnId:           vcnId,
		DisplayName:     jsii.String(config.DisplayName),
		FreeformTags:    config.FreeformTags,
		CidrBlock:       jsii.String(cidrBlock),
		SecurityListIds: jsii.Strings(securityListIds...),
	}

	subnet := &subnet{coresubnet.NewCoreSubnet(stack, jsii.String("subnet-public-cdktf-oci"), subnetConfig)}

	return subnet
}
