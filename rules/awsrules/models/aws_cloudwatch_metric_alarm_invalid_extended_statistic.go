// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsCloudwatchMetricAlarmInvalidExtendedStatisticRule checks the pattern is valid
type AwsCloudwatchMetricAlarmInvalidExtendedStatisticRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsCloudwatchMetricAlarmInvalidExtendedStatisticRule returns new rule with default attributes
func NewAwsCloudwatchMetricAlarmInvalidExtendedStatisticRule() *AwsCloudwatchMetricAlarmInvalidExtendedStatisticRule {
	return &AwsCloudwatchMetricAlarmInvalidExtendedStatisticRule{
		resourceType:  "aws_cloudwatch_metric_alarm",
		attributeName: "extended_statistic",
		pattern:       regexp.MustCompile(`^p(\d{1,2}(\.\d{0,2})?|100)$`),
	}
}

// Name returns the rule name
func (r *AwsCloudwatchMetricAlarmInvalidExtendedStatisticRule) Name() string {
	return "aws_cloudwatch_metric_alarm_invalid_extended_statistic"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchMetricAlarmInvalidExtendedStatisticRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchMetricAlarmInvalidExtendedStatisticRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchMetricAlarmInvalidExtendedStatisticRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchMetricAlarmInvalidExtendedStatisticRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`extended_statistic does not match valid pattern ^p(\d{1,2}(\.\d{0,2})?|100)$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}