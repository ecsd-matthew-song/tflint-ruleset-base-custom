// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsCurReportDefinitionInvalidTimeUnitRule checks the pattern is valid
type AwsCurReportDefinitionInvalidTimeUnitRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCurReportDefinitionInvalidTimeUnitRule returns new rule with default attributes
func NewAwsCurReportDefinitionInvalidTimeUnitRule() *AwsCurReportDefinitionInvalidTimeUnitRule {
	return &AwsCurReportDefinitionInvalidTimeUnitRule{
		resourceType:  "aws_cur_report_definition",
		attributeName: "time_unit",
		enum: []string{
			"HOURLY",
			"DAILY",
		},
	}
}

// Name returns the rule name
func (r *AwsCurReportDefinitionInvalidTimeUnitRule) Name() string {
	return "aws_cur_report_definition_invalid_time_unit"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCurReportDefinitionInvalidTimeUnitRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCurReportDefinitionInvalidTimeUnitRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCurReportDefinitionInvalidTimeUnitRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCurReportDefinitionInvalidTimeUnitRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`time_unit is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}