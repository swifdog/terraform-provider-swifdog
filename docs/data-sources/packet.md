# Packet Data Source

A packet represents a set of containers. To create a packet, you need at least one image and a unique name, as well as a project ID reference.

You can fetch an existing packet with the following code:

```terraform
data "swifdog_packet" "test-packet" {
  projectid = data.swifdog_project.demo-project.id
  name      = "test-packet"
}
```

In this case, the first packet with an equal name to "test-packet" will be fetched and used in further refs of this data.

## Argument Reference

- `projectid` - (Required, string) The id of the project to search in.
- `name` - (Required, string) The name of the packet to fetch.
