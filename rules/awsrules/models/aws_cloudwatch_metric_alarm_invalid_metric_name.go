// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsCloudwatchMetricAlarmInvalidMetricNameRule checks the pattern is valid
type AwsCloudwatchMetricAlarmInvalidMetricNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCloudwatchMetricAlarmInvalidMetricNameRule returns new rule with default attributes
func NewAwsCloudwatchMetricAlarmInvalidMetricNameRule() *AwsCloudwatchMetricAlarmInvalidMetricNameRule {
	return &AwsCloudwatchMetricAlarmInvalidMetricNameRule{
		resourceType:  "aws_cloudwatch_metric_alarm",
		attributeName: "metric_name",
		max:           255,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCloudwatchMetricAlarmInvalidMetricNameRule) Name() string {
	return "aws_cloudwatch_metric_alarm_invalid_metric_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchMetricAlarmInvalidMetricNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchMetricAlarmInvalidMetricNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchMetricAlarmInvalidMetricNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchMetricAlarmInvalidMetricNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"metric_name must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"metric_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}