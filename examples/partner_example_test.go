package example

import (
	"context"
	"fmt"
	"github.com/1080network/golang/partner"
	"github.com/1080network/golang/partner/proto/mica/partner/organizationv1"
	partnersvc "github.com/1080network/golang/partner/proto/mica/partner/servicev1"
	"github.com/1080network/golang/partner/proto/mica/partner/storev1"
	"github.com/1080network/golang/partner/proto/mica/partner/valuev1"
	"github.com/1080network/golang/partner/proto/micashared/common/enums/approvaltypev1"
	"github.com/1080network/golang/partner/proto/micashared/common/enums/barcodelocationv1"
	"github.com/1080network/golang/partner/proto/micashared/common/enums/barcodetypev1"
	"github.com/1080network/golang/partner/proto/micashared/common/enums/countryv1"
	"github.com/1080network/golang/partner/proto/micashared/common/enums/currencyv1"
	"github.com/1080network/golang/partner/proto/micashared/common/enums/custodialbankv1"
	"github.com/1080network/golang/partner/proto/micashared/common/enums/organizationcategoryv1"
	"github.com/1080network/golang/partner/proto/micashared/common/enums/regionv1"
	"github.com/1080network/golang/partner/proto/micashared/common/enums/unitv1"
	"github.com/1080network/golang/partner/proto/micashared/common/pingv1"
	commonv1 "github.com/1080network/golang/partner/proto/micashared/common/v1"
	"github.com/1080network/golang/shared"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
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

func TestCreateOrganization(t *testing.T) {
	ctx := context.TODO()
	micaclient, conn, err := initializeClient()
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	assert.NoError(t, err)
	request := &organizationv1.CreateOrganizationRequest{
		OrganizationRef: "2d7b57ed-c67c-4296-a7d0-90f9cfc750a8",
		Name:            "Example organization",
		Address: &commonv1.Address{
			StreetLines: []string{"1600 Pennsylvania Avenue, N.W."},
			Locality:    "Washington",
			Region:      regionv1.Region_REGION_US_DC,
			PostalCode:  "20500",
			Country:     countryv1.Country_COUNTRY_US,
		},
		Categories: []organizationcategoryv1.OrganizationCategory{organizationcategoryv1.OrganizationCategory_ORGANIZATION_CATEGORY_BOOKSTORE,
			organizationcategoryv1.OrganizationCategory_ORGANIZATION_CATEGORY_APPAREL},
		OperatingAccount: &commonv1.BankAccountDetail{
			CustodialBank: custodialbankv1.CustodialBank_CUSTODIAL_BANK_TEST_BANK,
			Bank:          "Example Bank",
			AccountRef:    "c57b2a85-9650-46c3-8267-651cbcb62aa6",
		},
		RevenueAccount: &commonv1.BankAccountDetail{
			CustodialBank: custodialbankv1.CustodialBank_CUSTODIAL_BANK_TEST_BANK,
			Bank:          "Example Bank",
			AccountRef:    "c57b2a85-9650-46c3-8267-651cbcb62aa6",
		},
		BarcodeType:     barcodetypev1.BarcodeType_BARCODE_TYPE_NONE,
		BarcodeLocation: barcodelocationv1.BarcodeLocation_BARCODE_LOCATION_NONE,
	}
	response, err := micaclient.CreateOrganization(ctx, request)
	assert.NoError(t, err)
	assert.Equal(t, organizationv1.CreateOrganizationResponse_STATUS_SUCCESS, response.Status)
}

func TestCreateStore(t *testing.T) {
	ctx := context.TODO()
	micaclient, conn, err := initializeClient()
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	assert.NoError(t, err)
	request := &storev1.CreateStoreRequest{
		OrganizationIdentifier: &storev1.CreateStoreRequest_OrganizationRef{OrganizationRef: "2d7b57ed-c67c-4296-a7d0-90f9cfc750a8"},
		StoreRef:               "20e97697-5b1c-46da-bf50-894da8bcaddc",
		StoreNumber:            "001",
		Address: &commonv1.Address{
			StreetLines: []string{"1600 Pennsylvania Avenue, N.W."},
			Locality:    "Washington",
			Region:      regionv1.Region_REGION_US_DC,
			PostalCode:  "20500",
			Country:     countryv1.Country_COUNTRY_US,
		},
	}
	response, err := micaclient.CreateStore(ctx, request)
	assert.NoError(t, err)
	assert.Equal(t, storev1.CreateStoreResponse_STATUS_SUCCESS, response.Status)
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
		OrganizationIdentifier: &organizationv1.GetOrganizationRequest_OrganizationRef{OrganizationRef: "2d7b57ed-c67c-4296-a7d0-90f9cfc750a8"},
	}
	response, err := micaClient.GetOrganization(ctx, request)
	assert.NoError(t, err)
	assert.Equal(t, organizationv1.GetOrganizationResponse_STATUS_SUCCESS, response.Status)
	fmt.Println(fmt.Sprintf("%v", response.Organization))
}

