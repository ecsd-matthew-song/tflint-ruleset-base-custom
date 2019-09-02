// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsCodedeployDeploymentGroupInvalidAppNameRule checks the pattern is valid
type AwsCodedeployDeploymentGroupInvalidAppNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCodedeployDeploymentGroupInvalidAppNameRule returns new rule with default attributes
func NewAwsCodedeployDeploymentGroupInvalidAppNameRule() *AwsCodedeployDeploymentGroupInvalidAppNameRule {
	return &AwsCodedeployDeploymentGroupInvalidAppNameRule{
		resourceType:  "aws_codedeploy_deployment_group",
		attributeName: "app_name",
		max:           100,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCodedeployDeploymentGroupInvalidAppNameRule) Name() string {
	return "aws_codedeploy_deployment_group_invalid_app_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodedeployDeploymentGroupInvalidAppNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodedeployDeploymentGroupInvalidAppNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodedeployDeploymentGroupInvalidAppNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodedeployDeploymentGroupInvalidAppNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"app_name must be 100 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"app_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}