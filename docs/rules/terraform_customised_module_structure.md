# terraform_customised_module_structure

Ensure that a module complies with a customised Terraform Module Structure consisting of the following:
* providers.tf
* data.tf
* locals.tf

## Example

```
$ tflint
1 issue(s) found:

Warning: Module should include an empty providers.tf file (terraform_customised_module_structure)

  on providers.tf line 1:
   (source code not available)

Reference: https://github.com/terraform-linters/tflint/blob/v0.37.0/docs/rules/terraform_customised_module_structure.md
```

## Why

As part of the standardisation of the creation of Terraform projects, it is recommended to include these 3 files `providers.tf`, `data.tf` & `locals.tf` for easier management.

## How To Fix

* Create empty files even if no `locals` or `data` blocks are defined
* Encourages the use of the `providers.tf` file
