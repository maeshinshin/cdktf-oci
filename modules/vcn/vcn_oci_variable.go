package vcn

import (
	"github.com/aws/jsii-runtime-go"
)

const (
	defaultVcnName  = "vcn-cdktf-oci"
	defaultDnsLabel = "vcncdktfoci"
)

var defaultCidrBlocks = &[]*string{
	jsii.String("10.0.0.0/8"),
	jsii.String("192.168.0.0/16"),
}
