# Floating IP Resources

Sometimes you want to expose more than "just" websites like gameservers, internal backend applications or file access
protocols like sftp. In this case, you need a solution to expose other ports than 80 or 443, what you can't do with our
Ingress Resources which only support unsecure (80) and secure (443) HTTP protocol.

A floating IP resource provides you a dedicated IP which you can use to expose a service like mentioned above. You can
expose multiple ports using only one IP, but we are limiting the ports you are allowed to use to prevent RFC violations.
An information about what you are not allowed to host on our infrastructure can be found in our terms of service on our
website.

Please note that you'll need both a v4 and v6 floating ip resource if you want to expose your endpoints on both ip
versions. If you only want to expose it either through v4 or v6, you only need one!

## Example

In the following example we want to expose a Packet which provides a FTP server over a dedicated ip v4 address.

```terraform
resource "swifdog_floating_ip" "my-dedicated-ip" {
  projectid = data.swifdog_project.test.id
  version   = 4

  endpoint {
    packetid      = swifdog_packet.sftp-packet.id
    containerport = 21
    targetport    = 21
  }
}
```

If you also want to expose your web packet (or anything else) on the same dedicated IP, you are able to add as lot
endpoints as you want like this inside your swifdog_floating_ip resource:

```terraform
  ...

  endpoint {
    packetid = swifdog_packet.web-packet.id
    containerport = 80
    targetport = 80
  }


  endpoint {
    packetid = swifdog_packet.web-packet.id
    containerport = 443
    targetport = 443
  }

...
```

Please notice, that you are NOT able to expose a targetport multiple times within a floating ip.

## Argument Reference

- `projectid` - (Required, string) The ID of the project where this resource should be in.
- `version` - (Optional, integer) The version of the IP address to be allocated like 4 or 6.
- `endpoint` - (Required, array) A list of packets to be exposed on the targetPort of the floating ip address 