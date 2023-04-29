# Project Resource

This code is for creating or accessing projects in Terraform. Projects group resources and can be used to create and manage them.

To create a project, use this code:

```terraform
resource "swifdog_project" "test" {
  name        = "test"
  description = "Mein erstes Swifdog Projekt mit Terraform!"
}
```

## Argument Reference

- `name` - (Required, string) The infrastructure-wide unique name of a project.
- `description` - (Optional, string) Visual information for you to keep organized when having multiple projects.
