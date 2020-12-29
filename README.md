# Go API client for RingCentral

[![Build Status][build-status-svg]][build-status-link]
[![Go Report Card][goreport-svg]][goreport-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]
[![Stack Overflow][stackoverflow-svg]][stackoverflow-link]
[![Twitter][twitter-svg]][twitter-link]

 [build-status-svg]: https://github.com/grokify/go-ringcentral/workflows/build/badge.svg
 [build-status-link]: https://github.com/grokify/go-ringcentral/actions
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-ringcentral
 [goreport-link]: https://goreportcard.com/report/github.com/grokify/go-ringcentral
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/go-ringcentral
 [docs-godoc-link]: https://pkg.go.dev/github.com/grokify/go-ringcentral
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/go-ringcentral/blob/master/LICENSE
 [stackoverflow-svg]: https://img.shields.io/badge/Stack%20Overflow-ringcentral-orange.svg
 [stackoverflow-link]: https://stackoverflow.com/questions/tagged/ringcentral
 [twitter-svg]: https://img.shields.io/twitter/follow/ringcentraldevs.svg?style=social&label=follow
 [twitter-link]: https://twitter.com/RingCentralDevs

## Overview

This is a API client built using [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) using this RingCentral API spec [`swagger_spec.yaml`](codegen/swagger_spec.yaml).

This API client does not include auth and relies on [`oauth2more/ringcentral`](https://github.com/grokify/oauth2more/tree/master/ringcentral) to create an `*http.Client` to use as a dependency injection.

## Installation

`$ go get github.com/grokify/go-ringcentral/...`

## Usage

See examples in the [`examples`](examples) directory. To get started, you can use [`examples/get_me/get_me.go`](examples/get_me/get_me.go).

By default, these use a `.env` file which can be specified by the `ENV_PATH` environment variable or a local `./.env` file. Using one `.env` file specified by `ENV_PATH` is useful so you only need one file to run all the examples.

## Documentation

The auto-generated Swagger files are in the `client` folder and you can find the Swagger docs there as `client/README.md`:

* Office: [`office/v1/client`](office/v1/client), [`office/v1/client/README.md`](office/v1/client/README.md)
* Engage Digital: [`engagedigital/v1/client`](engagedigital/v1/client), [`engagedigital/v1/client/README.md`](engagedigital/v1/client/README.md)
* Engage Voice: [`engagevoice/v1/client`](engagevoice/v1/client), [`engagevoice/v1/client/README.md`](engagevoice/v1/client/README.md)

## Related Packages

In addition to [`oauth2more/ringcentral`](https://github.com/grokify/oauth2more/tree/master/ringcentral), the following project provides a website with OAuth 2.0 authorization code flow example:

[`grokify/beego-oauth2-demo`](https://github.com/grokify/beego-oauth2-demo)
