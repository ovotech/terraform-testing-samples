# Terraform Testing Samples

This repository contains code samples demonstrating the Terraform IaC testing steps used by OVO tech production engineering, along with appropriate instructions and any helper resources / libraries that we’ve written to help enable this.

Contained in this repository is a sample module which provisions an AWS Elastic Container Registry for storing container images, along with appropriate examples and tests, as well as any reusable configurations and shared helper resources to suppliment the testing.

## Checkov
Checkov is an open source command line utility (approved by the OVO Security Engineering team) which we use to make sure all our IaC has all the correct security configurations and settings. If for example, an AWS S3 bucket created in our modules is missing encryption or is exposed publicly, this will be flagged by Checkov. 

To run this locally or in the pipeline, we simply need to [install the binary](https://www.checkov.io/1.Welcome/Quick%20Start.html) and run the CLI as follows:
```
checkov -d terraform/modules/aws_elastic_container_registry --quiet --output cli --framework terraform --download-external-modules false
```

## Tflint
Tflint is a pluggable linter that we make use of to ensure we can enforce custom rules statically e.g. the enforcement of certain tags on all our resources.

To run tflint just [install the binary](https://github.com/terraform-linters/tflint) and as an example, run the following commands:
```
tflint --init --config=terraform/.tflint.hcl
cd terraform/modules/aws_elastic_container_registry
terraform get
tflint --config=../../../terraform/.tflint.hcl
```

## Terratest
[Terratest](https://terratest.gruntwork.io/) is what is used in order to perform unit testing and integration testing on our Terraform IaC. The following tests are included in this repository:
* Unit test to ensure we can run `terraform plan` against an example usage of our Terraform module
* Version testing to ensure we can `terraform plan` with all the versions of Terraform and the AWS provider as specified by our version constraints in `versions.tf`
* Integration testing which will actually spin up infrastructure in AWS using `terraform apply`, check for various attributes against the live resources and then cleanup the tests using `terraform destroy`

**Note** As a pre-requisite for running theses tests against resources in AWS, you need to have set [CLI credentials for AWS](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html)

In order to run the test via the Go CLI all you need to do is run the following command in the test folder:
```
go test -v -timeout 30m
```

## Contributions
Contributions are more than welcome from both internal (OVO tech) and external contributors.

If you have write access to this repo, create a branch and a PR, otherwise a fork and a PR.
