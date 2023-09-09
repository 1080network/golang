package example

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/1080network/golang/partner"
	"github.com/1080network/golang/partner/proto/mica/partner/organizationv1"
	partnersvc "github.com/1080network/golang/partner/proto/mica/partner/servicev1"
	"github.com/1080network/golang/partner/proto/mica/partner/valuev1"
	"github.com/1080network/golang/partner/proto/micashared/common/enums/approvaltypev1"
	"github.com/1080network/golang/partner/proto/micashared/common/enums/barcodelocationv1"
	"github.com/1080network/golang/partner/proto/micashared/common/enums/barcodetypev1"
	"github.com/1080network/golang/partner/proto/micashared/common/enums/currencyv1"
	"github.com/1080network/golang/partner/proto/micashared/common/enums/organizationcategoryv1"
	"github.com/1080network/golang/partner/proto/micashared/common/enums/unitv1"
	"github.com/1080network/golang/partner/proto/micashared/common/pingv1"
	commonv1 "github.com/1080network/golang/partner/proto/micashared/common/v1"
	"github.com/1080network/golang/shared"
	"github.com/google/uuid"
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

	updateRequest := &organizationv1.UpdateOrganizationRequest{
		OrganizationKey:  response.Organization.OrganizationKey,
		Version:          response.Organization.Version,
		Name:             response.Organization.Name + "updated",
		Address:          response.Organization.Address,
		OperatingAccount: response.Organization.OperatingAccount,
		RevenueAccount:   response.Organization.RevenueAccount,
		BarcodeType:      barcodetypev1.BarcodeType_BARCODE_TYPE_EAN128,
		BarcodeLocation:  barcodelocationv1.BarcodeLocation_BARCODE_LOCATION_TOP,
	}
	updateResponse, err := micaClient.UpdateOrganization(ctx, updateRequest)
	assert.NoError(t, err)
	assert.Equal(t, organizationv1.UpdateOrganizationResponse_STATUS_SUCCESS, updateResponse.Status)
}

func TestObtainValueWithBasket(t *testing.T) {
	ctx := context.TODO()
	micaClient, conn, err := initializeClient()
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	lineItems := []*commonv1.LineItem{
		{
			Sequence:      1,
			LineItemGroup: "Boxed Items",
			ProductCode:   "021000658930",
			Description:   "Velveta Cheese",
			Quantity:      "1",
			Unit:          unitv1.Unit_UNIT_ITEM,
			UnitAmount:    "3.79",
			UnitTaxAmount: "0",
			LineAmount:    "3.79",
			LineTaxAmount: "0",
		},
		{
			Sequence:      2,
			LineItemGroup: "Tobacco",
			ProductCode:   "028200003034",
			Description:   "Marlboro Red 20 pack",
			Quantity:      "1",
			Unit:          unitv1.Unit_UNIT_ITEM,
			UnitAmount:    "22.50",
			UnitTaxAmount: "1.912",
			LineAmount:    "22.50",
			LineTaxAmount: "1.912",
		},
		{
			Sequence:      3,
			LineItemGroup: "Produce",
			ProductCode:   "90099332322",
			Description:   "Organic Bananas",
			Quantity:      "1",
			Unit:          unitv1.Unit_UNIT_POUND,
			UnitAmount:    "0.99",
			UnitTaxAmount: "0",
			LineAmount:    "0.99",
			LineTaxAmount: "0",
		},
		{
			Sequence:      4,
			LineItemGroup: "Produce",
			ProductCode:   "90099332322",
			Description:   "Organic Bananas",
			Quantity:      "1",
			Unit:          unitv1.Unit_UNIT_POUND,
			UnitAmount:    "0.99",
			UnitTaxAmount: "0",
			LineAmount:    "0.99",
			LineTaxAmount: "0",
		},
	}
	//bh9ndiSEVFhRiyjyREKw47tm6B6Pmg
	assert.NoError(t, err)
	valueRequest := &valuev1.ValueRequest{
		Uuek:                  "bh9ndiSEVFhRiyjyREKw47tm6B6Pmg",
		PartnerTransactionRef: uuid.NewString(),
		OrderNumber:           uuid.NewString(),
		Currency:              currencyv1.Currency_CURRENCY_USD,
		OrganizationKey:       "hron3n00MDV90UjuTAi71blY95IkwA",
		StoreKey:              "hron3n00RJFObOfhRfKxf1JJDHXnvw",
		Category:              organizationcategoryv1.OrganizationCategory_ORGANIZATION_CATEGORY_GROCERY,
		ClerkIdentifier:       uuid.NewString(),
		LineItems:             lineItems,
		TotalAmount:           "30.672",
		Adjustments:           nil,
	}
	obtainValueRequest := &valuev1.ObtainValueRequest{
		ApprovalType: approvaltypev1.ApprovalType_APPROVAL_TYPE_FULL,
		Value:        valueRequest,
	}
	response, err := micaClient.ObtainValue(ctx, obtainValueRequest)
	assert.NoError(t, err)
	assert.NotEqual(t, valuev1.ObtainValueResponse_STATUS_ERROR, response.Status)
	fmt.Println(fmt.Sprintf("%v", response))
}

func loadPemBlob(envVar string) ([]byte, error) {
	//expect the pem blob to be base64 encoded
	value, ok := os.LookupEnv(envVar)
	if !ok {
		return nil, errors.New(fmt.Sprintf("environment varible %s is not set", envVar))
	}
	return base64.StdEncoding.DecodeString(value)
}
