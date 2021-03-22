# Go

## Introduction
This repository contains the generated Go code from the 1080 Proto
(see: [proto repository](https://github.com/1080network/proto)) files
along with a small sample for showing how to use it.

## Getting Started

### Build
A sample main file is provided that makes Ping, Authenticate, and CreateServiceAccount calls. 

To build run:
```shell script
go build
```

### Test connectivity by calling Ping
The following two commands run the ten80 application and ping either the
partner or sp service:
```shell script
./ten80 partner ping -endpoint grpc.5m9gg.demo.1080.network:443
./ten80 sp ping -endpoint grpc.uparn.demo.1080.network:443
```

### To see how to do authentication and authorization
The command's usage text provides a step-by-step walkthrough of setting up your own jwt private key credential.

To print usage information run:
```shell script
./ten80 -h
```