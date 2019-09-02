// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsElasticBeanstalkApplicationVersionInvalidNameRule checks the pattern is valid
type AwsElasticBeanstalkApplicationVersionInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsElasticBeanstalkApplicationVersionInvalidNameRule returns new rule with default attributes
func NewAwsElasticBeanstalkApplicationVersionInvalidNameRule() *AwsElasticBeanstalkApplicationVersionInvalidNameRule {
	return &AwsElasticBeanstalkApplicationVersionInvalidNameRule{
		resourceType:  "aws_elastic_beanstalk_application_version",
		attributeName: "name",
		max:           100,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsElasticBeanstalkApplicationVersionInvalidNameRule) Name() string {
	return "aws_elastic_beanstalk_application_version_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElasticBeanstalkApplicationVersionInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElasticBeanstalkApplicationVersionInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElasticBeanstalkApplicationVersionInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElasticBeanstalkApplicationVersionInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 100 characters or less",
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
			return nil
		})
	})
}