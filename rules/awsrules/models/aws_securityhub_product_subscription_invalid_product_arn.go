// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsSecurityhubProductSubscriptionInvalidProductArnRule checks the pattern is valid
type AwsSecurityhubProductSubscriptionInvalidProductArnRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsSecurityhubProductSubscriptionInvalidProductArnRule returns new rule with default attributes
func NewAwsSecurityhubProductSubscriptionInvalidProductArnRule() *AwsSecurityhubProductSubscriptionInvalidProductArnRule {
	return &AwsSecurityhubProductSubscriptionInvalidProductArnRule{
		resourceType:  "aws_securityhub_product_subscription",
		attributeName: "product_arn",
		pattern:       regexp.MustCompile(`^.*\S.*$`),
	}
}

// Name returns the rule name
func (r *AwsSecurityhubProductSubscriptionInvalidProductArnRule) Name() string {
	return "aws_securityhub_product_subscription_invalid_product_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSecurityhubProductSubscriptionInvalidProductArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSecurityhubProductSubscriptionInvalidProductArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSecurityhubProductSubscriptionInvalidProductArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSecurityhubProductSubscriptionInvalidProductArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`product_arn does not match valid pattern ^.*\S.*$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}