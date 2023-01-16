# Reference: https://github.com/terraform-linters/tflint/blob/master/docs/user-guide/config.md
config {
  module = true
}

plugin "aws" {
    enabled = true
    version = "0.17.0"
    source  = "github.com/terraform-linters/tflint-ruleset-aws"
}

# Default tflint rules
# Reference: https://github.com/terraform-linters/tflint/tree/master/docs/rules
rule "terraform_naming_convention" {
  enabled = true
}

rule "terraform_comment_syntax" {
  enabled = true
}

rule "terraform_deprecated_index" {
  enabled = true
}

rule "terraform_deprecated_interpolation" {
  enabled = true
}

rule "terraform_documented_outputs" {
  enabled = true
}

rule "terraform_documented_variables" {
  enabled = true
}

rule "terraform_module_pinned_source" {
  enabled = true
}

rule "terraform_module_version" {
  enabled = true
}

rule "terraform_naming_convention" {
  enabled = true
}

rule "terraform_required_providers" {
  enabled = true
}

rule "terraform_required_version" {
  enabled = true
}

rule "terraform_standard_module_structure" {
  enabled = true
}

rule "terraform_typed_variables" {
  enabled = true
}

rule "terraform_unused_declarations" {
  enabled = true
}

rule "terraform_unused_required_providers" {
  enabled = true
}

rule "terraform_workspace_remote" {
  enabled = true
}

# AWS Plugin Rulesets
# Reference: https://github.com/terraform-linters/tflint-ruleset-aws/blob/master/docs/rules/README.md
rule "aws_s3_bucket_name" {
  enabled = true
  regex = "^[a-z\\-]+$"
}

rule "aws_iam_policy_gov_friendly_arns" {
  enabled = true
}

rule "aws_iam_role_policy_gov_friendly_arns" {
  enabled = true
}

rule "aws_iam_policy_document_gov_friendly_arns" {
  enabled = true
}

rule "aws_resource_missing_tags" {
  enabled = true
  tags = ["ModuleBy"]
}
