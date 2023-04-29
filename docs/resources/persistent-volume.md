# Persistent Volume Resource

All data within a container and its lifecycle are not persistent. To persist data across the lifecycle of a single container and share data between multiple packets, you can use persistent volumes (PVs).

To create a PV, you can use this Terraform code:

```terraform
resource "swifdog_persistent_volume" "example-pv" {
  projectid = swifdog_project.test.id
  name      = "demo"
  capacity  = "5GB"
}
```

A PV always belongs to a project. Therefore, you must provide a project ID. You can reference a project (using data/resource, as shown above) or provide a static string. We recommend referencing the project. The name is a unique identifier within a project and required. Capacity describes the size of the reserved storage space and is also required. Currently, only sizes in gigabytes are allowed in the storage plan. If necessary, larger hard drives will be possible in the future. Note that the "GB" suffix is required.

## Argument Reference

- `projectid` - (Required, string) The ID of the project where this resource should be in.
- `name` - (Required, string) The name of this persistent volume.
- `capacity` - (Required, string) The maximum capacity as an integer with the size (like "GB") as a suffix.
