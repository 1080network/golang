module github.com/1080network/golang/examples

go 1.19

replace (
	github.com/1080network/golang/partner => ../partner
	github.com/1080network/golang/serviceprovider => ../serviceprovider
	github.com/1080network/golang/shared => ../shared
)

require (
	github.com/1080network/golang/partner v1.0.0
	github.com/1080network/golang/serviceprovider v1.0.0
	github.com/1080network/golang/shared v1.0.0
	github.com/stretchr/testify v1.8.1
)

require (
	github.com/cockroachdb/apd/v3 v3.2.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.1.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.1.0 // indirect
	github.com/goccy/go-json v0.9.11 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/lestrrat-go/blackmagic v1.0.1 // indirect
	github.com/lestrrat-go/httpcc v1.0.1 // indirect
	github.com/lestrrat-go/httprc v1.0.4 // indirect
	github.com/lestrrat-go/iter v1.0.2 // indirect
	github.com/lestrrat-go/jwx/v2 v2.0.6 // indirect
	github.com/lestrrat-go/option v1.0.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20220427172511-eb4f295cb31f // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.49.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
