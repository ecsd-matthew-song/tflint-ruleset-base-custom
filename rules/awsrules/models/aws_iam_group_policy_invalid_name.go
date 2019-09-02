// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsIAMGroupPolicyInvalidNameRule checks the pattern is valid
type AwsIAMGroupPolicyInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIAMGroupPolicyInvalidNameRule returns new rule with default attributes
func NewAwsIAMGroupPolicyInvalidNameRule() *AwsIAMGroupPolicyInvalidNameRule {
	return &AwsIAMGroupPolicyInvalidNameRule{
		resourceType:  "aws_iam_group_policy",
		attributeName: "name",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w+=,.@-]+$`),
	}
}

// Name returns the rule name
func (r *AwsIAMGroupPolicyInvalidNameRule) Name() string {
	return "aws_iam_group_policy_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMGroupPolicyInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMGroupPolicyInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMGroupPolicyInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMGroupPolicyInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`name does not match valid pattern ^[\w+=,.@-]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}