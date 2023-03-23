# Terraform Provider GraphQL Schema

TODO introduction

## Build provider

Run the following command to build the provider

```shell
$ go build -o terraform-provider-gqlschema
```

## Test sample configuration

First, build and install the provider.

```shell
$ make install
```

Then, navigate to the `examples` directory. 

```shell
$ cd examples/provider-install-verification
```

Run the following command to initialize the workspace and apply the sample configuration.

```shell
$ terraform init && terraform apply
```
