Terraform Provider Okta
==================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Maintainers
-----------

This provider plugin is maintained by the Terraform team at [Articulate](https://articulate.com/).

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.11.x
-	[Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

Usage
---------------------

This plugin requires two inputs to run: the Okta organization name and the Okta API token. The Okta base url is not required and will default to "okta.com" if left out.

You can specify the inputs in your tf plan:

```
provider "okta" {
  org_name  = <okta instance name, e.g. dev-XXXXXX>
  api_token = <okta instance api token with the Administrator role>
  base_url  = <okta base url, e.g. oktapreview.com>
}
```

_or_ you can specify environment variables:

```
OKTA_ORG_NAME=<okta instance name, e.g. dev-XXXXXX>
OKTA_API_TOKEN=<okta instance api token with the Administrator role>
OKTA_BASE_URL=<okta base url, e.g. oktapreview.com>
```

Examples
--------

As we build out resources, we build concomitant acceptance tests that require use to create resource config that actually creates and modifies real resources. We decided to put these test fixtures to good use and provide them [as examples here.](./examples)

Building The Provider
---------------------

Simply clone the provider and run `make`. The `make` command combines `make deps && make build-plugins`, which ensures all dependencies are pulled down, builds binaries for Linux, OSX, and Windows and drops them in the default unix terraform plugin directory `~/.terraform.d/plugins`. Use `make build` for only a single binary for your host OS in your current directory. Note, we use `dep` for dependency management, this is triggered with `make deps`.

```sh
$ go get -d github.com/articulate/terraform-provider-okta
$ cd $GOPATH/src/github.com/articulate/terraform-provider-okta
$ make
```

For local development, we've found the commands below to be helpful. Run them from inside the `terraform-provider-okta` directory:

```sh
$ go build -o .terraform/plugins/$GOOS_$GOARCH/terraform-provider-okta
$ terraform init -plugin-dir=.terraform/plugins/$GOOS_$GOARCH
```

Using the provider
----------------------

Example terraform plan:

```
provider "okta" {
  org_name  = "dev-XXXXX"
  api_token = "XXXXXXXXXXXXXXXXXXXXXXXX"
  base_url  = "oktapreview.com"
}

resource "okta_user" "blah" {
  firstname = "blah"
  lastname  = "blergh"
  email     = "XXXXX@XXXXXXXX.XXX"
}
```

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-okta
...
```

In order to test the provider, you can simply run `make test`. The acceptance tests require an API token and a corresponding Okta org, if you use dotenv, you can `cp .env.sample .env` and add your Okta settings there, and prefix make test with `dotenv`.

```sh
$ make test
```

In order to run the full suite of acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

### Best Practices

We are striving to build a provider that is easily consumable and eventually can pass the HashiCorp community audit. In order to achieve this end we must ensure we are following HashiCorp's best practices. This can be derived either from their [documentation on the matter](https://www.terraform.io/docs/extend/best-practices/detecting-drift.html), or by using a simple well written [example as our template](https://github.com/terraform-providers/terraform-provider-datadog).
