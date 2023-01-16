# Elastic Container Registry

This module creates a simple container registry using AWS ECR.

## Examples

### Basic Usage
This example demonstrates the basic usage of the module
```hcl
module "ecr" {
  source  = "app.terraform.io/ovotech/container-registry/aws"
  version = "1.0.0"

  ecr_name = "my-container-registry"
}
```
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 3.0.1 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_ecr_name"></a> [ecr\_name](#input\_ecr\_name) | (Required) Name of the elastic container registry being created | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_repo_arn"></a> [repo\_arn](#output\_repo\_arn) | Full ARN of the repository. |
| <a name="output_repo_name"></a> [repo\_name](#output\_repo\_name) | Provided name of the ECR repository. |
| <a name="output_repo_uri"></a> [repo\_uri](#output\_repo\_uri) | The URL of the ECR repo. |
