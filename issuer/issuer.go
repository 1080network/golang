// Package issuer is the Mica SDK for issuers.
//
// The generated gRPC clients and messages live under mica/issuer; this file adds constructors for the
// clients an issuer uses. Connect with the mtls package from github.com/1080network/golang/common/v2.
package issuer

import (
	servicev2 "github.com/1080network/golang/issuer/v2/mica/issuer/service/v2"
	"google.golang.org/grpc"
)

// NewToMicaClient returns the client for calls an issuer makes to Mica.
func NewToMicaClient(cc grpc.ClientConnInterface) servicev2.IssuerToMicaServiceClient {
	return servicev2.NewIssuerToMicaServiceClient(cc)
}
