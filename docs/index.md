# Swifdog.io Provider

Swifdog.io provides you the infrastructure to securely deploy and manage your docker applications. This project helps
you to maintain the resources using Terraform.

At the moment we're supporting authentication through your basic credentials and custom authentication tokens. Please
create an account on https://www.swifdog.io/ first.

## Example usage

Integrate the provider into your project:

```terraform
provider "swifdog" {
  email    = "max@mustermann.de"
  password = "test"
}
```

## Argument Reference

The following arguments are supported:

- `email` - (Required for BASIC authentication, string) The email address of your account. You can pass it using the env
  variable `SWIFDOG_EMAIL`as well.
- `password` - (Required for BASIC authentication, string) The password of your account. You can pass it using the env
  variable `SWIFDOG_PASSWORD`as well.
- `api_token` - (Required for TOKEN authentication, string) A generated API token. You can pass it using the env variable `SWIFDOG_API_TOKEN`as well.
