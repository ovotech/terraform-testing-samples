package testhelpers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// GetTerraformVersionConstraintE returns the Terraform version string for the given module
// or an error if the provider cannot be found
func GetTerraformVersionConstraintE(srcDir string) (string, error) {
	files, err := os.ReadDir(srcDir)
	if err != nil {
		return "", err
	}

	vRegexp := regexp.MustCompile("required_version\\s*=\\s*\"([^\"]+)\"")

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if !strings.HasSuffix(file.Name(), ".tf") {
			continue
		}

		filename := fmt.Sprintf("%s/%s", srcDir, file.Name())
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			return "", err
		}

		f, diag := hclwrite.ParseConfig(content, file.Name(), hcl.Pos{Line: 1, Column: 1})
		if diag.HasErrors() {
			return "", errors.New(diag.Error())
		}

		for _, block := range f.Body().Blocks() {
			if block.Type() != "terraform" {
				continue
			}

			requiredVersionSetting := block.Body().GetAttribute("required_version")
			if requiredVersionSetting == nil {
				continue
			}

			val := requiredVersionSetting.BuildTokens(nil).Bytes()
			constraint := vRegexp.FindSubmatch(val)

			if constraint == nil || len(constraint) < 2 {
				continue
			}

			return string(constraint[1]), nil
		}
	}

	return "", fmt.Errorf("required_version setting not found")
}

// GetTerraformVersionConstraint returns the Terraform version string for the given module
// or fails the test if the version is not found
func GetTerraformVersionConstraint(t *testing.T, srcDir string) string {
	constraint, err := GetTerraformVersionConstraintE(srcDir)
	if err != nil {
		t.Fatalf(err.Error())
	}
	return constraint
}
