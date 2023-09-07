package example

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/1080network/golang/partner"
	"github.com/1080network/golang/partner/proto/mica/partner/organizationv1"
	partnersvc "github.com/1080network/golang/partner/proto/mica/partner/servicev1"
	"github.com/1080network/golang/partner/proto/micashared/common/enums/organizationcategoryv1"
	"github.com/1080network/golang/partner/proto/micashared/common/pingv1"
	"github.com/1080network/golang/shared"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"os"
	"testing"
)

func initializeClient() (partnersvc.PartnerToMicaServiceClient, *grpc.ClientConn, error) {
	certificateAuthorities, err := loadPemBlob("ROOT_CA_PEM")
	if err != nil {
		return nil, nil, err
	}
	clientSecretKey, err := loadPemBlob("CLIENT_KEY_PEM")
	if err != nil {
		return nil, nil, err
	}
	clientCertificate, err := loadPemBlob("CLIENT_CERTIFICATE_PEM")
	if err != nil {
		return nil, nil, err
	}
	conn, err := shared.SecureClientDial("api.hron3n.members.mica.io:443", certificateAuthorities, clientCertificate, clientSecretKey)
	if err != nil {
		return nil, nil, err
	}
	partnerClient := partner.BuildPartnerToMicaClient(conn)
	return partnerClient, conn, nil
}

func TestPing(t *testing.T) {
	client, conn, err := initializeClient()
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	assert.NoError(t, err)
	response, err := client.Ping(context.TODO(), &pingv1.PingRequest{})
	assert.NoError(t, err)
	assert.Equal(t, pingv1.PingResponse_STATUS_SUCCESS, response.Status)
}

func TestSearchOrganization(t *testing.T) {
	ctx := context.TODO()
	client, conn, err := initializeClient()
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	assert.NoError(t, err)
	request := &organizationv1.SearchOrganizationRequest{
		Category: organizationcategoryv1.OrganizationCategory_ORGANIZATION_CATEGORY_GROCERY,
	}
	response, err := client.SearchOrganization(ctx, request)
	assert.NoError(t, err)
	assert.Equal(t, organizationv1.SearchOrganizationResponse_STATUS_SUCCESS, response.Status)
	assert.Greater(t, len(response.Organizations), 0, "expected at least one organization")
	for _, organization := range response.Organizations {
		fmt.Println(fmt.Sprintf("%v", organization))
	}
}

func TestGetOrganization(t *testing.T) {
	ctx := context.TODO()
	micaClient, conn, err := initializeClient()
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	assert.NoError(t, err)
	request := &organizationv1.GetOrganizationRequest{
		OrganizationKey: "hron3n00MDV90UjuTAi71blY95IkwA",
	}
	response, err := micaClient.GetOrganization(ctx, request)
	assert.NoError(t, err)
	assert.Equal(t, organizationv1.GetOrganizationResponse_STATUS_SUCCESS, response.Status)
	fmt.Println(fmt.Sprintf("%v", response.Organization))
}

func TestUpdateOrganization(t *testing.T) {
	ctx := context.TODO()
	micaClient, conn, err := initializeClient()
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	assert.NoError(t, err)
	request := &organizationv1.GetOrganizationRequest{
		OrganizationKey: "hron3n00MDV90UjuTAi71blY95IkwA",
	}
	response, err := micaClient.GetOrganization(ctx, request)
	assert.NoError(t, err)
	assert.Equal(t, organizationv1.GetOrganizationResponse_STATUS_SUCCESS, response.Status)

	//updateRequest := &organizationv1.UpdateOrganizationRequest{
	//	OrganizationKey:               response.Organization.OrganizationKey,
	//	Version:                       response.Organization.Version,
	//	Name:                          "",
	//	Address:                       response.Organization.Address,
	//	DomesticAchRoutingNumber:      response.Organization.DomesticAchRoutingNumber,
	//	InternationalAchRoutingNumber: response.Organization.InternationalAchRoutingNumber,
	//	WireRoutingNumber:             response.Organization.WireRoutingNumber,
	//	SwiftRoutingNumber:            response.Organization.SwiftRoutingNumber,
	//	BankAccountNumber:             "",
	//}
}

func loadPemBlob(envVar string) ([]byte, error) {
	//expect the pem blob to be base64 encoded
	value, ok := os.LookupEnv(envVar)
	if !ok {
		return nil, errors.New(fmt.Sprintf("environment varible %s is not set", envVar))
	}
	return base64.StdEncoding.DecodeString(value)
}
