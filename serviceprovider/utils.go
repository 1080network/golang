package serviceprovider

import (
	serviceprovider "github.com/1080network/golang/serviceprovider/proto/serviceprovider/servicev1"
	"google.golang.org/grpc"
)

var JWTTestRoleServiceProvider = []string{"RoleServiceProviderAdmin"}

func BuildServiceProviderToMicaClient(cc grpc.ClientConnInterface) serviceprovider.ServiceProviderToMicaServiceClient {
	return serviceprovider.NewServiceProviderToMicaServiceClient(cc)
}
