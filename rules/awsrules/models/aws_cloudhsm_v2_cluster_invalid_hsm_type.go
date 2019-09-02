// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsCloudhsmV2ClusterInvalidHsmTypeRule checks the pattern is valid
type AwsCloudhsmV2ClusterInvalidHsmTypeRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsCloudhsmV2ClusterInvalidHsmTypeRule returns new rule with default attributes
func NewAwsCloudhsmV2ClusterInvalidHsmTypeRule() *AwsCloudhsmV2ClusterInvalidHsmTypeRule {
	return &AwsCloudhsmV2ClusterInvalidHsmTypeRule{
		resourceType:  "aws_cloudhsm_v2_cluster",
		attributeName: "hsm_type",
		pattern:       regexp.MustCompile(`^(hsm1\.medium)$`),
	}
}

// Name returns the rule name
func (r *AwsCloudhsmV2ClusterInvalidHsmTypeRule) Name() string {
	return "aws_cloudhsm_v2_cluster_invalid_hsm_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudhsmV2ClusterInvalidHsmTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudhsmV2ClusterInvalidHsmTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudhsmV2ClusterInvalidHsmTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudhsmV2ClusterInvalidHsmTypeRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`hsm_type does not match valid pattern ^(hsm1\.medium)$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}