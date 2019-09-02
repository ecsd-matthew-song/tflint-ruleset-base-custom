// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/wata727/tflint/tflint"
)

func Test_AwsCognitoResourceServerInvalidNameRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cognito_resource_server" "foo" {
	name = "example/server"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsCognitoResourceServerInvalidNameRule(),
					Message: `name does not match valid pattern ^[\w\s+=,.@-]+$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cognito_resource_server" "foo" {
	name = "example"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsCognitoResourceServerInvalidNameRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}