#!/bin/bash
# Regenerate the SDK from the output of the "Generate SDKs" workflow in https://github.com/1080network/proto.
#
# Usage: scripts/update-from-proto.sh <version> <generated-dir>
#   <version>        release tag the code was generated from: vX.Y.Z or vX.Y.Z-preN
#   <generated-dir>  the generated output (the artifact's `go` directory, i.e. gen/go)
#
# Download the generated code for a release with:
#   gh run download <run-id> -R 1080network/proto -n generated-sdks-<version>
set -euo pipefail

VERSION=${1:?usage: $0 <version> <generated-dir>}
GEN=${2:?usage: $0 <version> <generated-dir>}
ROOT=$(cd "$(dirname "$0")/.." && pwd)

if [[ ! "$VERSION" =~ ^v[0-9]+\.[0-9]+\.[0-9]+(-pre[0-9]+)?$ ]]; then
    echo "version '$VERSION' must look like v1.2.3 or v1.2.3-pre0" >&2
    exit 1
fi
if [[ ! -d "$GEN/mica/issuer" || ! -d "$GEN/micashared" ]]; then
    echo "'$GEN' does not look like generated Go output (expected mica/issuer and micashared)" >&2
    exit 1
fi

echo "Laying out $VERSION from $GEN"
# Generated code only: hand-written code (mtls/, issuer.go, acceptor.go, examples/) is left alone.
rm -rf "$ROOT/common/mica" "$ROOT/common/micashared" "$ROOT/issuer/mica" "$ROOT/acceptor/mica"
mkdir -p "$ROOT/common/mica" "$ROOT/issuer/mica" "$ROOT/acceptor/mica"
cp -r "$GEN/micashared" "$ROOT/common/"
cp -r "$GEN/mica/common" "$GEN/mica/member" "$ROOT/common/mica/"
cp -r "$GEN/mica/issuer" "$ROOT/issuer/mica/"
cp -r "$GEN/mica/acceptor" "$ROOT/acceptor/mica/"

# Minimum runtime versions the generated code needs (protoc-gen-go v1.36.12, protoc-gen-go-grpc v1.6.2).
# Pinned deliberately low so clients are not forced onto a newer Go toolchain than they need; raise these
# when the generator output requires it, and keep the `go` line of each go.mod as low as it can be.
GRPC_VERSION=v1.71.0
PROTOBUF_VERSION=v1.36.12
# The generated code only blank-imports protoc-gen-validate, so any version works; use an old one.
VALIDATE_VERSION=v1.2.1
# protoc-gen-validate pulls in a pre-split genproto, which overlaps google.golang.org/genproto/googleapis/rpc
GENPROTO_VERSION=v0.0.0-20250115164207-1a7da9e5054f
GO_MIN=1.22

# Resolve third-party dependencies first. The role modules' requirement on common is added afterwards,
# because that version isn't tagged yet; the workspace (go.work) resolves common locally for the build.
for m in common issuer acceptor examples; do
    (cd "$ROOT/$m" && go mod edit -droprequire=github.com/1080network/golang/common/v2 -go=$GO_MIN \
        && go get google.golang.org/grpc@$GRPC_VERSION google.golang.org/protobuf@$PROTOBUF_VERSION \
                  github.com/envoyproxy/protoc-gen-validate@$VALIDATE_VERSION \
                  google.golang.org/genproto@$GENPROTO_VERSION >/dev/null)
done
(cd "$ROOT" && go work use ./common ./issuer ./acceptor ./examples)
# Build and test in the workspace, which resolves common locally. This has to happen before the role modules
# require common at $VERSION, because that tag doesn't exist yet and a requirement on it would make go look
# it up remotely.
for m in issuer acceptor examples; do
    (cd "$ROOT/$m" && go mod edit -droprequire=github.com/1080network/golang/common/v2)
done
(cd "$ROOT" && go build ./common/... ./issuer/... ./acceptor/... ./examples/... && go test ./common/...)

# The role modules depend on common at the same version
for m in issuer acceptor examples; do
    (cd "$ROOT/$m" && go mod edit -require="github.com/1080network/golang/common/v2@$VERSION")
done
echo "$VERSION" > "$ROOT/VERSION"
echo "Done. Tag each module as <module>/${VERSION}, for example common/${VERSION}."