func TestUpdateOrganization(t *testing.T) {
	//this test works but a bug in the backend trashes the organization because we loose the categories
	t.SkipNow()
	return
	ctx := context.TODO()
	micaClient, conn, err := initializeClient()
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	assert.NoError(t, err)
	request := &organizationv1.GetOrganizationRequest{
		OrganizationIdentifier: &organizationv1.GetOrganizationRequest_OrganizationKey{OrganizationKey: ""},
	}
	response, err := micaClient.GetOrganization(ctx, request)
	assert.NoError(t, err)
	assert.Equal(t, organizationv1.GetOrganizationResponse_STATUS_SUCCESS, response.Status)

	updateRequest := &organizationv1.UpdateOrganizationRequest{
		OrganizationIdentifier: &organizationv1.UpdateOrganizationRequest_OrganizationKey{
			OrganizationKey: response.Organization.OrganizationKey,
		},
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

func TestObtainValueWithoutBasked(t *testing.T) {
	ctx := context.TODO()
	micaClient, conn, err := initializeClient()
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	//bh9ndiSEVFhRiyjyREKw47tm6B6Pmg
	assert.NoError(t, err)
	valueRequest := &valuev1.ValueRequest{
		Uuek:                  "bh9ndiSEVFhRiyjyREKw47tm6B6Pmg",
		PartnerTransactionRef: uuid.NewString(),
		OrderNumber:           uuid.NewString(),
		Currency:              currencyv1.Currency_CURRENCY_USD,
		OrganizationIdentifier: &valuev1.ValueRequest_OrganizationKey{
			OrganizationKey: "hron3n00MDV90UjuTAi71blY95IkwA",
		},
		StoreIdentifier: &valuev1.ValueRequest_StoreKey{StoreKey: "hron3n00RJFObOfhRfKxf1JJDHXnvw"},
		Category:        organizationcategoryv1.OrganizationCategory_ORGANIZATION_CATEGORY_GROCERY,
		ClerkIdentifier: "Clerk logged int",
		TotalAmount:     "30.672",
		Adjustments:     nil,
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

func TestObtainValueWithBasket(t *testing.T) {
	ctx := context.TODO()
	micaClient, conn, err := initializeClient()
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	assert.NoError(t, err)
	lineItems := []*commonv1.LineItem{
		{
			Sequence:      1,
			LineItemGroup: "Pasta and Grains",
			ProductCode:   "021000658930",
			Description:   "Velveta Shells and Cheese Original",
			Quantity:      "2",
			Unit:          unitv1.Unit_UNIT_ITEM,
			UnitAmount:    "3.79",
			UnitTaxAmount: "0.3127",
			LineAmount:    "7.58",
			LineTaxAmount: "0.6254",
		},
		{
			Sequence:      2,
			LineItemGroup: "Produce",
			ProductCode:   "90099332322",
			Description:   "Organic Bananas",
			Quantity:      "1",
			Unit:          unitv1.Unit_UNIT_POUND,
			UnitAmount:    "0.99",
			UnitTaxAmount: "0.082",
			LineAmount:    "0.99",
			LineTaxAmount: "0.082",
		},
	}
	valueRequest := &valuev1.ValueRequest{
		Uuek:                  "bh9ndiSEVFhRiyjyREKw47tm6B6Pmg",
		PartnerTransactionRef: uuid.NewString(),
		OrderNumber:           uuid.NewString(),
		Currency:              currencyv1.Currency_CURRENCY_USD,
		OrganizationIdentifier: &valuev1.ValueRequest_OrganizationKey{
			OrganizationKey: "hron3n00MDV90UjuTAi71blY95IkwA",
		},
		StoreIdentifier: &valuev1.ValueRequest_StoreKey{StoreKey: "hron3n00RJFObOfhRfKxf1JJDHXnvw"},
		Category:        organizationcategoryv1.OrganizationCategory_ORGANIZATION_CATEGORY_GROCERY,
		ClerkIdentifier: uuid.NewString(),
		LineItems:       lineItems,
		TotalAmount:     "30.672",
		Adjustments:     nil,
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
