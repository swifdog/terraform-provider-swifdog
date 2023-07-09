# Packet Resource

A packet represents a set of containers. To create a packet, you need at least one image and a unique name, as well as a project ID reference.

To create a packet using Terraform, you can use this code:

```terraform
resource "swifdog_packet" "hello-world" {
  projectid = data.swifdog_project.test.id
  name      = "hello-world"
  image     = "nginx"
}
```

You can add environment variables as follows. Note that you need to include the entire code block in the brackets of the resource above.

```terraform
env {
    key = "PMA_ABSOLUTE_URI"
    value = "demo.com/phpmyadmin/"
}
```

By default, you cannot access a packet by its name or ID. You need to release any ports before you can access them by the packet name. For example, with a database container, you can release a port as follows:

```terraform
internalport {
    containerport = 3306
    protocol = "tcp"
}
```

You can integrate the persistent volumes mentioned above into packets as follows:

```terraform
volume {
    volumeid = swifdog_persistent_volume.example-pv.id
    mountpath = "/var/www/demo"
}
```

When referencing volumes, it doesn't matter whether you use the volumename or the volumeid with the corresponding key.

## Argument Reference

- `projectid` - (Required, string) The ID of the project where this resource should be in.
- `name` - (Required, string) The name of a container which must be unique within a project.
- `image` - (Required, string) The container image to be run like "nginx:latest".
- `registryCredentialId` - (Optional, string) The linked image pull secret to pull the given image.
- `env` - (Optional, array) A list of key & values for setting environment variables within the container.
- `internalport` - (Optional, array) A list of ports and their protocols which should be exposed within a project.
- `volume ` - (Optional, array) A list of volumes to be mounted into the container.
