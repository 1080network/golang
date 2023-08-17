# Go

## Introduction
This repository contains the Mica Network SDK for golang.
The APIs use [gRPC](https://grpc.io) to communicate to or listen from the Mica Network. Most of the code in this repository
is generated from the [proto repository](https://github.com/1080network/proto)

### Cloning this repo
This repository uses a git submodule to get the proto definition code. You can clone with the following command to get the submodule initialized 

```shell
git clone --recurse-submodules git@github.com:1080network/golang.git
```
### Build
To build this code you need to install the pre-requisites listed [here](https://grpc.io/docs/languages/go/quickstart/#prerequisites) and have gnu Make.

Once you have the pre-requisites just run the following command:

```shell
make all
```

### Code examples
Check the examples directory for references on how to use the SDK