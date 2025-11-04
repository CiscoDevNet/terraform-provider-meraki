[![Tests](https://github.com/CiscoDevNet/terraform-provider-meraki/actions/workflows/test.yml/badge.svg)](https://github.com/CiscoDevNet/terraform-provider-meraki/actions/workflows/test.yml)

# Terraform Provider Meraki

The Meraki provider provides resources to interact with a Cisco Meraki Dashboard. It communicates with Meraki Dashboard via the REST API.

All resources and data sources have been tested with the following API releases.

| API | Version |
| ---------------- | ------- |
| Meraki Dashboard | 1.63.0  |

Documentation: <https://registry.terraform.io/providers/CiscoDevNet/meraki/latest>

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.24

## Support

As this is a community-driven project, support is provided by the community. If you encounter issues or have questions, please use the following resources:

- **GitHub Issues**: Report bugs or request features on our [GitHub Issues](https://github.com/CiscoDevNet/terraform-provider-meraki/issues) page.
- **Discussion Forums**: Engage with other users and contributors on [GitHub Discussions](https://github.com/CiscoDevNet/terraform-provider-meraki/discussions).

## Contributing

We welcome contributions from the community! If you'd like to contribute, please follow our [contribution guidelines](https://github.com/CiscoDevNet/terraform-provider-meraki/blob/main/CONTRIBUTING.md). Whether it's reporting bugs, suggesting features, or submitting pull requests, your involvement helps improve the project for everyone.

## Building The Provider

1. Clone the repository
2. Enter the repository directory
3. Build the provider using the Go `install` command:

```shell
go install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```shell
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

This Terraform Provider is available to install automatically via `terraform init`. If you're building the provider, follow the instructions to
[install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin)
After placing it into your plugins directory,  run `terraform init` to initialize it.

Additional documentation, including available resources and their arguments/attributes can be found on the [Terraform documentation website](https://registry.terraform.io/providers/CiscoDevNet/meraki/latest/docs).

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

More information about how the code generation works can be found in the [contribution guide](https://github.com/CiscoDevNet/terraform-provider-meraki/blob/main/CONTRIBUTING.md).

## Acceptance tests

In order to run the full suite of acceptance tests, run `make testall`. Make sure the respective environment variables are set (e.g., `MERAKI_API_KEY`).

Note: Acceptance tests create real resources.

```shell
make testall
```

More information about how the acceptance tests work can be found in the [contribution guide](https://github.com/CiscoDevNet/terraform-provider-meraki/blob/main/CONTRIBUTING.md).
