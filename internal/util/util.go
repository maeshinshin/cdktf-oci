package util

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FreeFormTags *map[string]*string

type Config struct {
	FreeformTags FreeFormTags
}

var TerraformVariables = make(map[string]*cdktf.TerraformVariable)

var DefaultTags = &map[string]*string{
	"provided-by": jsii.String("cdktf-oci"),
}

func GetTerraformVariable(stack constructs.Construct, name string) *cdktf.TerraformVariable {
	if variable, exists := TerraformVariables[name]; exists {
		return variable
	}

	tfvar := cdktf.NewTerraformVariable(stack, jsii.String(name), &cdktf.TerraformVariableConfig{
		Type:      jsii.String("string"),
		Default:   jsii.String(""),
		Sensitive: jsii.Bool(true),
	})
	TerraformVariables[name] = &tfvar

	return &tfvar
}
