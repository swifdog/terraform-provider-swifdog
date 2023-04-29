# Persistent Volume Data Source

Provide details about a persistent volume. A persistent volume is always inside a project, so please provide the
`projectid`!

```terraform
data "swifdog_persistent_volume" "demo-pv" {
  projectid = data.swifdog_project.test.id
  name      = "demo-pv"
}
```

## Argument Reference

- `projectid` - (Required, string) The id of the project to search in.
- `name` - (Required, string) The name of the persistent volume to fetch.
