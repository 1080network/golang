package partner

import (
	partner "github.com/1080network/golang/partner/proto/mica/partner/servicev1"
	"google.golang.org/grpc"
)

var JWTTestRolePartner = []string{"RolePartnerAdmin"}

func BuildPartnerToMicaClient(cc grpc.ClientConnInterface) partner.PartnerToMicaServiceClient {
	return partner.NewPartnerToMicaServiceClient(cc)
}
