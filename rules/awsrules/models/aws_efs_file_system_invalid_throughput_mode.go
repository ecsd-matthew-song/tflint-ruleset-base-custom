// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsEfsFileSystemInvalidThroughputModeRule checks the pattern is valid
type AwsEfsFileSystemInvalidThroughputModeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEfsFileSystemInvalidThroughputModeRule returns new rule with default attributes
func NewAwsEfsFileSystemInvalidThroughputModeRule() *AwsEfsFileSystemInvalidThroughputModeRule {
	return &AwsEfsFileSystemInvalidThroughputModeRule{
		resourceType:  "aws_efs_file_system",
		attributeName: "throughput_mode",
		enum: []string{
			"bursting",
			"provisioned",
		},
	}
}

// Name returns the rule name
func (r *AwsEfsFileSystemInvalidThroughputModeRule) Name() string {
	return "aws_efs_file_system_invalid_throughput_mode"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEfsFileSystemInvalidThroughputModeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEfsFileSystemInvalidThroughputModeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEfsFileSystemInvalidThroughputModeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEfsFileSystemInvalidThroughputModeRule) Check(runner *tflint.Runner) error {
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
					`throughput_mode is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}