// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsGameliftFleetInvalidBuildIDRule checks the pattern is valid
type AwsGameliftFleetInvalidBuildIDRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsGameliftFleetInvalidBuildIDRule returns new rule with default attributes
func NewAwsGameliftFleetInvalidBuildIDRule() *AwsGameliftFleetInvalidBuildIDRule {
	return &AwsGameliftFleetInvalidBuildIDRule{
		resourceType:  "aws_gamelift_fleet",
		attributeName: "build_id",
		pattern:       regexp.MustCompile(`^build-\S+|^arn:.*:build\/build-\S+`),
	}
}

// Name returns the rule name
func (r *AwsGameliftFleetInvalidBuildIDRule) Name() string {
	return "aws_gamelift_fleet_invalid_build_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGameliftFleetInvalidBuildIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGameliftFleetInvalidBuildIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGameliftFleetInvalidBuildIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGameliftFleetInvalidBuildIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`build_id does not match valid pattern ^build-\S+|^arn:.*:build\/build-\S+`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
