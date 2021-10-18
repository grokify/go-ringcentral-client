# Go API client for RingCentral

[![Build Status][build-status-svg]][build-status-link]
[![Go Report Card][goreport-svg]][goreport-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]
[![Stack Overflow][stackoverflow-svg]][stackoverflow-link]
[![Twitter][twitter-svg]][twitter-link]

 [build-status-svg]: https://github.com/grokify/go-ringcentral-client/workflows/build/badge.svg
 [build-status-link]: https://github.com/grokify/go-ringcentral-client/actions
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-ringcentral-client
 [goreport-link]: https://goreportcard.com/report/github.com/grokify/go-ringcentral-client
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/go-ringcentral-client
 [docs-godoc-link]: https://pkg.go.dev/github.com/grokify/go-ringcentral-client
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/go-ringcentral-client/blob/master/LICENSE
 [stackoverflow-svg]: https://img.shields.io/badge/Stack%20Overflow-ringcentral-orange.svg
 [stackoverflow-link]: https://stackoverflow.com/questions/tagged/ringcentral
 [twitter-svg]: https://img.shields.io/twitter/follow/ringcentraldevs.svg?style=social&label=follow
 [twitter-link]: https://twitter.com/RingCentralDevs

## Overview

<div align="center">

:bangbang: This module has been renamed to `go-ringcentral-client` from `go-ringcentral` to better work with Go modules and its new versioning requirements. Old versions are still available via `proxy.golang.org` :bangbang:

</div>

This module provides RingCentral API clients for RingCentral Office, Engage Digital and Engage Voice.

The clients are built using [OpenAPI Generator 4.x](https://github.com/OpenAPITools/openapi-generator) using this RingCentral API spec [`swagger_spec.yaml`](codegen/swagger_spec.yaml).

This API client does not include auth and relies on [`goauth/ringcentral`](https://github.com/grokify/goauth/tree/master/ringcentral) to create an `*http.Client` to use as a dependency injection.

## Installation

`$ go get github.com/grokify/go-ringcentral-client/...`

## Usage

See Office examples in the [`office/v1/examples`](office/v1/examples) directory. To get started, you can use [`office/v1/examples/get_me/get_me.go`](office/v1/examples/get_me/get_me.go).

By default, these use a `.env` file which can be specified by the `ENV_PATH` environment variable or a local `./.env` file. Using one `.env` file specified by `ENV_PATH` is useful so you only need one file to run all the examples.

## Documentation

The auto-generated Swagger files are in the `client` folder and you can find the Swagger docs there as `client/README.md`:

* Office: [`office/v1/client`](office/v1/client), [`office/v1/client/README.md`](office/v1/client/README.md)
* Engage Digital: [`engagedigital/v1/client`](engagedigital/v1/client), [`engagedigital/v1/client/README.md`](engagedigital/v1/client/README.md)
* Engage Voice: [`engagevoice/v1/client`](engagevoice/v1/client), [`engagevoice/v1/client/README.md`](engagevoice/v1/client/README.md)

## Related Packages

### GoAuth RingCentral

For RingCentral auth haandling see [`goauth/ringcentral`](https://github.com/grokify/goauth/tree/master/ringcentral). This is specifically a package to retrieve a `*http.Client` or `*oauth2.Token` without needing the auto-generated models.

### Beego Example

In addition to [`goauth/ringcentral`](https://github.com/grokify/goauth/tree/master/ringcentral), the following project provides a website with OAuth 2.0 authorization code flow example:

[`grokify/beegoutil`](https://github.com/grokify/beegoutil)
