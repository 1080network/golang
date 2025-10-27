# Go

## Introduction

This repository contains the Mica Network SDK for golang.
The APIs use [gRPC](https://grpc.io) to communicate to or listen from the Mica Network. Most of the code in this repository
is generated from the [proto repository](https://github.com/1080network/proto)

## Cloning this repo

This repository uses a git submodule to get the proto definition code. You can clone with the following command to get the submodule initialized 

```shell
git clone --recurse-submodules git@github.com:1080network/golang.git
```

## Manual release

As of now there is no automation to release this SDK so it is a manual process.

### One time setup

To build this code you need to install the pre-requisites listed [here](https://grpc.io/docs/languages/go/quickstart/#prerequisites) and have gnu Make.

### Release

1. Set VERSION environment variable: `export VERSION=vX.Y.Z` (`VERSION` is whatever the protos tag is from which this SDK is getting released)
2. `git checkout -b $VERSION`
3. `make all` to build the new SDK
4. Create PR with changes, get approval, merge into main
5. From main branch: `make publish`
6. Verify that the new release is available [here](https://github.com/1080network/golang/tags).

Note that this process assumes that the version of the SDK matches the version of the protos which is the recommended approach.

### Code examples

Check the examples directory for references on how to use the SDK
