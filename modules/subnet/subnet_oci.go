package subnet

import (
	"github.com/aws/constructs-go/constructs/v10"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/coresubnet"
	"github.com/maeshinshin/cdktf-oci/internal/util"
)

type subnet struct {
	coresubnet.CoreSubnet
}

func NewPublicAndPrivateSubnet(stack constructs.Construct, compartmentId *string, vcnId *string, publicSubnetCidrBlock string, privateSubnetCidrBlock string, publicSecurityListIds []*string, privateSecurityListIds []*string, optionsForPublic []Option, optionsForPrivate []Option) (*subnet, *subnet) {
	var privateSecurityListIdList []string
	var publicSecurityListIdList []string

	for _, id := range privateSecurityListIds {
		privateSecurityListIdList = append(privateSecurityListIdList, *id)
	}

	for _, id := range publicSecurityListIds {
		publicSecurityListIdList = append(publicSecurityListIdList, *id)
	}

	publicSubnet := NewPublicsubnet(stack, compartmentId, vcnId, privateSubnetCidrBlock, publicSecurityListIdList, optionsForPrivate...)
	privateSubnet := NewPrivatesubnet(stack, compartmentId, vcnId, privateSubnetCidrBlock, privateSecurityListIdList, optionsForPrivate...)

	return publicSubnet, privateSubnet
}

type Config struct {
	util.Config
	DisplayName string
}

type Option func(*Config)

func WithDisplayName(name string) Option {
	return func(c *Config) {
		c.DisplayName = name
	}
}
