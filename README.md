# Mica SDK for Go

Go client libraries for the [Mica](https://mica.io) network. Mica's API is [gRPC](https://grpc.io) over mutual TLS;
the messages and clients in this repository are generated from the
[public protos](https://github.com/1080network/proto).

## Which module do I need?

Each role is its own Go module, so you only depend on the code for the roles you play.

| Module | Use it if you are | Import path |
| --- | --- | --- |
| `issuer` | an issuer | `github.com/1080network/golang/issuer/v2` |
| `acceptor` | an acceptor | `github.com/1080network/golang/acceptor/v2` |
| `common` | everyone (shared messages, the `mtls` helper); pulled in automatically by the role modules | `github.com/1080network/golang/common/v2` |

```bash
go get github.com/1080network/golang/issuer/v2   # or acceptor/v2
```

Requires Go 1.23 or later.

## Connecting

Mica uses mutual TLS. You need the root CA, and the client certificate and key Mica issued to you.

```go
cfg, err := mtls.ConfigFromFiles("rootca.crt", "client.crt", "client.key")
conn, err := mtls.Dial("api.example.mica.io:443", cfg)
defer conn.Close()

issuerClient := issuer.NewToMicaClient(conn) // or acceptor.NewToMicaClient
```

Every generated client is also available directly, for example `pingv1.NewPingServiceClient(conn)` in
`github.com/1080network/golang/common/v2/mica/member/ping/v1`.

`examples/ping` is a complete program that connects and calls `Ping`, a good first check that your certificates
and network path work:

```bash
go run ./examples/ping -target api.example.mica.io:443 -ca rootca.crt -cert client.crt -key client.key
```

## Versions

SDK versions match the proto release they were generated from (`v2.0.0-pre0`, `v2.0.0`, and so on). Because each
module is versioned separately, git tags are prefixed with the module name, for example `issuer/v2.0.0`.
Version 1 of this SDK (the `serviceprovider`, `partner`, `discount` and `fullsdk` modules) remains available at its
existing `v1.x` tags and is not updated.

## Repository layout

| Path | Contents |
| --- | --- |
| `common/`, `issuer/`, `acceptor/` | one Go module each; `mica/` and `micashared/` inside are **generated** |
| `common/mtls/` | hand-written mutual TLS dialer |
| `issuer/issuer.go`, `acceptor/acceptor.go` | hand-written client constructors |
| `examples/` | runnable examples |
| `scripts/update-from-proto.sh` | regenerates the SDK for a proto release |
| `go.work` | lets the modules be developed together locally |

## Releasing (maintainers)

1. Run the **Generate SDKs** workflow in [1080network/proto](https://github.com/1080network/proto/actions) for the release tag.
2. Download its output: `gh run download <run-id> -R 1080network/proto -n generated-sdks-<version>`.
3. Run `scripts/update-from-proto.sh <version> <download-dir>/go`. It lays the generated code into the modules, builds everything and runs the tests.
4. Open a pull request with the changes and merge it.
5. Tag the merge commit once per module, `common` first: `common/<version>`, `issuer/<version>`, `acceptor/<version>`, and push the tags.
6. Check that a clean module can use it: `GOWORK=off go get github.com/1080network/golang/issuer/v2@<version>`.
