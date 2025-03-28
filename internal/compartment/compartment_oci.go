package compartment

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	"github.com/maeshinshin/cdktf-oci/generated/oracle/oci/identitycompartment"
)

func CreateCompartment(stack constructs.Construct) {
	compartmentConfig := &identitycompartment.IdentityCompartmentConfig{
		Description: jsii.String("CDKTF OCI Compartment"),
		Name:        jsii.String("cdktf-oci"),
	}

	identitycompartment.NewIdentityCompartment(stack, jsii.String("compartment-cdktf-oci"), compartmentConfig)
}
