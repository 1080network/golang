package connect

import (
	connect "github.com/1080network/golang/connect/proto/mica/connect/servicev1"
	"google.golang.org/grpc"
)

var JWTTestRoleConnect = []string{"RoleConnectAdmin"}

func BuildConnectClient(cc grpc.ClientConnInterface) connect.ConnectServiceClient {
	return connect.NewConnectServiceClient(cc)
}
