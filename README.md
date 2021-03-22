# Go

## Introduction
This repository contains the generated Go code from the 1080 Proto
(see: [proto repository](https://github.com/1080network/proto)) files
along with a small sample for showing how to use it.

## Getting Started

### Build
A sample main file is provided that makes a Ping call. To build run:
```shell script
go build
```

### Run
The following two commands run the ten80 application and ping either the
partner or sp service:
```shell script
./ten80 partner grpc.4p1nr.demo.1080.network:443
./ten80 sp grpc.uparn.demo.1080.network:443
```

## Using a credential

### Generate a public/private key pair
```shell script
ssh-keygen -t rsa -b 4096 -f key
ssh-keygen -f key.pub -e -m pem > key.pem
```

### Upload the public key to 1080 network