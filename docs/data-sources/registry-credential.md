# Registry Credential Data Source

Provide details about a registry credential (e.g. Image Pull Secret). A registry credential is always inside a project,
so please provide the
`projectid`!

```terraform
data "swifdog_registry_credential" "example-rc" {
  projectid   = data.swifdog_project.test.id
  registryUrl = "registry.gitlab.com"
  username    = "max@example.com"
}
```

## Argument Reference

- `projectid` - (Required, string) The id of the project to search in.
- `registryUrl` - (Required, string) The registry url (as you've created it)
- `username` - (Optional, string) You'll have to set it when you have multiple values for one registryUrl inside a given
  project to specify the correct credential.
