package discount

import (
	discount "github.com/1080network/golang/discount/proto/mica/discount/servicev1"
	"google.golang.org/grpc"
)

var JWTTestRoleDiscount = []string{"RoleDiscountAdmin"}

func BuildDiscountToMicaClient(cc grpc.ClientConnInterface) discount.DiscountToMicaServiceClient {
	return discount.NewDiscountToMicaServiceClient(cc)
}
