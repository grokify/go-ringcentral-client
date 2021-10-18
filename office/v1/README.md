# Go API client for RingCentral

[![Build Status][build-status-svg]][build-status-link]
[![Go Report Card][goreport-svg]][goreport-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]
[![Stack Overflow][stackoverflow-svg]][stackoverflow-link]
[![Twitter][twitter-svg]][twitter-link]

 [build-status-svg]: https://api.travis-ci.org/grokify/go-ringcentral-client.svg?branch=master
 [build-status-link]: https://travis-ci.org/grokify/go-ringcentral-client
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-ringcentral-client
 [goreport-link]: https://goreportcard.com/report/github.com/grokify/go-ringcentral-client
 [docs-godoc-svg]: https://img.shields.io/badge/docs-godoc-blue.svg
 [docs-godoc-link]: https://godoc.org/github.com/grokify/go-ringcentral-client
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/go-ringcentral-client/blob/master/LICENSE
 [stackoverflow-svg]: https://img.shields.io/badge/Stack%20Overflow-ringcentral-orange.svg
 [stackoverflow-link]: https://stackoverflow.com/questions/tagged/ringcentral
 [twitter-svg]: https://img.shields.io/twitter/follow/ringcentraldevs.svg?style=social&label=follow
 [twitter-link]: https://twitter.com/RingCentralDevs

## Overview

This is a API client built using [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) using this RingCentral API spec [`swagger_spec.yaml`](codegen/swagger_spec.yaml).

This API client does not include auth and relies on [`goauth/ringcentral`](https://github.com/grokify/goauth/tree/master/ringcentral) to create an `*http.Client` to use as a dependency injection.

## Installation

`$ go get github.com/grokify/go-ringcentral-client/...`

## Usage

See examples in the [`examples`](examples) directory. To get started, you can use [`examples/get_me/get_me.go`](examples/get_me/get_me.go).

By default, these use a `.env` file which can be specified by the `ENV_PATH` environment variable or a local `./.env` file. Using one `.env` file specified by `ENV_PATH` is useful so you only need one file to run all the examples.

## Documentation

The auto-generated Swagger files are in the [`client`](client) folder and you can find the Swagger docs there as [`client/README.md`](client/README.md)

## Related Packages

In addition to [`goauth/ringcentral`](https://github.com/grokify/goauth/tree/master/ringcentral), the following project provides a website with OAuth 2.0 authorization code flow example:

[`grokify/beegoutil`](https://github.com/grokify/beegoutil)
