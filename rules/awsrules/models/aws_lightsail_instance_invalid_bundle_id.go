// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsLightsailInstanceInvalidBundleIDRule checks the pattern is valid
type AwsLightsailInstanceInvalidBundleIDRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsLightsailInstanceInvalidBundleIDRule returns new rule with default attributes
func NewAwsLightsailInstanceInvalidBundleIDRule() *AwsLightsailInstanceInvalidBundleIDRule {
	return &AwsLightsailInstanceInvalidBundleIDRule{
		resourceType:  "aws_lightsail_instance",
		attributeName: "bundle_id",
		pattern:       regexp.MustCompile(`^.*\S.*$`),
	}
}

// Name returns the rule name
func (r *AwsLightsailInstanceInvalidBundleIDRule) Name() string {
	return "aws_lightsail_instance_invalid_bundle_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLightsailInstanceInvalidBundleIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLightsailInstanceInvalidBundleIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLightsailInstanceInvalidBundleIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLightsailInstanceInvalidBundleIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`bundle_id does not match valid pattern ^.*\S.*$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}