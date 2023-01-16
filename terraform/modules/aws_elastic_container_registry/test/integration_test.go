package test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	teststructure "github.com/gruntwork-io/terratest/modules/test-structure"
	helpers "github.com/ovotech/team-cppe/shared-resources/libs/test-helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEcrModule(t *testing.T) {
	dst := teststructure.CopyTerraformFolderToTemp(t, "..", "")
	opts := &terraform.Options{
		TerraformDir: dst,
		EnvVars: map[string]string{
			"AWS_REGION": "eu-west-1",
		},
		Vars: map[string]interface{}{
			"ecr_name": fmt.Sprintf("test-%s", strings.ToLower(random.UniqueId())),
		},
	}

	defer teststructure.RunTestStage(t, "cleanup", func() {
		terraform.Destroy(t, opts)
	})

	teststructure.RunTestStage(t, "apply", func() {
		terraform.Init(t, opts)
		terraform.ApplyAndIdempotent(t, opts)
	})

	teststructure.RunTestStage(t, "assert", func() {
		repo := helpers.GetEcrRepository(t, opts.Vars["ecr_name"].(string))
		assert.Equal(t, types.ImageTagMutabilityImmutable, repo.ImageTagMutability)
		assert.Equal(t, true, repo.ImageScanningConfiguration.ScanOnPush)
		assert.Equal(t, types.EncryptionTypeKms, repo.EncryptionConfiguration.EncryptionType)

		ecrTags := GetEcrTags(t, "eu-west-1", aws.ToString(repo.RepositoryArn))
		assert.Contains(t, ecrTags, "ModuleBy")
		assert.Equal(t, "OVO Tech Production Engineering", ecrTags["ModuleBy"])
	})
}

// GetEcrTagsE fetches the given ecr repositories tags and returns them as a string map of strings.
func GetEcrTagsE(t *testing.T, awsRegion string, arn string) (map[string]string, error) {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(awsRegion))
	if err != nil {
		t.Fatalf("Error initialising AWS config: %s", err)
	}

	ecrClient := ecr.NewFromConfig(cfg)
	out, err := ecrClient.ListTagsForResource(ctx, &ecr.ListTagsForResourceInput{
		ResourceArn: &arn,
	})
	if err != nil {
		return nil, err
	}

	tags := map[string]string{}
	for _, tag := range out.Tags {
		tag := tag
		tags[aws.ToString(tag.Key)] = aws.ToString(tag.Value)
	}

	return tags, nil
}

// GetEcrTags fetches the given ecr repositories tags and returns them as a string map of strings
func GetEcrTags(t *testing.T, awsRegion string, arn string) map[string]string {
	tags, err := GetEcrTagsE(t, awsRegion, arn)
	require.NoError(t, err)
	return tags
}
