package example

import (
	"github.com/1080network/golang/partner"
	partnersvc "github.com/1080network/golang/partner/proto/partner/servicev1"
	"github.com/1080network/golang/shared"
)

func createTestClient() (partnersvc.PartnerToMicaServiceClient, error){
	credential := shared.BuildJWTPerRpcCredential(partner.JWTTestRolePartner)
	conn, err := shared.TestClientDial("localhost:14000", credential)
	if err != nil{
		return nil, err
	}
	partnerClient := partner.BuildPartnerToMicaClient(conn)
	return partnerClient, nil
}