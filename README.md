# GOAPI-EXAMPLE

## Table of Contents
* [Summary](#summary)
* [Demo](#demo)
* [Please refer to GOAPI](#please-refer-to-goapi)

## Summary

This golang project is an example which uses the
[goapi](https://github.com/dcwangmit01/goapi) library to implement an API
server.

The goapi library is a framework for implementing an API server.  The
statically-linked all-in-one CLI binary built by this project may run as either
the client or the server, depending on the command line arguments.

The API server responds to both GRPC and JSON REST requests.  The framework
also includes an `/auth` endpoint which responds with JWT auth tokens,
middleware that enforces authentication, a user database, SSL encryption, and
other standard requirements for any API server.

## Demo

![Animated Image of Terminal](https://github.com/dcwangmit01/goapi-example/raw/master/demo/demo.gif)

## Please refer to GOAPI

[https://github.com/dcwangmit01/goapi](https://github.com/dcwangmit01/goapi)
