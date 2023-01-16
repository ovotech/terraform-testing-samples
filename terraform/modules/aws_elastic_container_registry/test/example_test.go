package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	teststructure "github.com/gruntwork-io/terratest/modules/test-structure"
	helpers "github.com/ovotech/team-cppe/shared-resources/libs/test-helpers"
)

func TestBasicUsageExample(t *testing.T) {
	dst := teststructure.CopyTerraformFolderToTemp(t, "..", "examples/basic-usage")
	helpers.UpdateModuleSourceToPath(t, dst, "*", "../..")
	terraform.InitAndPlan(t, &terraform.Options{
		TerraformDir: dst,
	})
}
