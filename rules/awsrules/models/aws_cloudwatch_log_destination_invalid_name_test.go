// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/wata727/tflint/tflint"
)

func Test_AwsCloudwatchLogDestinationInvalidNameRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cloudwatch_log_destination" "foo" {
	name = "test:destination"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsCloudwatchLogDestinationInvalidNameRule(),
					Message: `name does not match valid pattern ^[^:*]*$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cloudwatch_log_destination" "foo" {
	name = "test_destination"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsCloudwatchLogDestinationInvalidNameRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}