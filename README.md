# Go API client for RingCentral

[![Build Status][build-status-svg]][build-status-link]
[![Go Report Card][goreport-svg]][goreport-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]

 [build-status-svg]: https://api.travis-ci.org/grokify/go-ringcentral.svg?branch=master
 [build-status-link]: https://travis-ci.org/grokify/go-ringcentral
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-ringcentral
 [goreport-link]: https://goreportcard.com/report/github.com/grokify/go-ringcentral
 [docs-godoc-svg]: https://img.shields.io/badge/docs-godoc-blue.svg
 [docs-godoc-link]: https://godoc.org/github.com/grokify/go-ringcentral
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/go-ringcentral/blob/master/LICENSE

## Overview

This is a API client built using [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) using this RingCentral API spec [`swagger_spec.yaml`](codegen/swagger_spec.yaml).

This API client does not include auth and relies on [`oauth2more/ringcentral`](https://github.com/grokify/oauth2more/tree/master/ringcentral) to create an `*http.Client` to use as a dependency injection.

## Installation

`$ go get github.com/grokify/go-ringcentral/...`

## Usage

See examples in the [`examples`](examples) directory. To get started, you can use [`examples/get_me/get_me.go`](examples/get_me/get_me.go).

By default, these use a `.env` file which can be specified by the `ENV_PATH` environment variable or a local `./.env` file. Using one `.env` file specified by `ENV_PATH` is useful so you only need one file to run all the examples.

## Documentation

The auto-generated Swagger files are in the [`client`](client) folder and you can find the Swagger docs there as [`client/README.md`](client/README.md)

## Related Packages

In addition to [`oauth2more/ringcentral`](https://github.com/grokify/oauth2more/tree/master/ringcentral), the following project provides a website with OAuth 2.0 authorization code flow example:

[`grokify/beego-oauth2-demo`](https://github.com/grokify/beego-oauth2-demo)