# Ingress Rule Resource

An Ingress Rule allows you to receive web traffic from the internet and distribute it to your packets.

```terraform
resource "swifdog_ingress_rule" "example-ingress" {
  projectid = data.swifdog_project.test.id
  hostname  = "example.com"

  path {
    path          = "/"
    packetid      = swifdog_packet.hello-world.id
    containerport = 80
  }
  path {
    path          = "/blog"
    packetid      = swifdog_packet.wordpress.id
    containerport = 80
  }
}
```

You can enter an unlimited number of paths. Each path is set under the hostname.

For our example, this means the following scheme:

example.com/* -> swifdog_packet.hello-world:80
example.com/blog/* -> swifdog_packet.wordpress:80

An Ingress Rule without paths does not make sense.

## Argument Reference

- `projectid` - (Required, string) The ID of the project where this resource should be in.
- `hostname` - (Required, string) The hostname to be used.
- `path` - (Required, string) A list of path rules.
