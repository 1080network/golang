// Package acceptor is the Mica SDK for acceptors.
//
// The generated gRPC clients and messages live under mica/acceptor; this file adds constructors for the
// clients an acceptor uses. Connect with the mtls package from github.com/1080network/golang/common/v2.
package acceptor

import (
	servicev2 "github.com/1080network/golang/acceptor/v2/mica/acceptor/service/v2"
	"google.golang.org/grpc"
)

// NewToMicaClient returns the client for calls an acceptor makes to Mica.
func NewToMicaClient(cc grpc.ClientConnInterface) servicev2.AcceptorToMicaServiceClient {
	return servicev2.NewAcceptorToMicaServiceClient(cc)
}
