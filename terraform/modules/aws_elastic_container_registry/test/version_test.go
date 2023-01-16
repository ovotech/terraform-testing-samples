package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	teststructure "github.com/gruntwork-io/terratest/modules/test-structure"
	helpers "github.com/ovotech/team-cppe/shared-resources/libs/test-helpers"
)

func TestTerraformVersions(t *testing.T) {
	constraint := helpers.GetTerraformVersionConstraint(t, "..")
	available := helpers.GetAvailableVersions(t, "terraform")
	testVers := helpers.GetMatchingVersions(t, constraint, available)

	for _, version := range testVers {
		version := version
		t.Run(version, func(t *testing.T) {
			t.Parallel()

			dst := teststructure.CopyTerraformFolderToTemp(t, "..", ".")
			binaryPath := helpers.DownloadTerraformVersion(t, version)
			tfOptions := &terraform.Options{
				TerraformDir:    dst,
				TerraformBinary: binaryPath,
			}
			tfOptions.EnvVars = map[string]string{
				"AWS_REGION": "eu-west-1",
			}
			tfOptions.Vars = map[string]interface{}{
				"ecr_name": "cppe-container-registry",
			}
			terraform.InitAndPlan(t, tfOptions)
		})
	}
}

func TestAwsProviderVersions(t *testing.T) {
	constraint := helpers.GetProviderConstraint(t, "..", "aws")
	available := helpers.GetAvailableVersions(t, "terraform-provider-aws")
	testVers := helpers.GetMatchingVersions(t, constraint, available)

	for _, version := range testVers {
		version := version
		t.Run(version, func(t *testing.T) {
			t.Parallel()

			dst := teststructure.CopyTerraformFolderToTemp(t, "..", ".")
			helpers.UpdateProviderVersion(t, dst, "aws", version, "hashicorp/aws")
			terraform.InitAndPlan(t, &terraform.Options{
				TerraformDir: dst,
				EnvVars: map[string]string{
					"AWS_REGION": "eu-west-1",
				},
				Vars: map[string]interface{}{
					"ecr_name": "cppe-container-registry",
				},
			})
		})
	}
}
