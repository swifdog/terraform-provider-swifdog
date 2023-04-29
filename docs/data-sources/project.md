# Project Data Source

Provides basic information of a project (like the unique ID). This method will fetch the projects you have access to and
returns the first project which matches the given name.

## Example Usage

```terraform
data "swifdog_project" "test" {
  name      = "test"
}
```

## Argument Reference

- `name` - (Required, string) The name of the project to fetch
